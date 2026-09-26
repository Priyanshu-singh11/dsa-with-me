let data = [6,8,18,527,69,8,1]
let newData = []
for(let i=0;i<data.length;i++){
 for(let j=i+1;j<data.length;j++){
  if(data[i]>data[j]){
   [data[j],data[i]] = [data[i],data[j]]
  }
 }
}
console.log(data)
