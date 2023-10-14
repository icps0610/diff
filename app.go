package main

import (
    . "iDiff/conf"
    "iDiff/web"
)

func main() {
    router := web.Service()
    router.GET("/", web.Index)

    router.POST("/", web.IndexPost)

    router.Run(":" + Port)
}
