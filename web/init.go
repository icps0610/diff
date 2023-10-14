package web

import (
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"

    "iDiff/conf"
)

func Service() *gin.Engine {
    gin.SetMode(gin.ReleaseMode)

    router := gin.Default()
    router.Use(gzip.Gzip(gzip.DefaultCompression))
    router.Static("/tmp", conf.TempDirPath)
    router.Static("/static", conf.StaticDirPath)

    router.LoadHTMLGlob(conf.HtmlPath)

    return router
}
