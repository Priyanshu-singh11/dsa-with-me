
function maxProductOfTwoElementOptimize(arr) {
  let maxProduct = 1
  let largest = arr[0]
  let secondLargest = arr[0]
  for(let i=0;i<arr.length;i++){
   if(arr[i]>largest){
    secondLargest = largest
    largest = arr[i]
   }
  }
  maxProduct = (secondLargest)*(largest)
  return maxProduct
 }
 
 let ans = maxProductOfTwoElementOptimize(arr)
 console.log(ans)
