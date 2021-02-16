package main

import "github.com/ggdream/mochat-server/app"


func main() {
	if err := app.New(); err != nil {
		panic(err)
	}
}
