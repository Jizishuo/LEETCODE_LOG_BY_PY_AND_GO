class Solution:
    def halvesAreAlike(self, s: str) -> bool:
        c = ['a','e','i','o','u','A','E','I','O','U']
        a_l = 0
        b_l = 0
        for i in s[:len(s)//2]:
            if i in c:
                a_l+=1
        for i in s[len(s)//2:]:
            if i in c:
                b_l +=1
        return a_l==b_l