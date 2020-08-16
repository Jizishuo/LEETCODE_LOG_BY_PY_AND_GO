
from collections import deque

class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:
        row, col = len(image), len(image[0])
        p = image[sr][sc]
        deq = deque([(sr, sc)])
        if p == newColor:
            return image
        while deq:
            x, y = deq.popleft()
            if x in {-1, row} or y in {-1, col} or image[x][y] != p:
                continue
            image[x][y] = newColor
            deq.extend([(x - 1, y), (x, y - 1), (x + 1, y), (x, y + 1)])
        return image


