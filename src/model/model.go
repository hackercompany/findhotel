package model

import (
	"logger"

	"database/sql"
	"fmt"
	"regexp"
)

var log = logger.Log

// This is the main interface for data interation with the database
// It exposes function to validate the data in objects
// As well as holding a db handler for that specific transaction.
type Geolocation struct {
	IP, Ccode, Country, City, Lat, Long, MysteryValue string
	Handler                                           *sql.DB
}

func (g *Geolocation) Get() error {
	rows, err := g.Handler.Query(GET_IP_DATA, g.IP)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var ccode, country, city, mystryValue string
		var lat, long float64

		if err := rows.Scan(&ccode, &country, &city, &lat, &long, &mystryValue); err != nil {
			break
		}
		g.Ccode = ccode
		g.Country = country
		g.City = city
		g.MysteryValue = mystryValue
		// This is done as in memory objects all work on strings
		// For db storage optimisation specific datatypes have been defined
		g.Lat = fmt.Sprintf("%f", lat)
		g.Long = fmt.Sprintf("%f", long)
	}
	return nil
}

func (g *Geolocation) Insert() error {
	insert, err := g.Handler.Query(INSERT_IP_DATA, g.IP, g.Ccode, g.Country, g.City, g.Lat, g.Long, g.MysteryValue)

	if err != nil {
		log.Errorln("Geolocation Insert", "Mysql Error", err.Error())
		return err
	}
	defer insert.Close()

	return nil
}

func (g *Geolocation) Validate() bool {
	valid := true
	if !g.ValidIp() || !g.ValidCc() || !g.ValidCountry() || !g.ValidCity() || !g.ValidLatLon() {
		valid = false
	}
	return valid
}

func (g *Geolocation) ValidIp() bool {
	// Currently only supports IPv4
	// TODO: Make re expression configurable and add IPv6 support
	validEx := regexp.MustCompile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if !validEx.MatchString(g.IP) {
		return false
	}
	return true
}

func (g *Geolocation) ValidCc() bool {
	// ISO country codes are 2 digits in length
	if len(g.Ccode) == 2 {
		return true
	}
	return false
}

func (g *Geolocation) ValidCountry() bool {
	if g.Country != "" {
		return true
	}
	return false
}

func (g *Geolocation) ValidCity() bool {
	if g.Country != "" {
		return true
	}
	return false
}

func (g *Geolocation) ValidLatLon() bool {
	validEx := regexp.MustCompile(`^(\-?\d+(\.\d+)?)$`)
	if !validEx.MatchString(g.Lat) {
		return false
	}
	if !validEx.MatchString(g.Long) {
		return false
	}
	return true
}
