//Algorithm no 1
package main

import (
	"fmt"
)

func main(){
	var N int
    fmt.Scanf("%d", &N)
	socks := make([]int, N)
	res := 0
    for i, _ := range socks {
		var tmp int
		fmt.Scanf("%d", &tmp)
		k, found := Find(socks, tmp)
		if found {
			res += 1
			socks[i] = 0
			socks[k] = 0
		} else {
			socks[i] = tmp
		}
	}
	fmt.Println(res)
}

func Find(slice []int, val int) (int, bool) {
    for i, item := range slice {
        if item == val {
            return i, true
        }
    }
    return -1, false
}