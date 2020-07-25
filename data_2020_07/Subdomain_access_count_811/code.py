class Solution:
    def subdomainVisits(self, cpdomains: List[str]) -> List[str]:
        count = {}
        result = []
        number = 0
        for cp in cpdomains:
            for i in range(len(cp)):
                if cp[i] == ' ':
                    number = int(cp[:i])
                    if cp[i + 1:] not in count.keys():
                        count[cp[i + 1:]] = number
                    else:
                        count[cp[i + 1:]] = count[cp[i + 1:]] + number
                if cp[i] == '.':
                    if cp[i + 1:] not in count.keys():
                        count[cp[i + 1:]] = number
                    else:
                        count[cp[i + 1:]] = count[cp[i + 1:]] + number
            number = 0
        for k, v in count.items():
            result.append(str(v) + ' ' + k)
        return result

