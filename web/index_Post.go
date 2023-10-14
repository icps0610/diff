package web

import (
    "fmt"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"

    . "iDiff/conf"
    "iDiff/diff"
    "iDiff/io"
    "iDiff/script"
)

func IndexPost(c *gin.Context) {
    var files []string
    form, _ := c.MultipartForm()
    for _, f := range form.File["uploads"] {
        path := filepath.Join(TempDirPath, f.Filename)
        c.SaveUploadedFile(f, path)

        files = append(files, f.Filename)
    }

    var url string
    if len(files) >= 2 {
        filePath1 := filepath.Join(TempDirPath, files[0])
        filePath2 := filepath.Join(TempDirPath, files[1])

        fileContent1 := io.Readfile(filePath1)
        fileContent2 := io.Readfile(filePath2)

        if fileNotBlank(fileContent1, fileContent2) {
            diff1, diff2, allDiff, _ := diff.Run(fileContent1, fileContent2)

            io.SaveJson(diff1, Diff1Path)
            io.SaveJson(diff2, Diff2Path)
            io.SaveJson(allDiff, AllPath)

        }
        url = fmt.Sprintf(`/?file1=%s&file2=%s`, script.EnBase64(files[0]), script.EnBase64(files[1]))
    }

    c.Redirect(http.StatusMovedPermanently, url)
}

func fileNotBlank(f1, f2 []string) bool {
    return len(f1) > 0 && len(f2) > 0
}
