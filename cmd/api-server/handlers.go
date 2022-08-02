package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
	"log"
)

func returnResult(c *gin.Context, data interface{}, err error) {
	if err != nil {
		log.Println(fmt.Sprintf("Error:%v", err))
		c.JSON(500, "Error when processing request")
	} else if core.IsNil(data) {
		c.JSON(404, data)
	} else {
		c.JSON(200, data)
	}
}
