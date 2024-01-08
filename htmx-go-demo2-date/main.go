package main

// Source: https://github.com/Bamimore-Tomi/go-templates-guide/tree/examaple-08 with updates

import (
	"log"
	"os"
	"text/template"
	"time"
)

type data struct {
	IsAdmin bool
	Users   []User
}

type User struct {
	Name      string
	Birthdate time.Time
	Test      string
}

// Declare type pointer to a template
var temp *template.Template

func init() {
	temp = template.Must(template.New("template-08.html").Funcs(funcMap).ParseFiles("template-08.html"))
}

func dateSince(Birthdate time.Time) string {
	return time.Since(Birthdate).String()
}

var funcMap = template.FuncMap{
	"calculatedDate": dateSince,
}

var (
	html = `<div><h1>GoLinuxCloud</h1>
			<p>This is an html document!</p></div>`
)

func main() {

	users := []User{
		{
			Name:      "karim",
			Birthdate: time.Now().Add(-2 * time.Hour),
			Test:      html,
		},
		{
			Name:      "fadel",
			Birthdate: time.Now().Add(-3 * time.Hour),
			Test:      html,
		},
		{
			Name:      "",
			Birthdate: time.Now().Add(-4 * time.Hour),
			Test:      html,
		},
	}

	test := data{IsAdmin: true, Users: users}

	//timeNow := time.Now()
	err := temp.Execute(os.Stdout, &test)
	if err != nil {
		log.Fatalln(err)
	}
}

// Hi,

// Time before formatting : 2021-10-04 18:01:59.6659258 +0100 WAT m=+0.004952101
// Time After formatting : 09-04-2021
