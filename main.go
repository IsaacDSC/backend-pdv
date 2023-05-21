package main

import (
	"backend-pdv/src/infra"
	"backend-pdv/src/infra/environments"
)

func main() {
	environments.StartEnv()
	infra.StartServer()
}
