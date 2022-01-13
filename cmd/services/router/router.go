package router

import "github.com/gin-gonic/gin"

func New(release bool) (router *gin.Engine) {
	if release {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.Use(gin.Recovery())
	r.RedirectTrailingSlash = false

	return r
}
