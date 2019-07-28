package model

var (
	GET_IP_DATA    = "SELECT ccode, country, city, lat, lon, mystry FROM ip_data WHERE ip=?"
	INSERT_IP_DATA = "INSERT INTO ip_data (ip, ccode, country, city, lat, lon, mystry) VALUES(?, ?, ?, ?, ?, ?, ?)"
	DELETE_IP_DATA = "DELETE FROM ip_data WHERE ip=?"
)
