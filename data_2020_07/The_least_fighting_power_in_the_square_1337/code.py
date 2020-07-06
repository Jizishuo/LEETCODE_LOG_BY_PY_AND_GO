class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        # n = len(mat)
        # pow = [sum(line) for line in mat]
        # ans = list(range(n))
        # ans.sort(key=lambda i: (pow[i],i))
        # return ans[:k]

        result = [(sum(mat[i]), i) for i in range(len(mat))]
        result.sort()
        return [i[1] for i in result[:k]]





