
class Solution:
    def isUnique(self, astr: str) -> bool:
        return len(astr) == len(set(astr))