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
		c.JSON(http.StatusInternalServerError, "")
	}

	var result interface{}

	client := resty.New().EnableTrace()

	if body.Headers != nil {
		client.SetHeaders(body.Headers)
	}

	if body.Params != nil {
		client.SetQueryParams(body.Params)
	}

	if body.Method == "POST" {
		_, err = client.R().
			SetBody(body.Data).
			SetResult(&result).
			SetHeaders(body.Headers).
			SetQueryParams(body.Params).
			Post(body.Url)
	} else {
		_, err = client.R().
			SetResult(&result).
			SetQueryParams(body.Params).
			Get(body.Url)
	}

	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, "")
	}

	c.JSON(http.StatusOK, result)
}

func main() {
	r := gin.Default()
	r.POST("", proxyHandle)
	err := r.Run(":8001")
	if err != nil {
		return
	}
}
