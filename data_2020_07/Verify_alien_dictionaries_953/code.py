class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        # return words == sorted(words, key=lambda w: [order.index(x) for x in w])

        for i in range(len(words) - 1):
            if self.compare(words[i], words[i + 1], order):
                return False
        return True

    def compare(self, x, y, order):
        for i in range(len(x)):
            if i >= len(y):
                return True
            if order.index(x[i]) < order.index(y[i]):
                return False
            if order.index(x[i]) > order.index(y[i]):
                return True
        return False

