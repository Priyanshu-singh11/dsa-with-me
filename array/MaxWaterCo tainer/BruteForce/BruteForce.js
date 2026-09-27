function  maximumWater(cont) {
 let maxWater = 0
 let leftContIdx = 0
 let rightContIdx = 0
 for(let i=0;i<cont.length;i++){
  for(let j=i+1;j<cont.length;j++){
   let width = j-i
   let height;
   if(cont[i]>cont[j]){
    height = cont[j]
    //leftCont = cont[i]
    //rightCont = cont[j]
   }else{
    height = cont[i]
    
   }
   if(maxWater<height*width){
    maxWater = height*width
    leftContIdx = i
    rightContIdx = j
   }
  }
 }
 console.log(leftContIdx,rightContIdx)
 console.log(cont[leftContIdx,rightContIdx],cont[rightContIdx])
 return maxWater
}

let cont = [1,8,6,2,5,4,8,3,7]
let res = maximumWater(cont)
console.log("maxWater ",res)

