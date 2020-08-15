package main

func main(){
	ch:=make(chan int)

	var writeCh chan<-int=ch//只能写不能读
	var readCh <-chan int=ch//只能写不能读

}