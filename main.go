// Code generated by hertz generator.

package main

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/renjs2001/WebManager/client"
	"github.com/renjs2001/WebManager/db"
)

func main() {
	client.Init()
	db.Init()
	h := server.Default()

	register(h)
	h.Spin()
}
