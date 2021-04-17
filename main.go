package main

import (
	"fmt"
	"sync"
	"time"
)

type Bucket struct {
	mu sync.Mutex
	cap int  //令牌容量
	num int //每次加入几个
	ch chan int64
	token int
	times int
}


var num int64
func NewBucket(r int,b int,times int) *Bucket {
	bucket := &Bucket{
		   cap:   b,
		   num:  r,
		   ch:    make(chan int64,b),
		   times: times,
	   }
	 go bucket.startTicker()
     return bucket
}


func (this *Bucket) startTicker() {
	for i := 0;i < cap(this.ch) ; i++  {
		for j := 0; j < this.num ; i++  {
			this.ch <- this.SetToken()
		}
	}
}


//生成token放入桶内
func (this *Bucket)SetToken()int64{
	num = num + 1
	return num
}


//向桶里存入token
func (this *Bucket)Add(){
        this.mu.Lock()
        this.ch <- this.SetToken()
		defer this.mu.Unlock()
}

//读取数据
func (this *Bucket)GetToken(n int){
	this.mu.Lock()
	for i := 0;i < n; i++  {
          str := <-this.ch
          fmt.Println(str)
	}
	this.mu.Unlock()
}


func main()  {
	//NewBucket(2,5,1)
	bu :=NewBucket(2,4,5)
	for{
		bu.GetToken(2)
		fmt.Println(time.Now().Format("2016-04-02 15:04:05"))
		time.Sleep(time.Second)
	}
}
