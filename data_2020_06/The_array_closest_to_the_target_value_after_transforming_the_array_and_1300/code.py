import math

class Solution:
    def findBestValue(self, arr, target: int) -> int:
        # print(arr, target)
        maxx = sum(arr)
        if maxx <= target:
            return max(arr)
        mixx = math.ceil(target/len(arr))
        # mixx = math.floor(target/len(arr))
        # mixx = target//len(arr)
        print(mixx)
        mixx_arr = min(arr)
        if mixx_arr > mixx:
            return mixx

        arr.remove(mixx_arr)
        return self.findBestValue(arr, target-mixx_arr)
        # summ = sum(arr)
        # if summ <= target:
        #     return max(arr)
        # l = len(arr)
        # val = target // l
        # summ, last = val * l, 0
        # while summ < target:
        #     last = summ
        #     summ = 0
        #     for i in range(l):
        #         summ += arr[i] if val > arr[i] else val
        #     val += 1
        # return val - 2 if abs(target - summ) >= abs(target - last) else val - 1




S = Solution().findBestValue([60864,25176,27249,21296,20204], 56803)
# S = Solution().findBestValue([4,9,3], 10)

# S = Solution().findBestValue([48772,52931,14253,32289,75263], 40876)
print(S)

# aa = [4, 9, 3]
# aa.remove(3)
# print(aa)