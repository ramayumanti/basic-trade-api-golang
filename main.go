package main

import "basictrade-api/routers"

var PORT = ":9090"

func main() {
	routers.StartServer().Run(PORT)
}
