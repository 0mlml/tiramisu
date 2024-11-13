package backend

import (
	"github.com/gin-gonic/gin"
)

var router *gin.Engine
var db *DB

func Main() {
	var err error
	db, err = newDB(config.String("db_path"))
	if err != nil {
		panic(err)
	}

	router = gin.Default()
	router.MaxMultipartMemory = 8 << 20

	initializeRoutes(router)

	router.Run(":8080")
}
