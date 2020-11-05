class Solution:
    def minCount(self, coins: List[int]) -> int:
        if not  coins:
            return 0
        num = 0
        for i in range(len(coins)):
            num +=(coins[i]//2)
            if coins[i]%2!=0:
                num +=1
        return num