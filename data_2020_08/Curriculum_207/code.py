class Solution:
    def canFinish(self, numCourses, prerequisites) -> bool:
        # 2, [[1,0],[0,1]]
        # 0 表示没有访问过（白）
        # 1 表示访问过了（黑）
        visited = [0]*numCourses
        adjacency = [[] for _ in range(numCourses)]
        def dfs(i):
            if visited[i] == 1:
                return False
            visited[i] =1
            for j in adjacency[i]:
                if not dfs(j):
                    return False
            visited[i] = 0
            return True
        for cur, pre in prerequisites:
            adjacency[cur].append(pre)

        for i in range(numCourses):
            if not dfs(i):
                return False
        return True
