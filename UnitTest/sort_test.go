package UnitTest

import (
	"testing"
)

type sortCase struct {
	Source []int
	Expected []int
}

func TestBubblingSort(t *testing.T) {
	createBubblingSortTestCase(t,&sortCase{[]int{1, 2, 3},[]int{1,2,3}})
	createBubblingSortTestCase(t,&sortCase{[]int{},[]int{}})
	createBubblingSortTestCase(t,&sortCase{[]int{0,1,0,0},[]int{0,0,0,1}})
	createBubblingSortTestCase(t,&sortCase{[]int{-1,-9,-5,9,5,1},[]int{-9,-5,-1,1,5,9}})
}

func createBubblingSortTestCase(t *testing.T, s *sortCase) {
	//帮助定位调用函数的位置
	t.Helper()
	temp:=s.Source
	BubblingSort(temp)
	if  !SliceEqual(temp,s.Expected) {
		t.Fatalf("slice %d sort result expected %d, but %d got",
			s.Source,s.Expected, temp)
	}
}

func SliceEqual(a,b[]interface{})bool{
	if a==nil{
		return b==nil
	}
	if b==nil {
		return a==nil
	}
	if len(a)!= len(b){
		return false
	}
	for i,x:=range a{
		if x!=b[i]{
			return false
		}
	}
	return true
}