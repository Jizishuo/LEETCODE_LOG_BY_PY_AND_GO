class Solution:
    def maxProfit(self, prices) -> int:
        res = 0
        for i in range(len(prices)-1):
            print(prices[i], prices[i+1])
            if prices[i]<prices[i+1]:
                print("--", prices[i], prices[i + 1])
                res = res + prices[i]-prices[i+1]
        return res

Solution().maxProfit([7,1,5,3,6,4])