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

//export DebugGoGC
func DebugGoGC() {
	go func() {
		for {
			p := make([]int, 1024*1024)
			p[0] = 1
			time.Sleep(1 * time.Second)
		}
	}()
}

//export GetGoMap
func GetGoMap() map[int]int {
	return make(map[int]int)
}

func main() {}
