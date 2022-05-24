package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
)

type ViewData struct {
	Title   string
	Message string
}

func main() {
	arrowDown := '↓'
	smile := '🙃'

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		data := ViewData{
			Title: "Robert's CV",
			Message: fmt.Sprintf("Hi %c! You can choose the language of my CV bellow %c",
				smile, arrowDown),
		}
		tmpl, _ := template.ParseFiles("index.html")
		if err := tmpl.Execute(c.Writer, data); err != nil {
			log.Fatal(err)
		}
	})
	router.GET("/cv/en", func(c *gin.Context) {
		c.File("Robert Mulyukov.pdf")
	})
	router.GET("/cv/ru", func(c *gin.Context) {
		c.File("Мулюков Роберт.pdf")
	})

	fmt.Printf("Starting server at port %s\n", os.Getenv("PORT"))
	if err := http.ListenAndServe(":"+os.Getenv("PORT"), router); err != nil {
		log.Fatal(err)
	}
}
