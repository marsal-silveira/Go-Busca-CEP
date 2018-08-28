package main

import (
	"cep-provider/app/api"
)

func main() {

	server := api.ConfigureServer()
	server.ListenAndServe()
}
