function flatArray(list){
    let newArr = [];
    for (let item of list){
        if(Array.isArray(item)){
            newArr.push(...flatArray(item))
        }else{
            
            newArr.push(item)
        }
    }
    return newArr
}
let result = flatArray(arr)
console.log(result)
