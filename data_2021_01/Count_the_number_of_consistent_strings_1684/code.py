class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        res = 0
        al = set(list(allowed))
        for w in words:
            wl = set(list(w))
            if al >= wl:
                res +=1
        return res
