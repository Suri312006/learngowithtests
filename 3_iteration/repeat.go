package iteration

import "strings"

func Repeat(in string, count int) string {
    // declaration of string variable only (explicit version)
    // var repeated string
    // for i := 0; i < count; i++{
    //     repeated += in
    // }

    return strings.Repeat(in, count)
    

    //return repeated
}
