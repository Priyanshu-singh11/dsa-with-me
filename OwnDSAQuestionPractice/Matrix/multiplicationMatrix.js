function multiplyMatrix(arr1,arr2){
 
 let tempArr = []
 for(let i=0;i<arr1.length;i++){
  let matrix = [];
  for(let j=0;j<arr1[i].length;j++){
   let add = arr1[i][j]*arr2[i][j]
    matrix.push(add)
  }
  tempArr.push(matrix)
 }
 return tempArr
}

let arr1 = [
 [7,8,9],
 [6,7,8],
 [1,4,6]
]
let arr2 = [
 [7,8,9],
 [6,7,8],
 [1,4,6]
]

let result = multiplyMatrix(arr1,arr2)
console.log(result)
