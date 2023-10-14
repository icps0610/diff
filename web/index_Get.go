package web

import (
    "github.com/gin-gonic/gin"
    "net/http"

    . "iDiff/conf"
    "iDiff/io"
    "iDiff/script"
)

func Index(c *gin.Context) {
    file1 := c.Query("file1")
    file2 := c.Query("file2")

    fileName1 := script.DeBase64(file1)
    fileName2 := script.DeBase64(file2)

    var diff1, diff2, allDiff []Diff
    var same = "nothing"
    if filesNotBlank(fileName1, fileName2) {
        diff1 = io.ReadJson(Diff1Path)
        diff2 = io.ReadJson(Diff2Path)
        allDiff = io.ReadJson(AllPath)

        if checkSame(allDiff) {
            same = "true"
        } else {
            same = "false"
        }
    }

    c.HTML(http.StatusOK, `index.html`, gin.H{
        "fileName1": fileName1,
        "fileName2": fileName2,
        "diff1":     diff1,
        "diff2":     diff2,
        "allDiff":   allDiff,
        "same":      same,
    })

}

func filesNotBlank(f1, f2 string) bool {
    return f1 != "" && f2 != ""
}

func checkSame(allDiff []Diff) bool {
    return len(allDiff) == 0
}
