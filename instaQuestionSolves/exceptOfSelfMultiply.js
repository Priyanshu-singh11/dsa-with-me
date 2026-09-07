let arr = [1,4,5,6,7]
function OptimizedSumExceptSelf(arr){
 let total = arr.reduce((sum,elem)=>{
  return sum *= elem
 })
 
 return arr.map((elem)=>total/elem)
 }

let result=OptimizedSumExceptSelf(arr)
console.log(result)
