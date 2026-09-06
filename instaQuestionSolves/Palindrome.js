function palindromeChecker(str) {
    let n = str.length-1;
    let i = 0
    for(n;n>=0;n--){
        if(str[i]===str[n]){
            return true
        }else{
            return false
        }
        i++;
    }
}
let ans = palindromeChecker("madam")
console.log(ans)
