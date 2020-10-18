class Solution:
    def sumOddLengthSubarrays(self, arr: List[int]) -> int:
        res = 0
        for n in range(1, len(arr)+1, 2):
            for i in range(len(arr)-n+1):
                res += sum(arr[i:i+n])
        return res
