package rb

import (
	"github.com/algorithms/constant"
	graphics "github.com/algorithms/graphics"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
	"image"
	"image/color"
	"image/draw"
	"image/gif"
	"image/png"
	"math"
	"os"
	"strconv"
)

type TreeNode struct {
	Val    int
	Red    bool
	LChild bool
	Left   *TreeNode
	Right  *TreeNode
	Parent *TreeNode
}

func Insert(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val, Red: false}
		CreateGifImg(root, &constant.Anim)
		return root
	}
	p := root
	for p != nil {
		if val < p.Val {
			if p.Left == nil {
				p.Left = &TreeNode{Val: val, Red: true, LChild: true, Parent: p}
				CreateGifImg(getRoot(p), &constant.Anim)
				fixup(p.Left)
				break
			}
			p = p.Left
		} else {
			if p.Right == nil {
				p.Right = &TreeNode{Val: val, Red: true, LChild: false, Parent: p}
				CreateGifImg(getRoot(p), &constant.Anim)
				fixup(p.Right)
				break
			}
			p = p.Right
		}
	}
	return getRoot(p)
}

func fixup(c *TreeNode) {
	for c.Parent != nil && c.Parent.Red {
		if c.Parent.LChild {
			uncle := c.Parent.Parent.Right
			if uncle != nil && uncle.Red {
				c.Parent.Red = false
				uncle.Red = false
				c.Parent.Parent.Red = true
				c = c.Parent.Parent
				CreateGifImg(getRoot(c), &constant.Anim)
			} else {
				if !c.LChild {
					c = c.Parent
					rotateLeft(c)
					CreateGifImg(getRoot(c), &constant.Anim)
				}
				c.Parent.Red = false
				c.Parent.Parent.Red = true
				CreateGifImg(getRoot(c), &constant.Anim)
				rotateRight(c.Parent.Parent)
				CreateGifImg(getRoot(c), &constant.Anim)
			}
		} else {
			uncle := c.Parent.Parent.Left
			if uncle != nil && uncle.Red {
				c.Parent.Red = false
				uncle.Red = false
				c.Parent.Parent.Red = true
				c = c.Parent.Parent
				CreateGifImg(getRoot(c), &constant.Anim)
			} else {
				if c.LChild {
					c = c.Parent
					rotateRight(c)
					CreateGifImg(getRoot(c), &constant.Anim)
				}
				c.Parent.Red = false
				c.Parent.Parent.Red = true
				CreateGifImg(getRoot(c), &constant.Anim)
				rotateLeft(c.Parent.Parent)
				CreateGifImg(getRoot(c), &constant.Anim)
			}
		}
	}
	if c.Parent == nil {
		//到达根节点，涂为黑色
		c.Red = false
		CreateGifImg(getRoot(c), &constant.Anim)
	}
}

func getRoot(current *TreeNode) *TreeNode {
	for current.Parent != nil {
		current = current.Parent
	}
	return current
}

func rotateLeft(c *TreeNode) {
	r := c.Right
	if r == nil {
		return
	}
	c.Right = r.Left
	if r.Left != nil {
		r.Left.LChild = false
		r.Left.Parent = c
	}
	r.Parent = c.Parent
	if c.Parent != nil {
		if c.LChild {
			c.Parent.Left = r
			r.LChild = true
		} else {
			c.Parent.Right = r
		}
	}
	r.Left = c
	c.Parent = r
	c.LChild = true
}

func rotateRight(c *TreeNode) {
	l := c.Left
	if l == nil {
		return
	}
	c.Left = l.Right
	if l.Right != nil {
		l.Right.LChild = true
		l.Right.Parent = c
	}
	l.Parent = c.Parent
	if c.Parent != nil {
		if c.LChild {
			c.Parent.Left = l
		} else {
			c.Parent.Right = l
			l.LChild = false
		}
	}
	l.Right = c
	c.Parent = l
	c.LChild = false
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return int(1 + math.Max(float64(depth(root.Left)), float64(depth(root.Right))))
}

/**
图形化
*/

const (
	width  = 1440
	height = 900
	r      = 30
)

var palette = []color.Color{color.White, color.Black, color.RGBA{255, 0, 0, 255}} //调色板

func showTree(root *TreeNode, row, depth, x, y, r int, left bool, img draw.Image) {
	if root == nil {
		return
	}
	var xx, yy int
	if root.Val != 0 {
		if row == 0 {
			xx = x
			yy = r
		} else {
			if left {
				xx = x - 2*depth/row*r
			} else {
				xx = x + 2*depth/row*r
			}
			yy = r + 3*row*r
		}
		var col color.Color
		if root.Red {
			col = palette[2]
		} else {
			col = palette[1]
		}
		graphics.Circle(image.Point{xx, yy}, r, root.Red, img, col)
		graphics.Line(x, y, xx, yy, img, palette[1])
		showNum(xx, yy, root.Val, img)
	}
	if root.Left != nil {
		showTree(root.Left, row+1, depth, xx, yy, r, true, img)
	}
	if root.Right != nil {
		showTree(root.Right, row+1, depth, xx, yy, r, false, img)
	}
}

func showNum(x, y, number int, img draw.Image) {
	colFont := color.RGBA{255, 255, 255, 255}
	point := fixed.Point26_6{fixed.Int26_6((x - 6) * 64), fixed.Int26_6((y + 6) * 64)}
	d := &font.Drawer{img, image.NewUniform(colFont), basicfont.Face7x13, point}
	d.DrawString(strconv.Itoa(number))
}

func CreatePngImg(root *TreeNode) *image.NRGBA {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	showTree(root, 0, depth(root), width/2, r, r, false, img)
	return img
}

func GeneratePng(img draw.Image, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	png.Encode(f, img)
}

func CreateGifImg(root *TreeNode, anim *gif.GIF) {
	img := image.NewPaletted(image.Rect(0, 0, width, height), palette)
	showTree(root, 0, depth(root), width/2, r, r, false, img)
	anim.Delay = append(anim.Delay, 200)
	anim.Image = append(anim.Image, img)
}

func GenerateGif(anim *gif.GIF, filename string) {
	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	gif.EncodeAll(f, anim)
}
