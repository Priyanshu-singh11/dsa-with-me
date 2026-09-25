
let arr = [3,-4,5,4,-1,7,-8]
let currElem = 0;
let maxElem = -Infinity

for(let i=0;i<arr.length;i++){
 currElem += arr[i]
 if(currElem>maxElem){
  maxElem = currElem
 }
 if(currElem<0){
  currElem = 0
 }
}
console.log(maxElem)
