package script

import (
    "fmt"

    "diff/config"
)

func ArrCompare(a1, a2 []config.Data) ([]config.Data, []config.Data, []config.Data, []config.Comp) {
    var d1, d2, same []config.Data
    var tmp = map[int]map[int]string{}
    for _, e1 := range a1 {
        idx := arrIncludeIndex(a2, e1.Str)
        if idx != -1 {
            same = append(same, e1)
        } else {
            d1 = append(d1, e1)
            tmp[e1.Idx] = map[int]string{}
            tmp[e1.Idx][0] = e1.Str
        }
    }
    for _, e2 := range a2 {
        idx := arrIncludeIndex(a1, e2.Str)
        if idx == -1 {
            d2 = append(d2, e2)
            if tmp[e2.Idx] == nil {
                tmp[e2.Idx] = map[int]string{}
            }
            tmp[e2.Idx][1] = e2.Str
        }
    }
    var all []config.Comp
    for i, v := range tmp {
        k := to_s(i)
        all = append(all, config.Comp{k, v[0], v[1]})
    }
    bubbleSort(all)
    return d1, d2, same, all
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
