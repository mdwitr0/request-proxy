package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
)

type ProxyQuery struct {
	Url     string            `json:"url"`
	Method  string            `json:"method"`
	Params  map[string]string `json:"params"`
	Headers map[string]string `json:"headers"`
	Data    interface{}       `json:"data"`
}

func proxyHandle(c *gin.Context) {
	var body ProxyQuery
	err := c.ShouldBindJSON(&body)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "")
	}

	client := resty.New()

	var result interface{}
	_, err = client.R().
		SetBody(body.Data).
		SetHeaders(body.Headers).
		SetQueryParams(body.Params).
		SetResult(&result).
		EnableTrace().
		Post(body.Url)

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, "")
	}

	c.JSON(http.StatusOK, result)
}

func main() {
	r := gin.Default()
	r.POST("", proxyHandle)
	err := r.Run()
	if err != nil {
		return
	}
}
