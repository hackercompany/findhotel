package logger

import (
	"fmt"
	"os"

	"config"

	"github.com/Sirupsen/logrus"
)

var Log = logrus.New()

func DoInit() {
	logfile, err := os.OpenFile(config.Config.GetString("logger.file_path"), os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(fmt.Sprintf("Can't open log file %s", err.Error()))
	}
	Log.Out = logfile
}
