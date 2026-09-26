
let name = "Hello"
let repStr = []
for(let i=0;i<name.length;i++){
 let isHave = false
 for(let j=0;j<repStr.length;j++){
  
  if(name[i]===repStr[j]){
   isHave = true
   break
  }
 }
 if(!isHave){
   repStr.push(name[i])
  }
}
console.log(repStr)
