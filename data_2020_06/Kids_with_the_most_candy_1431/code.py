class Solution:
    def kidsWithCandies(self, candies: List[int], extraCandies: int) -> List[bool]:
        t = max(candies)
        res = []

        for i in candies:
            if i+extraCandies>=t:
                res.append(True)
            else:res.append(False)
        return res