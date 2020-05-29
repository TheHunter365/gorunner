package main

import "github.com/thehunter365/gorunner/runner"

func main() {

	server := runner.NewServer(":8080")
	server.Start()
	
	
}
