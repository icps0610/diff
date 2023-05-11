package main

import (
    "fmt"
    "github.com/gin-contrib/gzip"
    "github.com/gin-gonic/gin"
    "net/http"
    "runtime"

    "diff/config"
    "diff/io"
    "diff/script"
)

var (
    port       = "3000"
    tmpDirPath = GetTmpPath()

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
        var diff1, diff2, all []config.Data
        var same int

        if fName1 != "" && fName2 != "" {
            diff1 = io.ReadJson(d1Path)
            diff2 = io.ReadJson(d2Path)
            all = io.ReadJson(allPath)
            if len(diff1) > 0 || len(diff2) > 0 {
                same = 1
            } else {
                same = 2
            }
        }

        c.HTML(http.StatusOK, `index.html`, gin.H{
            "fName1": fName1,
            "fName2": fName2,
            "diff1":  diff1,
            "diff2":  diff2,
            "all":    all,
            "same":   same,
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

        if len(fNames) > 1 {
            fileData1 := io.Readfile(tmpDirPath + fNames[0])
            fileData2 := io.Readfile(tmpDirPath + fNames[1])

            diff1, diff2, all, _ := script.ArrCompare(fileData1, fileData2)

            io.SaveJson(diff1, d1Path)
            io.SaveJson(diff2, d2Path)
            io.SaveJson(all, allPath)
        }

        url := fmt.Sprintf(`/?fName1=%s&fName2=%s`, script.EnBase64(fNames[0]), script.EnBase64(fNames[1]))
        c.Redirect(http.StatusMovedPermanently, url)

    })
    router.Run(":" + port)
}

func GetTmpPath() string {
    if runtime.GOOS == "windows" {
        return `z:\`
    }
    return `/tmp/`
}

var _ = fmt.Println
