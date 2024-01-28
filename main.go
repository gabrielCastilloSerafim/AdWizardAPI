package main

func main() {
	mongoClient := createDatabase()
	startServer(mongoClient)
}
