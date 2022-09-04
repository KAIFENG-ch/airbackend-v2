package main

import "test/api"

func main() {
	e := api.Router()
	_ = e.Run(":8080")
}
