package main

import (
	"fmt"
	"strconv"
	"time"
)

func timeDemo()  {
	now := time.Now().Format("2006-01-02 15:04")
	parseInt, _ := strconv.ParseInt(now, 10, 32)
	fmt.Println(parseInt)
}

