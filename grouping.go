//Algorithm no 3
package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scanln(&s)
	point := strings.Count(s, ".")
	if point > 1 {
	 s = strings.Replace(s, ".", "", -1)
	}
	for i := 0; i < len(s); i++ {
		tmp := s[i:i+1]
		fmt.Println(tmp + strings.Repeat("0", len(s) - i - 1))
	}
}

