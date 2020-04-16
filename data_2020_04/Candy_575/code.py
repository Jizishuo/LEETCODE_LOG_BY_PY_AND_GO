class Solution:
    def distributeCandies(self, candies: List[int]) -> int:
        each = len(candies) // 2
        candies_set = set(candies)
        return each if each <= len(candies_set) else len(candies_set)

