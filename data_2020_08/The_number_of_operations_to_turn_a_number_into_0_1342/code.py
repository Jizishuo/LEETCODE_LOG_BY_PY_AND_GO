class Solution:
    def numberOfSteps (self, num: int) -> int:

        for ans in itertools.count():
            if not num:
                return ans
            num = num -1 if num & 1 else num >>1
