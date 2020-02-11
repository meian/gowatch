package main

import (
	"os"

	"github.com/meian/gowatch/app"
)

func main() {
	app.NewApp().Run(os.Args)
}
