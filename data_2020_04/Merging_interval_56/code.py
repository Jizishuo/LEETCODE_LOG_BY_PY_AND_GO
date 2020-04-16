class Solution:
    def merge(self, intervals):
        intervals.sort(key=lambda x:x[0])
        res = []
        for v in intervals:
            if not res or res[-1][1] < v[0]:
                res.append(v)
            else:
                res[-1][1] = max(res[-1][1], v[1])


        return res



asd = Solution().merge([[1,3],[2,6],[8,10],[15,18]])

# print([1,2][-1])