function removeDuplicate(data){
    if(data.length===0){
        return 0
    }
    let newArr = []
    
    for(let i=0;i<=data.length-1;i++){
        let isFound = false;
        for(let j=0;j<=newArr.length-1;j++){
        if(data[i]===newArr[j]){
            isFound = true;
            break
        }
    }
    
    if(!isFound){
        newArr.push(data[i])
    }
    }
    return newArr
}
let arr = [2,4,3,6,6,6,2,2,2]
let result = removeDuplicate(arr)
console.log(result)
