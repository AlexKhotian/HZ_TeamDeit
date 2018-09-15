package main

import "HZ_proj/HTTPHandler"

func main() {
	server := HTTPHandler.ServerRoutineFactory()
	server.RunServer()
}
