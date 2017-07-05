package main

import "server"

func main() {

	server.Run(server.Options{Addr: ":8080"})
}
