package main

import "fmt"

//#include <stdio.h>
//void callC() {
//	printf("hello c go\n");
//}
import "C"

func main() {
	fmt.Println("hello go")
	C.callC()
}
