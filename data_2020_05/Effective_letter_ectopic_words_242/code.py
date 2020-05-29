class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        # s_d = {}
        # t_d = {}
        #
        # for i in s:
        #     if i not in s_d:
        #         s_d.update(**{i:1})
        #     else:
        #         s_d[i] +=1
        #
        # for i in t:
        #     if i not in t_d:
        #         t_d.update(**{i:1})
        #     else:
        #         t_d[i] +=1
        #
        # if s_d == t_d:
        #     return True
        # else:
        #     return False

        return sorted(s) == sorted(t)