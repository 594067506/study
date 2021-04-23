package learn

import "fmt"

//切片实现的线性表结构
type  Stable struct {
	Length  int  //数组目前元素的个数
	Array   []int
}


func  InitTable()  {
	var st Stable
	st.Array = []int{}
	st.Array = append(st.Array,7)
	st.Array = append(st.Array,7)
	st.Array = append(st.Array,7)
	st.Array = append(st.Array,7)
	st.Array = append(st.Array,9)
	st.Length = 5


	fmt.Println(st.Array)
	index:= BinarySearch(&st,8)

	fmt.Println("index===",index)
}

func BinarySearch( st *Stable, key int ) int{
	var  low ,height,mid int
	low = 0; height = st.Length-1
	for low <=height{
		mid =  (low+height)/2
		if key==st.Array[mid] {
			//mid正好是是要查找的值,检查下是不是最右
			fmt.Println("CheckRight",mid)
			return CheckRight(st,mid)
		}else if key>st.Array[mid] {
			low = mid+1;
		}else{
			height = mid-1
		}
	}
	//走到这里说明没有查到,返回要插入的位置
	return low
}

//查到到的元素是不是最右位置
func CheckRight( st *Stable,index int) int {
	//从当前位置的后一个元素依次遍历当出现相同的元素位置找到
	var  key_word = st.Array[index]
	var i int
	for i =index+1 ; i<=st.Length-1; i++  {
		if key_word != st.Array[i]{
			return i-1
		}
	}
	//说明 不满足循环条件这种情况说明最后一个位置就是我们要找的最右位置
	return  i-1
}

