function maxWaterContainer(cont){
 let maxWater = 0;
 let left = 0
 let right = cont.length-1;
 while(left<right){
  let width = right-left
  let height;
  if(cont[left]>cont[right]){
   height = cont[right]
   right -= 1
  }else{
   height = cont[left]
   left += 1
  }
  let area = width*height
  if(area>maxWater){
   maxWater = area
  }
 }
 return maxWater
}
let res = maxWaterContainer(cont)
console.log(res)
