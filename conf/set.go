package conf

import (
    "fmt"
    "os"
    "path/filepath"
    "runtime"
)

var (
    Port = "8030"

    RootDirPath = getRootDirPath()
    TempDirPath = getTempDirPath()

    StaticDirPath = filepath.Join(RootDirPath, "static")
    HtmlPath      = filepath.Join(RootDirPath, `templates`, `*.html`)

    Diff1Path = filepath.Join(TempDirPath, "diff1.json")
    Diff2Path = filepath.Join(TempDirPath, "diff2.json")
    AllPath   = filepath.Join(TempDirPath, "all.json")
)

type Diff struct {
    Idx  int
    Ele1 string
    Ele2 string
}

func init() {
    os.Mkdir(TempDirPath, 0777)
}

func getRootDirPath() string {
    path, err := os.Executable()
    printError(err)
    if string(path[:11]) == `z:\go-build` {
        path, err = os.Getwd()
        printError(err)
        return path + `\`
    }
    path = filepath.Dir(path)

    if runtime.GOOS == "windows" {
        return path + `\`
    }
    return path + `/`
}

func getTempDirPath() string {
    if runtime.GOOS == "windows" {
        return `z:\diff`
    }
    return `/tmp/diff`
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
