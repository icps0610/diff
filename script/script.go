package script

import (
    "fmt"

    "diff/config"
)

func ArrCompare(arr1, arr2 []config.Data) ([]config.Data, []config.Data, []config.Data, []config.Comp) {
    var diff1, diff2, same []config.Data
    var tmp map[int]map[int]string

    for _, e1 := range arr1 {
        idx := arrIncludeIndex(arr2, e1.Str)
        if idx != -1 {
            same = append(same, e1)
        } else {
            diff1 = append(diff1, e1)
            tmp[e1.Idx] = map[int]string{}
            tmp[e1.Idx][0] = e1.Str
        }
    }
    for _, e2 := range arr2 {
        idx := arrIncludeIndex(arr1, e2.Str)
        if idx == -1 {
            diff2 = append(diff2, e2)
            if tmp[e2.Idx] == nil {
                tmp[e2.Idx] = map[int]string{}
            }
            tmp[e2.Idx][1] = e2.Str
        }
    }
    var all []config.Comp
    for i, v := range tmp {
        all = append(all, config.Comp{k, v[0], v[1]})
    }
    bubbleSort(all)
    return diff1, diff2, same, all
}

func CheckKB(arr []config.Data) []config.Data {
    var narr []config.Data
    for _, e := range arr {
        kb := getKB(e.Str)
        if kb == "" {
            narr = append(narr, e)
        } else {
            c := fmt.Sprintf("wusa /uninstall /kb:%s /quiet /norestar", kb)
            narr = append(narr, config.Data{e.Idx, c})
        }
    }
    return narr
}

func getKB(str string) string {
    keyword := `http:\/\/support\.microsoft\.com\/\?kbid=(\d+)`

    match := Scans(str, keyword)
    if len(match) > 0 {
        return match[1]
    }
    return ""
}
