package main

import "fmt"

func main(){
	fmt.Println("AAAA");
	fmt.Println("BBBB");

	var c string = "CCCCC";
	fmt.Println(c);

	var a, b int = 12, 7;

	var sum int = a + b;
	fmt.Println("a + b = ",sum);

	var d float64 = 12.0;
	var e float64 = 4.0;

	var f float64 = d / e;
	fmt.Println("d / e = ", f);
}