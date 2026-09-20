
func maxProfit(prices []int) int {
    var maxProfit,bestBuy = 0, prices[0]
    for i:=0;i<len(prices);i++{
        if prices[i]<bestBuy{
           bestBuy = prices[i]
        }
        profit := prices[i]-bestBuy
        if profit>maxProfit{
            maxProfit = profit
        }
    }
    return maxProfit
}
