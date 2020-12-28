package main

import (
	"github.com/algorithms/constant"
	"github.com/algorithms/rb"
	"image/gif"
	"math/rand"
)

func main() {
	count := 20
	constant.Anim = gif.GIF{LoopCount: 50}
	root := rb.Insert(nil, 13)
	nums := make(map[int]int)
	for i := 0; i < count; i++ {
		num := rand.Intn(100) + 1
		if nums[num] > 0 || num == 13 {
			continue
		} else {
			nums[num] = num
			root = rb.Insert(root, num)
		}
	}
	rb.GenerateGif(&constant.Anim, "rb.gif")
	rb.GeneratePng(rb.CreatePngImg(root), "rb.png")
}
