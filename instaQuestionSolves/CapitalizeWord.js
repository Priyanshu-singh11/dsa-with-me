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
