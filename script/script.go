package script

import (
    "diff/config"
)

func ArrCompare(arr1, arr2 []string) ([]config.Data, []config.Data, []config.Data, []string) {
    var diff1, diff2, allDiff []config.Data
    var same []string

    diff1, same = compare(arr1, arr2, true)
    diff2, _ = compare(arr2, arr1, false)

    allDiff = append(allDiff, diff1...)
    allDiff = append(allDiff, diff2...)

    QuickSort(allDiff)
    return diff1, diff2, allDiff, same
}

func compare(arr1, arr2 []string, compareOrd bool) ([]config.Data, []string) {
    diff := make([]config.Data, 0, len(arr1))
    same := make([]string, 0, len(arr1))

    // 找重複
    arr2Map := make(map[string]struct{}, len(arr2))
    for _, ele := range arr2 {
        arr2Map[ele] = struct{}{}
    }

    for idx, ele := range arr1 {
        if _, ok := arr2Map[ele]; !ok {
            var ele1, ele2 string
            if compareOrd {
                ele1 = ele
            } else {
                ele2 = ele
            }
            data := config.Data{idx + 1, ele1, ele2}
            diff = append(diff, data)
        } else {
            same = append(same, ele)
        }
    }
    return diff, same
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

// func Scans(keyword, str string) []string {
//     re := regexp.MustCompile(keyword)
//     match := re.FindStringSubmatch(str)
//     if len(match) > 0 {
//         return match[1:]
//     }
//     return match
// }
