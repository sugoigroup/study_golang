package main

import (
	"./testserver"
	"./utils"
)


func main() {
	utils.PrintBasic()

	testserver.Runserver()
}
