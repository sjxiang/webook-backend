package main

func main() {
	server, err := initServer()
	if err != nil {
		panic(err)
	}
	server.Start()
}
