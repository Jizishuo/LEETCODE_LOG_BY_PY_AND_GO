class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:

        num_s = ""
        for i in digits:
            num_s = num_s + str(i)

        nnn = int(num_s) + 1

        res = []
        for ii in str(nnn):
            res.append(int(ii))
        return res