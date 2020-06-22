class Solution:
    def cuttingRope(self, n: int) -> int:
        if n==2:return 1
        if n==3:return 2
        dp =[0]*(n+1)
        dp[2]=2
        dp[3]=3
        for i in range(4,n+1):
            ii=i-1
            j=1
            ans=0
            while ii>=j:
                if dp[ii]*j>ans:
                    ans=dp[ii]*j
                ii-=1
                j+=1
            dp[i]=ans
        return dp[n]

