
 func removeStars(s string) string {
    var newStr = []byte{}
    for i:=0;i<len(s);i++{
        if s[i]=='*'{
            if len(newStr)>0{
                newStr = newStr[:len(newStr)-1]
            }
        }else{
            newStr = append(newStr,s[i])
        }
    }
    return string(newStr)
}
