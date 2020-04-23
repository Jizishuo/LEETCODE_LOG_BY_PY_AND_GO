class Solution:
    def countCharacters(self, words, chars: str) -> int:


        sum = 0
        for i in words:
            for j in i:
                if i.count(j) <= chars.count(j):
                    a = 1
                    continue
                else:
                    a = 0
                    break
            if a ==1:
                sum+=len(i)
        return sum

