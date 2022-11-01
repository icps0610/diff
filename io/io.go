package io

import (
    "encoding/json"
    "fmt"
    "github.com/malexdev/utfutil"
    "io/ioutil"
    "strings"

    "diff/config"
    "diff/script"
)

func Readfile(path string) []config.Data {
    content, _ := utfutil.ReadFile(path, utfutil.UTF8)
    var lines []config.Data
    for idx, line := range strings.Split(string(content), "\n") {
        line = strings.TrimSpace(line)
        if len(line) > 0 {
            lines = append(lines, config.Data{idx, line})
        }
    }
    return script.ArrUniq(lines)
}

func ReadJson(path string) []config.Data {
    file := readContent(path)
    var reaJson []config.Data
    err := json.Unmarshal([]byte(file), &reaJson)
    printError(err)
    return reaJson
}
func ReadJsonC(path string) []config.Comp {
    file := readContent(path)
    var reaJson []config.Comp
    err := json.Unmarshal([]byte(file), &reaJson)
    printError(err)
    return reaJson
}

func readContent(path string) string {
    dat, err := ioutil.ReadFile(path)
    printError(err)
    return string(dat)
}

func SaveJson(datas interface{}, path string) {
    file, err := json.MarshalIndent(datas, "", " ")
    printError(err)
    err = ioutil.WriteFile(path, file, 0777)
    printError(err)
}

func saveLines(lines []string, path string) {
    data := []byte(strings.Join(lines, "\n"))
    ioutil.WriteFile(path, data, 0777)
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
