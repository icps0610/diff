package config

import (
    "runtime"
)

var TempDirPath = getTempPath()

type Data struct {
    Idx  int
    Ele1 string
    Ele2 string
}

func getTempPath() string {
    if runtime.GOOS == "windows" {
        return `z:\`
    }
    return `/tmp/`
}
