package script

import (
    "diff/config"
)

func ArrCompare(arr1, arr2 []string) ([]config.Data, []config.Data, []config.Data, []string) {
    var diff1, diff2, all []config.Data
    var same []string

    for idx, ele := range arr1 {
        if arrIncludeIndex(arr2, ele) == -1 {
            data := config.Data{idx + 1, ele, ""}
            all = append(all, data)
            diff1 = append(diff1, data)
        } else {
            same = append(same, ele)
        }
    }
    for idx, ele := range arr2 {
        if arrIncludeIndex(arr1, ele) == -1 {
            data := config.Data{idx + 1, "", ele}
            all = append(all, data)
            diff2 = append(diff2, data)
        }
    }

    BubbleSort(all)
    return diff1, diff2, all, same
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
