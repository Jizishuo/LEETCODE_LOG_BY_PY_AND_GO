class Solution:
    def isLongPressedName(self, name: str, typed: str) -> bool:
        m,n = len(name), len(typed)
        i=j=0
        while j<n:
            if i<m and name[i] == typed[j]:
                i +=1
            elif j==0 or typed[j-1] !=typed[j]:
                return False
            j+=1
        return i==m

