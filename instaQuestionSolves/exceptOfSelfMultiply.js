let arr = [1,4,5,6,7]
function OptimizedMultiplyExceptSelf(arr){
 let total = arr.reduce((sum,elem)=>{
  return sum *= elem
 })
 
 return arr.map((elem)=>total/elem)
 }

let result=OptimizedMultiplyExceptSelf(arr)
console.log(result)
