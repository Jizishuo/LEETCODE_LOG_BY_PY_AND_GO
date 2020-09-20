class Solution:
    def relativeSortArray(self, arr1: List[int], arr2: List[int]) -> List[int]:
        bins = [0 for _ in range(1001)]
        res = []
        for i in arr1:
            bins[i] += 1
        for i in arr2:
            res += [i] * bins[i]
            bins[i] = 0
        for i in range(len(bins)):
            res += [i] * bins[i]
        return res
