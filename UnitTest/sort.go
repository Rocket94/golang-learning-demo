package UnitTest

func BubblingSort(a []int) {
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}

//func SliceEqual(a,b[]int)bool{
//	if a==nil{
//		return b==nil
//	}
//	if b==nil {
//		return a==nil
//	}
//	if len(a)!= len(b){
//		return false
//	}
//	for i,x:=range a{
//		if x!=b[i]{
//			return false
//		}
//	}
//	return true
//}
