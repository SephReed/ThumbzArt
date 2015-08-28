package main

import (
	"encoding/xml"
    "net/http"
    "html/template"

    "github.com/labstack/echo"
    "github.com/unrolled/render" // or "gopkg.in/unrolled/render.v1"
)

type ExampleXml struct {
    XMLName xml.Name `xml:"example"`
    One     string   `xml:"one,attr"`
    Two     string   `xml:"two,attr"`
}

// var templates = template.Must(template.ParseGlob("templates/*"))

func main() {
	template.Must(template.ParseGlob("templates/sub_templates/*.tmpl"))

    r := render.New(render.Options{ IndentJSON: true, Layout: "layout" })

    e := echo.New()

    // Routes
    e.Get("/", func(c *echo.Context) error {
        r.JSON(c.Response().Writer(), http.StatusOK, map[string]string{"welcome": "This is rendered JSON!"})
        return nil
    })

    e.Get("/test", func(c *echo.Context) error {
        c.Response().Writer().Write([]byte("Welcome, visit sub pages now."))
        return nil
    })

    e.Get("/data", func(c *echo.Context) error {
        r.Data(c.Response().Writer(), http.StatusOK, []byte("Some binary data here."))
        return nil
    })

	e.Get("/text", func(c *echo.Context) error {
		r.Text(c.Response().Writer(), http.StatusOK, "Plain text here")
        return nil
    })


	e.Get("/html", func(c *echo.Context) error {
		pageData := map[string]string{"name": "Seph", "css": "/css/test.css" };

		r.HTML(c.Response().Writer(), http.StatusOK, "example", pageData)
	    return nil
    })


	e.Static("/css/", "public/css")
	e.Static("/images/", "public/images")
	e.Favicon("images/favicon.ico")



    e.Run(":3000")
}




