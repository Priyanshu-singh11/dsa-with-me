let arr1 = [4,6,89,9,0]
let arr2 = [5,8,5,7,8]
let arr3 = [2,4,5,6,7,]


function commonNumberArray(...arr){
 let newArr = [];
 let isFound = false
 for(let i=0; i<arr.length;i++){
  for(let j=0; j<arr[i].length;j++){
  for(let k=0;k<newArr.length;k++){
  
   if(newArr[k]===arr[i][j]){
    isFound = true;
    break
   }
  }
  
   if(!isFound){
    //console.log(arr[i][j])
    newArr.push(arr[i][j])
 }
 }
 }
 return newArr
}
let res = commonNumberArray(arr1,arr2,arr3)
console.log(res)
