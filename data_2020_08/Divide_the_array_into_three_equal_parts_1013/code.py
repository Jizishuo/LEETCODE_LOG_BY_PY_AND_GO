class Solution:
    def canThreePartsEqualSum(self, A: List[int]) -> bool:
        mid = sum(A)/3
        if mid %1 ==0:
            num = 0
            l = []
            for i in range(len(A)):
                num += A[i]
                if num == mid:
                    l.append(num)
                    num = 0
                if len(l)==2 and i != len(A)-1:
                    last = sum(A[i+1:])
                    l.append(last)
                    return l[0]==l[1]==l[2]
        return False