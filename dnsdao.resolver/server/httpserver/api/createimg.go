package api

import (
	"fmt"
	"github.com/golang/freetype"
	"golang.org/x/image/font"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"time"
	"unicode"
)

func GenerateRandInt(min, max int) int {
	rand.Seed(time.Now().Unix()) //随机种子
	return rand.Intn(max-min) + min
}
func rotate90(m image.Image) image.Image {
	rotate90 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			//  设置像素点
			rotate90.Set(m.Bounds().Max.Y-x, y, m.At(y, x))
		}
	}
	return rotate90
}

// 旋转180度
func rotate180(m image.Image) image.Image {
	rotate180 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dx(), m.Bounds().Dy()))
	// 矩阵旋转
	for x := m.Bounds().Min.X; x < m.Bounds().Max.X; x++ {
		for y := m.Bounds().Min.Y; y < m.Bounds().Max.Y; y++ {
			//  设置像素点
			rotate180.Set(m.Bounds().Max.X-x, m.Bounds().Max.Y-y, m.At(x, y))
		}
	}
	return rotate180
}

// 旋转270度
func rotate270(m image.Image) image.Image {
	rotate270 := image.NewRGBA(image.Rect(0, 0, m.Bounds().Dy(), m.Bounds().Dx()))
	// 矩阵旋转
	for x := m.Bounds().Min.Y; x < m.Bounds().Max.Y; x++ {
		for y := m.Bounds().Max.X - 1; y >= m.Bounds().Min.X; y-- {
			// 设置像素点
			rotate270.Set(x, m.Bounds().Max.X-y, m.At(y, x))
		}
	}
	return rotate270

}
func writeLabel(img image.Image, label string, x, y int, fontColor color.Color, size float64, fontPath string) (image.Image, error) {
	bound := img.Bounds()
	// 创建一个新的图片
	rgba := image.NewRGBA(image.Rect(0, 0, bound.Dx(), bound.Dy()))
	// 读取字体
	fontBytes, err := ioutil.ReadFile(fontPath)
	if err != nil {
		return rgba, err
	}
	myFont, err := freetype.ParseFont(fontBytes)
	if err != nil {
		return rgba, err
	}

	draw.Draw(rgba, rgba.Bounds(), img, bound.Min, draw.Src)
	c := freetype.NewContext()
	c.SetDPI(72)
	c.SetFont(myFont)
	c.SetFontSize(size)
	c.SetClip(rgba.Bounds())
	c.SetDst(rgba)
	uni := image.NewUniform(fontColor)
	c.SetSrc(uni)
	c.SetHinting(font.HintingNone)

	// 在指定的位置显示
	pt := freetype.Pt(x, y+int(c.PointToFixed(size)>>6))
	if _, err := c.DrawString(label, pt); err != nil {
		return rgba, err
	}

	return rgba, nil
}
func UpperNumber(name string) int {
	count := 0
	for _, s := range name {
		if unicode.IsUpper(s) {
			count++
		}
	}
	return count
}
func CreateImage(name string, nametype string, createImg string) error {
	number := GenerateRandInt(1, 16)
	src, err := os.Open(fmt.Sprintf("./%s/%d.png", nametype, number))
	if err != nil {
		log.Println(err)
		return err
	}
	img, err := png.Decode(src)
	//img, err := gif.Decode(src)
	if err != nil {
		log.Println(err)
		return err
	}
	count := UpperNumber(name)
	labellen := len(name)*10 + count*2
	x := (270 - labellen) / 2
	y := 130
	y1 := 0
	if x < 0 {
		sp := len(name) / 2
		label1 := name[0:sp]
		label2 := name[sp:]
		labellen1 := len(label1)*10 + count*2
		labellen2 := len(label2)*10 + count*2
		x1 := (270 - labellen1) / 2
		x2 := (270 - labellen2) / 2
		x = (x1 + x2) / 2
		y1 = y + 20
		writeimage, err := writeLabel(img, label1, x, y, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		writeimage, err = writeLabel(writeimage, label2, x, y1, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		if err != nil {
			log.Println(err)
			return err
		}
		f, err := os.Create(createImg)
		if err != nil {
			log.Println(err)
			return err
		}
		defer func() { _ = f.Close() }()

		err = png.Encode(f, writeimage)
		if err != nil {
			log.Println(err)
			return err
		}
	} else {
		writeimage, err := writeLabel(img, name, x, y, color.RGBA{255, 255, 255, 255}, 18, "./PingFang Heavy.ttf")
		if err != nil {
			log.Println(err)
			return err
		}
		f, err := os.Create(createImg)
		if err != nil {
			log.Println(err)
			return err
		}
		defer func() { _ = f.Close() }()

		err = png.Encode(f, writeimage)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}
