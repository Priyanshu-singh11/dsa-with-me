let arr = [1,4,5,6,7]
function OptimizedMultiplyExceptSelf(arr){
 let total = arr.reduce((sum,elem)=>{
  return sum *= elem
 })
 
 return arr.map((elem)=>total/elem)
 }

let result=OptimizedMultiplyExceptSelf(arr)
console.log(result)

let arr = [1,4,5,6,7]


function multiplyExceptSelf(arr){
 let newArr = []
 for(let i=0;i<arr.length;i++){
  let sum = 1;
    for(let j=0;j<arr.length;j++){
      if(i!==j){
       sum *= arr[j]
       //console.log(sum,arr[j])
      }
    }
    newArr.push(sum)
}
return newArr
}

let res=multiplyExceptSelf(arr)
console.log(res)
