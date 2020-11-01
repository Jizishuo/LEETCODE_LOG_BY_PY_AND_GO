class Solution:
    def destCity(self, paths: List[List[str]]) -> str:
        a = set()
        b = set()
        for p in paths:
            a.add(p[0])
            a.add(p[1])
            b.add(p[0])
        return (a - b).pop()
