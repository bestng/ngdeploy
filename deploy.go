package main

import (
	"fmt"
	"os"
	"text/template"
)

type Lang struct {
	Name   string
	Domain string
	Port   string
	Path   string
}

func (lang Lang) Deploy() {
	userfile := lang.Domain + ".config"
	config, err := os.Create(userfile)
	defer config.Close()
	if err != nil {
		panic(err)
	}

	switch {
	case lang.Name == "node":
		// do ...
	case lang.Name == "rails":
		t, _ := template.ParseFiles("config/rails/nginx.config")
		t.Execute(config, &lang)
	default:
		// do ..
	}
}

func main() {
	rails := Lang{"rails", "ruby.clss.co", "80"}
	rails.Deploy()
	fmt.Println(rails.Name)
}
