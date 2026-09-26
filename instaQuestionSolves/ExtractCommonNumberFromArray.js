let arr1 = [4,6,89,9,0]
let arr2 = [5,8,5,7,8]


function commonNumberArray(...arr){
 let newArr = [];
 for(let i=0;i<arr.length;i++){
  for(let j=0;j<arr[i].length;j++){
   let isHave = false
   for(let k=0;k<newArr.length;k++){
    if(newArr[k]===arr[i][j]){
     isHave = true
    }
   }
   if(!isHave){
    newArr.push(arr[i][j])
   }
  }
 }
 
 return newArr
}
let res = commonNumberArray(arr1,arr2,arr3)
console.log(res)
