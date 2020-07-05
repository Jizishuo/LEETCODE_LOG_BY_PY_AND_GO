class Solution:
    def wordPattern(self, pattern: str, str: str) -> bool:
        s = str.split(' ')
        if len(s) != len(pattern):
            return False
        d = {}
        for i, x in enumerate(s):
            if pattern[i] not in d:
                if x in d.values():
                    return False
                d[pattern[i]] = x
            else:
                if x != d[pattern[i]]:
                    return False
        return True