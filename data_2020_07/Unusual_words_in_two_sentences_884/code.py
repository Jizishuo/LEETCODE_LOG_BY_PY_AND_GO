class Solution:
    def uncommonFromSentences(self, A: str, B: str) -> List[str]:
        AA = A.split(" ")
        BB = B.split(" ")
        d = {}
        All = AA+BB
        for i in All:
            t = d.get(i)
            if t == None:
                d[i]=1
            else:
                d[i] += 1

        res = []
        for k, v in dd.items():
            if v == 1:
                res.append(k)
        return res

