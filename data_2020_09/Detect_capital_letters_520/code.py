class Solution:
    def detectCapitalUse(self, word: str) -> bool:
        count = 0
        for i in word:
            if 65 <= ord(i) <= 90:
                count +=1
        if count ==0:
            return True
        elif count == len(word):
            return True
        elif count == 1 and 65<=ord(word[0])<=90:
            return True
        else:
            return False