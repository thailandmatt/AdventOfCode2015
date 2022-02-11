package main

import (
	"crypto/md5"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	key := "yzbqklnj"

	i := 0
	for {
		data := []byte(key + strconv.Itoa(i))
		test := fmt.Sprintf("%x", md5.Sum(data))
		if strings.HasPrefix(test, "000000") { //change to 5 or 6 zeros for part 1 or 2
			fmt.Printf("Number:%v, hash:%v", i, test)
			break
		}
		i++
	}
}
