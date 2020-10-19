class Solution:
    def game(self, guess: List[int], answer: List[int]) -> int:
        res = 0
        for i in range(len(guess)):
            if guess[i] == answer[i]:
                res +=1
        return res