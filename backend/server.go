package backend

import (
	"flag"

	"github.com/gin-gonic/gin"
)

var (
	router *gin.Engine
	db     *DB
	dbPath = flag.String("db", "tiramisu.db", "Path to the database file")
	port   = flag.String("port", "8080", "Port to run the server on")
)

func Main() {
	var err error
	db, err = newDB(*dbPath)
	if err != nil {
		panic(err)
	}

	router = gin.Default()
	router.MaxMultipartMemory = 8 << 20

	initializeRoutes(router)

	router.Run(":" + *port)
}
