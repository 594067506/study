package tests

import (
	"fmt"
	"hash/fnv"
)


func hash(s string) uint32 {
	h := fnv.New32a()
	h.Write([]byte(s))
	return h.Sum32()
}


func Initfun() {
	fmt.Println(hash("HelloWorBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBbbbbbajljdkljlka立刻就爱上看来大家快来加奥克兰BBBBBBBBBBBBBBBBBBld"))
	fmt.Println(hash("你"))
}
