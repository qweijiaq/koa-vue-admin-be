package users

import (
	"github.com/gin-gonic/gin"
	"github.com/qweijiaq/user/router"
	"log"
)

func init() {
	log.Println("init users router")
	router.Register(&RouterUser{})
}

type RouterUser struct {
}

func (*RouterUser) Route(r *gin.Engine) {
	h := &HandlerUser{}
	r.POST("/project/login/getCaptcha", h.getCaptcha)
}
