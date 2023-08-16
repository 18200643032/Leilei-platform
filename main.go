package main

import (
	"strings"
	"net/http"
	"github.com/gin-gonic/gin"
	"Leilei-platform/config"
	"Leilei-platform/router"
)

func main() {
	gin.SetMode(config.Config.BaseModel)
	r := router.SystemRouter()
	if err := http.ListenAndServe(strings.Join([]string{config.Config.BaseHost,
		":", config.Config.BasePort}, ""), r); err != nil {
		panic("Port occupancy!!!")
	}

}
