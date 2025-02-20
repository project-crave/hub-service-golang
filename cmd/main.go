package main

import (
	"crave/hub/cmd/lib"
)

func main() {

	//router.Use()

	startLib()
}

func startLib() {
	lib.Start(nil)
}
