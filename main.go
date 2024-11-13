package main

import (
	"log"
	"net/http"
	"path/filepath"

	"github.com/0mlml/cfgparser"
	"github.com/gin-gonic/gin"
)

func main() {
	defaultConfig := &cfgparser.Config{}
	defaultConfig.Literal(
		map[string]bool{},
		map[string]string{},
		map[string]int{
			"port": 8080,
		},
		map[string]float64{},
	)

	cfgparser.SetDefaultConfig(defaultConfig)

	config := &cfgparser.Config{}
	if err := config.From("server.cfg"); err != nil {
		log.Fatalf("Error parsing config file: %v\nMake sure you have created server.cfg from example.cfg", err)
	}

	token := config.String("token")
	if len(token) == 0 {
		log.Fatalf("Token is empty in config\n")
	}

	router := gin.Default()
	router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		// Source
		file, err := c.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "get form err: %s", err.Error())
			return
		}

		filename := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, filename); err != nil {
			c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
			return
		}

		c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
	})
	router.Run(":8080")
}
