package main

import "HZ_proj/Backend/HTTPHandler"

func main() {
	server := HTTPHandler.ServerRoutineFactory()
	server.RunServer()
}
