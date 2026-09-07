let str = "hello this is javaScript"
function captalize(str,index){
 let newStr = "";
 for(let i=0;i<=str.length-1;i++){
  if(index===i){
   newStr += str[i].toUpperCase()
  }else{
   newStr += str[i]
  }
 }
 return newStr
}
console.log(captalize(str,10))







function OptimizeCaptialize(str,word){
 let newStr = "";
 for (let elem of str) {
  if(elem===word){
   newStr += elem.toUpperCase()
  }
  else{
   newStr += elem
  }
 }

 return newStr
}
