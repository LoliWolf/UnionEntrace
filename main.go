package main

import "UnionEntrace/caller"

func main() {
	r := setupRouter()
	caller.Init()

	r.Run(":28080")
}
