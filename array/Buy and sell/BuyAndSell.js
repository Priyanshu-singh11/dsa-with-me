let price = [7,1,5,3,6,4];

let bestBuy = price[0];
let maxProfit = 0;

for(let i=0;i<price.length;i++){
 if(bestBuy>price[i]){
  bestBuy = price[i]
 }
 profit = price[i]-bestBuy
 console.log(profit)
 if(profit>maxProfit){
  maxProfit = profit
 }
}
console.log(maxProfit)
