package main

import (
	"UnionEntrace/caller"
	"UnionEntrace/router"
)

func main() {
	r := router.SetupRouter()
	caller.Init()

	r.Run(":28080")
}
