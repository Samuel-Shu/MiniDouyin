package utils

import (
	"fmt"
	"testing"
)

func TestGetVideo(t *testing.T) {
	privateUrl := GetVideo("ceshi.mp4")
	fmt.Println(privateUrl)
}
