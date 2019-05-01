package main

import (
	"github.com/gin-gonic/gin"
)

type aboutData struct {
	baseTemplateData
}

func aboutPage(c *gin.Context) {

data := new(aboutData)

defer resp(c, 200, "about.html", data)

data.DisableHH = true
}
