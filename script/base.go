package script

import (
    "encoding/base64"
    "fmt"
)

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
