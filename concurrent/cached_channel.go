package main

/*
	 两种方式判断通道是否被关闭
		1.i, ok := <-ch1 //通道关闭后再取值时ok = false
		2.for range ch2  //通道关闭后会退出for range循环
*/
//func main() {
//	ch1 := make(chan int)
//	ch2 := make(chan int)
//
//	// 开启goroutine将0~20的数发送到ch1中
//	go func() {
//		for i := 0; i < 100; i++ {
//			ch1 <- i
//		}
//		close(ch1)
//	}()
//
//	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
//	go func() {
//		for {
//			i, ok := <-ch1 //通道关闭后再取值时ok = false
//			if !ok {
//				break
//			}
//			ch2 <- i * i
//		}
//		close(ch2)
//	}()
//
//	for i := range ch2 {
//		// 通道关闭后会退出for range循环
//		fmt.Println(i)
//	}
//}
