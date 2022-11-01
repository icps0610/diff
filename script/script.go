package script

import (
    "diff/config"
)

func ArrCompare(arr1, arr2 []string) ([]config.Data, []config.Data, []config.Data, []string) {
    var diff1, diff2, all []config.Data
    var same []string
    diff1, all, same = compare(arr1, arr2, all, "1")
    diff2, all, _ = compare(arr2, arr1, all, "2")
    BubbleSort(all)
    return diff1, diff2, all, same
}

func compare(arr1, arr2 []string, all []config.Data, compare string) ([]config.Data, []config.Data, []string) {
    var diff []config.Data
    var same []string

    for idx, ele := range arr1 {
        if arrIncludeIndex(arr2, ele) == -1 {
            if compare == "1" {
                ele1, ele2 := ele, ""
            } else {
                ele1, ele2 := "", ele
            }
            data := config.Data{idx + 1, ele1, ele2}
            all = append(all, data)
            diff = append(diff, data)
        } else {
            same = append(same, ele)
        }
    }
    return diff, all, same
}

// func CheckKB(datas []config.Data) []config.Data {
//     var narr []config.Data
//     for _, e := range datas {
//         keyword := `support\.microsoft\.com\/\?kbid=(\d+)`
//         scans := Scans(e.Ele1, keyword)

//         fmt.Println(scans)

//         if len(scans) > 0 {
//             c := fmt.Sprintf("wusa /uninstall /kb:%s /quiet /norestar", scans[0])
//             narr = append(narr, config.Data{e.Idx, c, ""})
//         } else {
//             narr = append(narr, e)
//         }
//     }

//     return narr
// }
