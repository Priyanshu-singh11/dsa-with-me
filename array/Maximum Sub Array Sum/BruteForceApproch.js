
let n = [3,-4,5,4,-1,7,-8]
let maxElem = -Infinity;
let possibleValuesOfMax = []
let startIdx = 0
let endIdx = 0
let maxSubArray = []
for(let sta=0;sta<n.length;sta++){
 let currElem = 0
 let val = ""
 for(let end=sta;end<n.length;end++){
  currElem += n[end]
  if(currElem>maxElem){
   possibleValuesOfMax.push(currElem)
   startIdx = sta
   endIdx = end
   maxElem = currElem
  }
  //maxElem = Math.max(currElem,maxElem)
 }
}
maxSubArray.push(...n.slice(startIdx,endIdx+1))
console.log(maxElem)
console.log(possibleValuesOfMax)
console.log(maxSubArray)
