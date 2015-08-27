package main

import "github.com/go-martini/martini"

func main() {
  m := martini.Classic()

  m.Get("/", func() string {
    return "Hello world!"
  })

  m.Get("/misty_bath", func() string {
    return "THERE's SHIT EVERYWHERE!!!!!!"
  })

  m.Run()
}