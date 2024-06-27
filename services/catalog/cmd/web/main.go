package main

import "github.com/malikilamalik/stacksandbrews/services/catalog/server"

func main() {
	s := server.Init()
	s.Start()
}
