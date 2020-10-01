class Solution:
    def countPrimeSetBits(self, L: int, R: int) -> int:
        z = [0,0,1,1,0,1,0,1,0,0,0,1,0,1,0,0,0,1,0,1,0,0,0,1]
        n = 0
        for i in range(L, R+1):
            q =bin(i)
            s =q.count('1')
            if z[s]==1:
                n+=1
        return n
