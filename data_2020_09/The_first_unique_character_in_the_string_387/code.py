class Solution:
    def firstUniqChar(self, s: str) -> int:
        h = {}
        for i in s:
            h[i] = h.get(i, 0) + 1

        for key in h.keys():
            if h[key] == 1:
                return s.find(key)
        return -1