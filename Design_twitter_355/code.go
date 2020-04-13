package Design_twitter_355


type Twitter struct {
	user map[int][]int
	twi [][2]int
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		user : map[int][]int{},
		twi : [][2]int{},
	}
}


/** Compose a new tweet. */
func (t *Twitter) PostTweet(userId int, tweetId int)  {
	if len(t.user[userId]) == 0 {
		t.user[userId] = append(t.user[userId], userId)
	}
	temp := [2]int{userId, tweetId}
	t.twi = append(t.twi, temp)
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (t *Twitter) GetNewsFeed(userId int) []int {
	if len(t.user[userId]) == 0 {
		t.user[userId] = append(t.user[userId], userId)
	}
	res := []int{}
	count := 0
	for i := len(t.twi)-1; i>=0 && count<10; i -- {
		for _, a := range t.user[userId] {
			if a == t.twi[i][0] {
				res = append(res, t.twi[i][1])
				count ++
			}
		}
	}
	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Follow(followerId int, followeeId int)  {
	if followeeId == followerId {
		return
	}
	if len(t.user[followerId]) == 0 {
		t.user[followerId] = append(t.user[followerId], followerId)
	}
	for _, v := range t.user[followerId] {
		if v == followeeId {
			return
		}
	}
	t.user[followerId] = append(t.user[followerId], followeeId)
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (t *Twitter) Unfollow(followerId int, followeeId int)  {
	if followerId == followeeId || len(t.user[followerId]) == 0 {
		return
	}
	if t.user[followerId][len(t.user[followerId])-1] == followeeId {
		t.user[followerId] = t.user[followerId][:len(t.user[followerId])-1]
	} else {
		for i:=1;i<len(t.user[followerId])-1; i++ {
			if t.user[followerId][i] == followeeId {
				t.user[followerId] = append(t.user[followerId][0:i], t.user[followerId][i+1:]...)
			}
		}
	}
}


/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */