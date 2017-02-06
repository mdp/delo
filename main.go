package main

import (
	"bufio"
	"flag"
	"fmt"
	"image"
	"os"

	"image/color"

	"github.com/mdp/monochromeoled"
	"github.com/mdp/smallfont"
	"golang.org/x/exp/io/i2c"
)

var height = flag.Int("height", 32, "Pixel height of the OLED")

func main() {
	flag.Parse()
	img := image.NewRGBA(image.Rect(0, 0, 128, *height)) // 52Pi OLED size
	ctx := smallfont.Context{
		Font:  smallfont.Font8x8,
		Dst:   img,
		Color: color.White,
	}

	// Open the i2c device
	// TODO: allow for customizing the address
	d, err := monochromeoled.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x3c, 128, *height)
	if err != nil {
		panic(err)
	}
	defer d.Close()

	// Split the input by newline and print each to a subsequent line
	scanner := bufio.NewScanner(os.Stdin)
	lines, line := *height/ctx.Font.Height, 0
	for scanner.Scan() {
		if line >= lines {
			break
		}
		message := []byte(scanner.Text())
		err := ctx.Draw(message[0:15], 0, line*ctx.Font.Height)
		if err != nil {
			fmt.Println(err)
			break
		}
		line++
	}

	// Clear the display
	if err := d.Clear(); err != nil {
		panic(err)
	}
	// Set the display image to our text
	if err := d.SetImage(0, 0, img); err != nil {
		panic(err)
	}
	if err := d.Draw(); err != nil {
		panic(err)
	}
}
