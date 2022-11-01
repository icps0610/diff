package script

import (
    "encoding/base64"
    "fmt"
    "path"
    "regexp"
    "strconv"

    "diff/config"
)

func ArrUniq(arr []config.Data) []config.Data {
    var n []config.Data
    for _, e := range arr {
        if arrIncludeIndex(n, e.Str) == -1 {
            n = append(n, e)
        }
    }
    return n
}
func arrIncludeIndex(arr []config.Data, str string) int {
    for i, e := range arr {
        if str == e.Str {
            return i
        }
    }
    return -1
}
func arrRemoveByIndex(arr []config.Data, idx int) []config.Data {
    arr[idx] = arr[len(arr)-1]
    return arr[:len(arr)-1]
}

func bubbleSort(arr []config.Comp) {
    size := len(arr)
    for i := 0; i < size; i++ {
        for j := 1; j < size; j++ {
            oi := to_i(arr[j].Idx)
            ti := to_i(arr[j-1].Idx)
            if oi < ti {
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
func to_i(s string) int {
    i, _ := strconv.Atoi(s)
    return i
}

func to_s(i int) string {
    return strconv.Itoa(i)
}

func EnBase64(str string) string {
    return base64.StdEncoding.EncodeToString([]byte(str))
}

func DeBase64(str string) string {
    s, err := base64.StdEncoding.DecodeString(str)
    printError(err)
    return string(s)
}

func BaseName(str string) string {
    return path.Base(str)
}

func printError(err error) {
    if err != nil {
        fmt.Println(err)
    }
}
