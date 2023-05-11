package diff

import (
    "diff/config"
    "diff/script"
)

func Run(arr1, arr2 []string) ([]config.Data, []config.Data, []config.Data, []config.Data) {
    var diff1, diff2, allDiff, same []config.Data

    diff1, same = compare(arr1, arr2, true)
    diff2, _ = compare(arr2, arr1, false)

    allDiff = append(allDiff, diff1...)
    allDiff = append(allDiff, diff2...)

    script.QuickSort(allDiff)
    return diff1, diff2, allDiff, same
}

func compare(arr1, arr2 []string, compareOrd bool) ([]config.Data, []config.Data) {
    var diff, same []config.Data

    // 找重複
    arr2Map := make(map[string]struct{}, len(arr2))
    for _, ele := range arr2 {
        arr2Map[ele] = struct{}{}
    }

    for idx, element := range arr1 {
        if _, ok := arr2Map[element]; !ok {
            var element1, element2 string
            if compareOrd {
                element1 = element
            } else {
                element2 = element
            }
            data := config.Data{idx + 1, element1, element2}
            diff = append(diff, data)
        } else {
            data := config.Data{idx + 1, element, element}
            same = append(same, data)
        }
    }
    return diff, same
}
