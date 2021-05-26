package tests

import (
	"fmt"
	"study/learn"
	"testing"
)

//单元测试
func TestBinarySearch(t *testing.T)  {

	type params struct {
		st learn.Stable
		want int
	}

	sts:=[]params{
		{
			st: learn.Stable{
				Length: 5,
				Array: []int{7,7,7,7,9},
			},
			want: 4,
		},
		{
			st: learn.Stable{
				Length: 5,
				Array: []int{7,7,8,7,9},
			},
			want: 2,
		},
	}

	for  i,v:=range sts {
		fmt.Printf("%p :",&sts[i].st)
		got:= learn.BinarySearch(&sts[i].st,8)
		if got !=v.want {
			t.Errorf(" got =%d ,want=%d" ,got,v.want)
		}
	}

}


//性能测试
func BenchmarkBinarySearch(b *testing.B)  {
	st:=learn.Stable{
		Length: 5,
		Array: []int{7,7,8,7,9},
	}

	for i:=0;i<b.N ;i++{
		learn.BinarySearch(&st,8)
	}

}