/**
 * @author  Solomon Ng <solomon.wzs@gmail.com>
 * @version 1.0
 * @date    2019-02-22
 * @license MIT
 */

package main

import "C"

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println("init libgoc")
}

//export GoFoo
func GoFoo() {
	go func() {
		for {
			p := make([]int, 1024*1024)
			p[0] = 1
			time.Sleep(1 * time.Second)
		}
	}()
}

//export GoBar
func GoBar() {
	fmt.Println("Hello world")
}

func main() {}
