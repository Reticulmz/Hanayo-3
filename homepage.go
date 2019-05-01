package main

import (
	"github.com/gin-gonic/gin"
)

type homepageData struct {
	baseTemplateData
}

func homepagePage(c *gin.Context) {

data := new(homepageData)

defer resp(c, 200, "home.html", data)

data.DisableHH = true
}
