class Solution:
    def minTimeToVisitAllPoints(self, points: List[List[int]]) -> int:
        res = 0
        for j in range(1, len(points)):
            times0 = abs(points[j][0] - points[j - 1][0])
            times1 = abs(points[j][1] - points[j - 1][1])
            distance = times1 - times0 if times1 > times0 else times0 - times1
            res += min(times0, times1) + distance
        return res


