package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
)

func main() {
	router := gin.Default()

	router.GET("/*any", func(c *gin.Context) {
		c.File("Robert Mulyukov.pdf")
	})

	fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))
	if err := http.ListenAndServe(os.Getenv("PORT"), router); err != nil {
		log.Fatal(err)
	}
}
