class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        news = sorted(nums)
        l = []
        for num in nums:
            l.append(news.index(num))
        return l