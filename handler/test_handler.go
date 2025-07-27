package handler

import (
	"UnionEntrace/caller"
	"github.com/gin-gonic/gin"
	"github.com/nacos-group/nacos-sdk-go/v2/vo"
	"net/http"
)

func TestHandler(c *gin.Context) {
	config, err := caller.ConfigClient.GetConfig(vo.ConfigParam{
		DataId: "test_config",
	})
	if err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"config": config,
	})
}
