package io

import (
    "encoding/json"
    "fmt"
    "github.com/malexdev/utfutil"
    "io/ioutil"
    "strings"

    . "iDiff/conf"
    // "diff/script"
)

func Readfile(path string) []string {
    content, _ := utfutil.ReadFile(path, utfutil.UTF8)
    return split(string(content))
}

func ReadJson(path string) []Diff {
    dat, err := ioutil.ReadFile(path)
    printError(err)
    var reaJson []Diff
    err = json.Unmarshal(dat, &reaJson)
    printError(err)
    return reaJson
}

func SaveJson(datas interface{}, path string) {
    file, err := json.MarshalIndent(datas, "", "  ")
    printError(err)
    err = ioutil.WriteFile(path, file, 0777)
    printError(err)
}

func split(content string) []string {
    var lines []string
    for _, line := range strings.Split(string(content), "\n") {
        line = strings.TrimSpace(line)
        if len(line) > 0 {
            lines = append(lines, line)
        }
    }
    return lines
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
