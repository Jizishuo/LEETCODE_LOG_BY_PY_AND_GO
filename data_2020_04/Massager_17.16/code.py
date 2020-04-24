class Solution:
    def massage(self, nums: List[int]) -> int:
        
        nuw, last =0, 0
        for no in nums:
            last, now = now, max(last+no, now)

        return now