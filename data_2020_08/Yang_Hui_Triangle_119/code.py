class Solution:
    def getRow(self, rowIndex: int) -> List[int]:
        res = [1]*(rowIndex+1)
        for i in range(2, rowIndex+1):
            for j in range(1, i):
                index = i-j
                res[index] = res[index]+res[index-1]
        return res
