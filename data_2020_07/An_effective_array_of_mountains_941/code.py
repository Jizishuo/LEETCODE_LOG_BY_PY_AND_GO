class Solution:
    def validMountainArray(self, A: List[int]) -> bool:
        up, down = False, False

        for i in range(1, len(A)-1):
            if (A[i]-A[i-1])>0:
                if down:
                    return False
                up = True
            elif(A[i]-A[i-1])<0:
                if not up:
                    return False
                down = True
            else:
                return False
        return up and down
