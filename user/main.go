package main

import (
	"github.com/gin-gonic/gin"

	srv "github.com/qweijiaq/common"
	_ "github.com/qweijiaq/user/api"
	"github.com/qweijiaq/user/router"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)
	srv.Run(r, "users", ":80")
}
