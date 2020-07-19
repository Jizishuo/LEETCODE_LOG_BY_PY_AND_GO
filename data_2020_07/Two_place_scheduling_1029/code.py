class Solution:
    def twoCitySchedCost(self, costs: List[List[int]]) -> int:
        costs.sort(key=lambda x : x[0]-x[1])

        res = 0
        n = len(costs)

        res += sum([i[0] for i in costs[:n//2]])  # 前半部分去A地
        res += sum([i[1] for i in costs[n//2:]])

        return res