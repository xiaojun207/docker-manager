package model

import (
	"log"
	"math"
	"time"
)

/**
 *	这是一个简单的登录锁，避免被暴力破解密码
 *	错误3次以上，按错误次数，指数增加等待时间
 */
type UserLock struct {
	LockNum     SyncMap
	lastTime    SyncMap
	MaxErrorNum int
}

func NewUserLock() UserLock {
	res := UserLock{}
	res.Init()
	return res
}

func (e *UserLock) Init() {
	e.LockNum = SyncMap{}
	e.lastTime = SyncMap{}
	e.MaxErrorNum = 3
}

func (e *UserLock) AddLockNum(username string) {
	num, _ := e.LockNum.LoadInt(username)
	log.Println("num:", num)
	e.LockNum.Store(username, num+1)
	e.lastTime.Store(username, time.Now().Unix())
}

func (e *UserLock) ClearLockNum(username string) {
	e.LockNum.Delete(username)
	e.lastTime.Delete(username)
}

func (e *UserLock) Locked(username string) bool {
	num, _ := e.LockNum.LoadInt(username)
	if num >= e.MaxErrorNum {
		lastTime, _ := e.lastTime.LoadInt64(username)
		// 如果最后一次错误时间，加上等待时间，已经过了，则可以登录
		return time.Now().Unix() < lastTime+(int64)(math.Pow(3, float64(num)))
	}
	return false
}
