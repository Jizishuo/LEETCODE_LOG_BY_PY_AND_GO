class Solution:
    def findTheDistanceValue(self, arr1: List[int], arr2: List[int], d: int) -> int:
        res = 0
        for i in arr1:
            for ii in arr2:
                if abs(i-ii) <= d:
                    break
            else:
                res +=1
        return res
