class Solution:
    def uniqueOccurrences(self, arr: List[int]) -> bool:
        m = {}
        for i in arr:
            if i in m:
                d[i] +=1
            else:
                d[i]=1
        return len(d.values())==len(set(d.values()))