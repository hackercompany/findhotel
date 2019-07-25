package middleware

import (
	"database/sql"
	"fmt"

	"config"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var DBHandler *sql.DB

func DoInit() {
	var DbConnErr error

	MySqlUsername := config.Config.GetString("database.user")
	MySqlPassword := config.Config.GetString("database.pass")
	MySqlHost := config.Config.GetString("database.host")
	MySqlPort := config.Config.GetString("database.port")
	MySqlDB := config.Config.GetString("database.db")
	DBHandler, DbConnErr = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", MySqlUsername, MySqlPassword, MySqlHost, MySqlPort, MySqlDB))

	if DbConnErr != nil {
		panic(DbConnErr)
	}

	DBHandler.SetMaxOpenConns(config.Config.GetInt("database.max_connections"))
	DBHandler.SetMaxIdleConns(config.Config.GetInt("database.max_idle"))

}

func MySQLConnector(c *gin.Context) {
	//put db object to gin context
	c.Set("mysql", DBHandler)
	c.Next()
}
