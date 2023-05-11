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
    port = "3000"

    d1Path  = config.TempDirPath + "d1.json"
    d2Path  = config.TempDirPath + "d2.json"
    allPath = config.TempDirPath + "all.json"
)

func main() {
    gin.SetMode(gin.ReleaseMode)
    router := gin.Default()
    router.Use(gin.Logger())
    router.Use(gzip.Gzip(gzip.DefaultCompression))
    router.LoadHTMLGlob(`templates/*.html`)

    router.Static("/static", "static")
    router.Static("/tmp", config.TempDirPath)

    router.GET("/", func(c *gin.Context) {
        var fName1 = script.DeBase64(c.Query("fName1"))
        var fName2 = script.DeBase64(c.Query("fName2"))

        var diff1, diff2, allDiff []config.Data
        var same = "nothing"
        if haveCompareFile(fName1, fName2) {
            diff1 = io.ReadJson(d1Path)
            diff2 = io.ReadJson(d2Path)
            allDiff = io.ReadJson(allPath)
            if checkSame(allDiff) {
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
        var fNames = uploadFile(c)
        var url string
        if len(fNames) > 1 {
            fileData1 := io.Readfile(fNames[0])
            fileData2 := io.Readfile(fNames[1])
            if fileNotBlank(fileData1, fileData2) {
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

func haveCompareFile(f1, f2 string) bool {
    return f1 != "" && f2 != ""
}

func checkSame(allDiff []config.Data) bool {
    return len(allDiff) == 0
}

func uploadFile(c *gin.Context) []string {
    var fNames []string
    form, _ := c.MultipartForm()
    for _, f := range form.File["uploads"] {
        path := config.TempDirPath + f.Filename
        fNames = append(fNames, f.Filename)
        c.SaveUploadedFile(f, path)
    }
    return fNames
}
func fileNotBlank(f1, f2 []string) bool {
    return len(f1) > 0 && len(f2) > 0
}
