package main

import (
	"golang.org/x/time/rate"
)


//这个包 主要来测试目前常用的限流算法 令牌桶、漏桶

func main()  {
	
	rate.NewLimiter()

}


