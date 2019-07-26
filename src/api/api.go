package api

import (
	"database/sql"

	"logger"
	"model"

	"github.com/gin-gonic/gin"
)

type Response struct {
	IP      string `json:"ip,omitempty"`
	City    string `json:"city,omitempty"`
	Country string `json:"country,omitempty"`
	Lat     string `json:"latitude,omitempty"`
	Long    string `json:"longitude,omitempty"`
	Ccode   string `json:"country_code,omitempty"`
	MVal    string `json:"mystery_value,omitempty"`
	Status  string `json:"status,omitempty"`
}

var log = logger.Log

// App health URL
// Response format:
//	{"message": "pong"}
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Ip information URL
// Response Format:
//	{
//	  "ip": "100.10.114.117",
//	  "city": "Lake Nellaberg",
//	  "country": "Kenya",
//	  "latitude": "-81.805504",
//	  "longitude": "-36.449104",
//	  "country_code": "GQ",
//	  "status": "success"
//	}

func IpInfo(c *gin.Context) {
	resp := Response{Status: "fail"}
	ipAddress := c.Query("ip")

	ipDAO := model.Geolocation{IP: ipAddress, Handler: c.MustGet("mysql").(*sql.DB)}
	// Sanatising input param
	if !ipDAO.ValidIp() {
		c.JSON(400, resp)
		return
	}

	err := ipDAO.Get()
	if err != nil {
		c.JSON(400, resp)
		log.Errorln("IpInfo", "MySQL error", err.Error())
		return
	}

	resp.IP = ipDAO.IP
	resp.City = ipDAO.City
	resp.Country = ipDAO.Country
	resp.Lat = ipDAO.Lat
	resp.Long = ipDAO.Long
	resp.Ccode = ipDAO.Ccode
	resp.MVal = ipDAO.MysteryValue
	resp.Status = "success"
	c.JSON(200, resp)
	return
}
