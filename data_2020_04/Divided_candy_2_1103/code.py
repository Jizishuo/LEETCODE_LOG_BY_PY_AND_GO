class Solution:
    def distributeCandies(self, candies: int, num_people: int) -> List[int]:
        i = 0
        j = 1
        a = [0]*num_people
        while candies> 0:
            a[i] += j
            candies -= j
            if i == num_people-1:
                i = -1
            if candies <=j:
                a[i+1] += candies
                break
            i += 1
            j += 1
        return a
