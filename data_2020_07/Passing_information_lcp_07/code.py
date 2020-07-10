class Solution:
    def numWays(self, n: int, relation: List[List[int]], k: int) -> int:
        # d = [[]for _ in range(n)]
        # for i, j in relation:
        #     d[i].append(j)
        # q = [0]
        # for i in range(k):
        #     q = [j for i in q for j in d[i]]
        # return q.count(n-1)
        d = [1] + [0] * (n - 1)
        for _ in range(k):
            t = [0] * n
            for i, j in relation:
                t[j] += d[i]
            d = t
        return d[-1]

