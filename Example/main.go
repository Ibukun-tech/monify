package main

import (
	"context"

	"github.com/Ibukun-tech/monify"
)

func main() {
	ctx := context.Background()
	m := monify.NewMonify("jsjsjsj")
	m.GenerateToken(ctx)
}
