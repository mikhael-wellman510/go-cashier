package main

import (
	"mikhael-project-go/cmd/app"
)

func main() {

	app := app.App{}

	app.ConnectDb()
	app.Routes()
	app.Run()

}
