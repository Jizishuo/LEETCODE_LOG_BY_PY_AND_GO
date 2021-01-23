class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        n, l = len(mat), 0
        for i in range(n):
            l += mat[i][i]
            mat[i][i] = 0
            l += mat[i][n-i-1]
        return l