class Twitter:

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.users = dict()  # key:value分别对应用户 userId(int) 和 其关注者(list)
        self.posts = []  # 发布的帖子，每个元素格式为 [userId, tweetId]



    def postTweet(self, userId: int, tweetId: int) -> None:
        """
        Compose a new tweet.
        """
        self.posts.append([userId, tweetId])
        if not self.users.get(userId):
            self.users[userId] = []



    def getNewsFeed(self, userId: int) -> List[int]:
        """
        Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent.
        """
        if userId not in self.users.keys():
            return
        else:
            ids = [userId] + self.users.get(userId)
            tmp = []
            count = 10
            for post in self.posts[-1:-(len(self.posts)+1):-1]:
                if count > 0:
                    if post[0] in ids:
                        tmp.append(post[1])
                        count -= 1
            return tmp

    def follow(self, followerId: int, followeeId: int) -> None:
        """
        Follower follows a followee. If the operation is invalid, it should be a no-op.
        """
        if followerId not in self.users.keys():
            self.users[followerId] = []
            self.users[followerId].append(followeeId)
        else:
            self.users[followerId].append(followeeId)


    def unfollow(self, followerId: int, followeeId: int) -> None:
        """
        Follower unfollows a followee. If the operation is invalid, it should be a no-op.
        """
        if followerId not in self.users.keys():
            return
        else:
            if followeeId in self.users[followerId]:
                self.users[followerId].remove(followeeId)
            else:
                return

# Your Twitter object will be instantiated and called as such:
# obj = Twitter()
# obj.postTweet(userId,tweetId)
# param_2 = obj.getNewsFeed(userId)
# obj.follow(followerId,followeeId)
# obj.unfollow(followerId,followeeId)