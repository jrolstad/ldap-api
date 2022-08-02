package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jrolstad/ldap-api/internal/pkg/core"
	"net/http"
)

func returnJsonResult(c *gin.Context, data interface{}, err error) {
	if err != nil {
		c.Status(http.StatusInternalServerError)
	} else if core.IsNil(data) {
		c.Status(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, data)
	}
}
