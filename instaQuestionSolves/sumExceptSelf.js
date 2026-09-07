let arr = [1,4,5,6,7]


function sumExceptSelf(arr){
 let newArr = []
 for(let i=0;i<arr.length;i++){
  let sum = 0;
    for(let j=0;j<arr.length;j++){
      if(i!==j){
       sum += arr[j]
      }
    }
    newArr.push(sum)
}
return newArr
}

let result=sumExceptSelf(arr)
console.log(result)



function OptimizedSumExceptSelf(arr){
 let total = arr.reduce((sum,elem)=>{
  return sum += elem
 })
 
 return arr.map((elem)=>total-elem)
 }

let result=OptimizedSumExceptSelf(arr)
console.log(result)
