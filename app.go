package main

import (
    "fmt"
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
    "net/http"

    "diff/config"
    "diff/diff"
    "diff/io"
    "diff/script"
)

var (
    port       = "3000"
    tmpDirPath = script.GetTempPath()

    d1Path  = tmpDirPath + "d1.json"
    d2Path  = tmpDirPath + "d2.json"
    allPath = tmpDirPath + "all.json"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gzip.Gzip(gzip.DefaultCompression))
    router.LoadHTMLGlob(`templates/*.html`)

    router.Static("/static", "static")
    router.Static("/tmp", tmpDirPath)

    router.GET("/", func(c *gin.Context) {
        var fName1 = script.DeBase64(c.Query("fName1"))
        var fName2 = script.DeBase64(c.Query("fName2"))
        var diff1, diff2, allDiff []config.Data
        var same string

        if fName1 != "" && fName2 != "" {
            diff1 = io.ReadJson(d1Path)
            diff2 = io.ReadJson(d2Path)
            allDiff = io.ReadJson(allPath)
            if len(diff1) > 0 || len(diff2) > 0 {
                same = "true"
            } else {
                same = "false"
            }
        }

        c.HTML(http.StatusOK, `index.html`, gin.H{
            "fName1":  fName1,
            "fName2":  fName2,
            "diff1":   diff1,
            "diff2":   diff2,
            "allDiff": allDiff,
            "same":    same,
        })
    })

    router.POST("/", func(c *gin.Context) {
        var fNames []string
        form, _ := c.MultipartForm()
        for _, f := range form.File["uploads"] {
            path := tmpDirPath + f.Filename
            fNames = append(fNames, f.Filename)
            c.SaveUploadedFile(f, path)
        }

        var url string
        if len(fNames) > 1 {
            fileData1 := io.Readfile(tmpDirPath + fNames[0])
            fileData2 := io.Readfile(tmpDirPath + fNames[1])
            if len(fileData1) > 0 && len(fileData2) > 0 {

                diff1, diff2, allDiff, _ := diff.Run(fileData1, fileData2)

                io.SaveJson(diff1, d1Path)
                io.SaveJson(diff2, d2Path)
                io.SaveJson(allDiff, allPath)
            }

            url = fmt.Sprintf(`/?fName1=%s&fName2=%s`, script.EnBase64(fNames[0]), script.EnBase64(fNames[1]))
        }

        c.Redirect(http.StatusMovedPermanently, url)

    })
    router.Run(":" + port)
}

var _ = fmt.Println
