function reverseString(str) {
    let n = str.length-1;
    let newStr = "";
    for(n;n>=0;n--){
        newStr += str[n]
    }
    return newStr
}
let ans = reverseString("tihbohs");
console.log(ans)
