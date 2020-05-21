class Solution:
    def replaceSpace(self, s: str) -> str:
        # return s.replace(" ", "%20")
        res = []
        for i in s:
            if i == ' ':
                res.append("%20")
            else:
                res.append(i)
        return "".join(res)