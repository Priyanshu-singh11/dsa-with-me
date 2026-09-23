let n = [3,-4,5,4,-1,7,-8]
let maxElem = 0
for(let sta=0;sta<n.length;sta++){
 let currElem = 0
 for(let end=sta;end<n.length;end++){
  currElem += n[end]
  maxElem = Math.max(currElem,maxElem)
 }
}

console.log(maxElem)


