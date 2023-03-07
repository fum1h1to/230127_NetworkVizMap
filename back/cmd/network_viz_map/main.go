package main

import (
	"NetworkVizMap/web/viz_server"
)

func main() {
	server := viz_server.CreateWebServer()
	server.StartServer()
}