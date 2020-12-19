class Solution:
    def commonChars(self, A: List[str]) -> List[str]:
        # 初始化列表，记录字符出现最少次数的部分
        min_freq = [0] * 26

        # 先记录第一个字符串的字符出现次数
        for ch in A[0]:
            min_freq[ord(ch) - ord('a')] += 1

        # 遍历后续，统计并比较取字符出现次数较小的情况
        for i in range(1, len(A)):
            other_freq = [0] * 26
            for ch in A[i]:
                other_freq[ord(ch) - ord('a')] += 1
            # 与 min_freq 中字符出现次数比对
            for j in range(26):
                min_freq[j] = min(min_freq[j], other_freq[j])

        ans = []
        # 将 min_freq 对应次数的输出字符
        for i in range(len(min_freq)):
            for _ in range(min_freq[i]):
                ans.append(chr(i + ord('a')))

        return ans
