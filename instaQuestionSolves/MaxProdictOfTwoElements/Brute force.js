
function maxProductOfTwoElement(arr){
  let maxProduct = 0
  for(let i=0;i<arr.length;i++){
   for(let j=i+1;j<arr.length;j++){
    let product = (arr[i]-1)*(arr[j]-1)
    if(product>maxProduct){
     maxProduct = product
    }
   }
  }
  return maxProduct
 }
 
 let arr = [3,4,5,2]
 let res = maxProductOfTwoElement(arr)
 console.log(res)
 
