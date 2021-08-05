package main

/*
	给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

	设计一个算法来计算你所能获取的最大利润。你可以尽可能地完成更多的交易（多次买卖一支股票）。

	注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。

	输入: [7,1,5,3,6,4]  即使知道第二天股价比第一天低，但是因为最初没有股票，所以也卖出。但是第三天比第二天的股价高，所以第二天就得购买股票获取利润
	输出: 7

	输入: [1,2,3,4,5]  股价一直在涨，如果第二购买股票，那么第三天必须卖出，因为不能连续购买股票，只能一次买入，等股票涨到最高点卖出
	输出: 4

	输入: [7,6,4,3,1]  股价一直在跌，就不会买入股票，又因为最开始就没买入股票，所以后面也就没有卖出的操作
	输出: 0
*/
func maxProfit(prices []int) int {
	sum := 0
	for i := 0; i < len(prices); i++ {
		if prices[i] <= prices[i+1] {
			sum = prices[len(prices)-1] - prices[0]
		} else if prices[i] > prices[i+1] {
			return 0
		} else if prices[0] > prices[1] {
			sum += prices[2*i+2] - prices[2*i+1]
		}
	}
	return sum
} 

func main() {

}
