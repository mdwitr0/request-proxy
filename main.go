package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
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

	client := resty.New().EnableTrace()

	if body.Headers != nil {
		client.SetHeaders(body.Headers)
	}

	if body.Params != nil {
		client.SetQueryParams(body.Params)
	}

	var result interface{}

	switch body.Method {
	case "POST":
		_, err = client.R().
			SetBody(body.Data).
			SetResult(&result).
			Post(body.Url)
		break
	case "GET":
		_, err = client.R().
			SetResult(&result).
			Get(body.Url)
		break
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
