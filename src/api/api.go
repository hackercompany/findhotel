package api

import (
	"logger"
	"model"

	"database/sql"

	"github.com/gin-gonic/gin"
)

type Response struct {
	IP      string `json:"ip,ommitempty"`
	City    string `json:"city,ommitempty"`
	Country string `json:"country,ommitempty"`
	Lat     string `json:"latitude,ommitempty"`
	Long    string `json:"longitude,ommitempty"`
	Ccode   string `json:"country_code,ommitempty"`
	Status  string `json:"status,ommitempty"`
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
	if !ipDAO.ValidIp() {
		// Sanatising input param
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
	resp.Status = "success"
	c.JSON(200, resp)
	return
}
