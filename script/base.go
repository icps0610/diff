package script

import (
    "encoding/base64"
    "fmt"
    "regexp"

    "diff/config"
)

func arrIncludeIndex(arr []string, str string) int {
    for i, e := range arr {
        if str == e {
            return i
        }
    }
    return -1
}

func BubbleSort(arr []config.Data) {
    size := len(arr)
    for i := 0; i < size; i++ {
        for j := 1; j < size; j++ {
            if arr[j].Idx < arr[j-1].Idx {
                arr[j], arr[j-1] = arr[j-1], arr[j]
            }
        }
    }
}

func Scans(keyword, str string) []string {
    re := regexp.MustCompile(keyword)
    match := re.FindStringSubmatch(str)
    if len(match) > 0 {
        return match[1:]
    }
    return match
}

func EnBase64(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

func DeBase64(str string) string {
    s, err := base64.StdEncoding.DecodeString(str)
    printError(err)
    return string(s)
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
