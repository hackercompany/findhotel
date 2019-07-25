package main

import (
	"config"
	"middleware"
	"model"

	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

func main() {
	config.DoInit()
	logger.DoInit()
	log := logger.log

	dataFile, err := os.Open(config.Config.GetString("sample_data_file"))
	if err != nil {
		log.Errorln("Import Cron", "File not found")
		panic(err)
	}

	middleware.DoInit()

	totalRecords := 0
	successInserts := 0
	failedInserts := 0
	start := time.Now()

	r := csv.NewReader(dataFile)

	var wg sync.WaitGroup

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

		totalRecords += 1

		if len(record) < 6 {
			failedInserts += 1
			continue
		}

		geo := model.Geolocation{Handler: middleware.DBHandler,
			IP:           record[0],
			Ccode:        record[1],
			Country:      record[2],
			City:         record[3],
			Lat:          record[4],
			Long:         record[5],
			MysteryValue: record[6],
		}

		if !geo.Validate() {
			failedInserts += 1
			continue
		}
		wg.Add(1)

		// TODO: Better atomic way to maintain counts
		go func(suc *int, fail *int) {
			defer wg.Done()
			err = geo.Insert()
			if err != nil {
				failedInserts += 1
			} else {
				successInserts += 1
			}
		}(&successInserts, &failedInserts)
	}
	wg.Wait()

	timeElapsed := time.Since(start)

	fmt.Println("Total Records", totalRecords)
	fmt.Println("Successful Insertions", successInserts)
	fmt.Println("Failed Insertions", failedInserts)
	fmt.Println("Time Taken", timeElapsed.Seconds(), "Sec")

}
