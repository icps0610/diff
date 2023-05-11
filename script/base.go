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

func QuickSort(arr []config.Data) {
    if len(arr) <= 1 {
        return
    }

    left, right := 0, len(arr)-1

    for i, pivot := 1, arr[0]; i <= right; {
        switch {
        case arr[i].Idx < pivot.Idx:
            left++
            arr[i], arr[left] = arr[left], arr[i]
            i++
        case arr[i].Idx > pivot.Idx:
            arr[i], arr[right] = arr[right], arr[i]
            right--
        default:
            i++
        }
    }

    arr[0], arr[left] = arr[left], arr[0]

    QuickSort(arr[:left])
    QuickSort(arr[left+1:])
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
