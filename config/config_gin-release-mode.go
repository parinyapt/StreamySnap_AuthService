package config

import "github.com/gin-gonic/gin"

func GinReleaseModeSetup() {
	gin.SetMode(gin.ReleaseMode)
}