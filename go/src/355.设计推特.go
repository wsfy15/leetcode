/*
 * @lc app=leetcode.cn id=355 lang=golang
 *
 * [355] 设计推特
 */

// @lc code=start
type tweet struct {
	userId int
	id int
	time int
}

type Twitter struct {
	time int
	followers map[int]*list.List // 用户的粉丝
	tweets map[int]*list.List // 每个用户看到的推文
	selfTweet map[int][]*tweet // 每个用户自己的推文
}


/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		followers: make(map[int]*list.List),
		tweets:    make(map[int]*list.List),
		selfTweet: make(map[int][]*tweet),
	}
}


/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int)  {
	// 加入自己的推特列表
	if _, ok := this.tweets[userId]; !ok {
		this.tweets[userId] = list.New()
	}
	this.tweets[userId].PushFront(&tweet{
		userId: userId,
		id:     tweetId,
		time: this.time,
	})
	this.selfTweet[userId] = append(this.selfTweet[userId], &tweet{
		userId: userId,
		id:     tweetId,
		time: this.time,
	})

	// 加入粉丝的推特列表
	if lists, ok := this.followers[userId]; ok {
		for l := lists.Front(); l != nil; l = l.Next() {
			id := l.Value.(int)
			this.tweets[id].PushFront(&tweet{
				userId: userId,
				id:     tweetId,
				time: this.time,
			})
		}
	}

	this.time++
}


/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	res := []int{}
	if list, ok := this.tweets[userId]; ok {
		i := 0
		for l := list.Front(); l != nil && i < 10; l = l.Next() {
			res = append(res, l.Value.(*tweet).id)
			i++
		}
	}
	return res
}


/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int)  {
	this.checkExist(followeeId)
	if followeeId == followerId {
		return
	}
	this.checkExist(followerId)

	// 检查是否已在粉丝列表里
	list := this.followers[followeeId]
	for l := list.Front(); l != nil; l = l.Next() {
		if l.Value.(int) == followerId {
			return
		}
	}

	this.followers[followeeId].PushFront(followerId)

	list = this.tweets[followerId]
	newTweet := this.selfTweet[followeeId]
	index := len(newTweet) - 1
	if index < 0 {
		return
	}

	for l := list.Front(); l != nil; {
		if l.Value.(*tweet).time < newTweet[index].time {
			list.InsertBefore(&tweet{
				userId: newTweet[index].userId,
				id:     newTweet[index].id,
				time:   newTweet[index].time,
			}, l)
			index--

			if index < 0 {
				break
			}
		} else {
			l = l.Next()
		}
	}

	for index >= 0 {
		list.PushBack(&tweet{
			userId: newTweet[index].userId,
			id:     newTweet[index].id,
			time:   newTweet[index].time,
		})
		index--
	}
}


/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
	this.checkExist(followeeId)
	if followeeId == followerId {
		return
	}
	this.checkExist(followerId)

	list := this.followers[followeeId]
	for l := list.Front(); l != nil; l = l.Next() {
		if l.Value.(int) == followerId {
			list.Remove(l)
			break
		}
	}

	// 删除推文
	list = this.tweets[followerId]
	for l := list.Front(); l != nil; {
		tmp := l.Next()
		if l.Value.(*tweet).userId == followeeId {
			list.Remove(l)
		}
		l = tmp
	}
}

func (this *Twitter) checkExist(userId int) {
	if _, ok := this.followers[userId]; !ok {
		this.followers[userId] = list.New()
	}
	if _, ok := this.tweets[userId]; !ok {
		this.tweets[userId] = list.New()
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
// @lc code=end

