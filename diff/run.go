package diff

import (
    . "iDiff/conf"
)

func Run(arr1, arr2 []string) ([]Diff, []Diff, []Diff, []Diff) {
    var diff1, diff2, allDiff, same []Diff

    diff1, same = compare(arr1, arr2, true)
    diff2, _ = compare(arr2, arr1, false)

    allDiff = append(allDiff, diff1...)
    allDiff = append(allDiff, diff2...)

    QuickSort(allDiff)
    return diff1, diff2, allDiff, same
}

func compare(arr1, arr2 []string, compareOrd bool) ([]Diff, []Diff) {
    var diff, same []Diff

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
            data := Diff{idx + 1, element1, element2}
            diff = append(diff, data)
        } else {
            data := Diff{idx + 1, element, element}
            same = append(same, data)
        }
    }
    return diff, same
}

func QuickSort(arr []Diff) {
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
