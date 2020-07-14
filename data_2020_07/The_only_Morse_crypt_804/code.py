class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        l1 = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]

        letter = 'abcdefghijklmnopqrstuvwxyz'
        d = {}
        for i, v in enumerate(letter):
            d[v]=l1[i]

        res = []
        for v in words:
            s = ''
            for w in v:
                s += d[w]
            res.append(s)
        return len(set(res))