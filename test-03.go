package main

import (
	"fmt"
)

func main() {
	s :="hello"
	c :=[]byte(s)
	fmt.Println(c)
	  //convert string to []byte type
	  c[0] = 'c'
	  s2 :=string(c)
	   //convert back to string type
	   fmt.Printf(:%s\n , s2)

}