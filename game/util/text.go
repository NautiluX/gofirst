package util

import (
	"bytes"
	"image/color"
	"log"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

var (
	font *text.GoTextFaceSource
)

func DrawWrappedText(screen *ebiten.Image, str string, x, y, maxWidth int, maxHeight int, fontSize int, clr color.Color) {
	words := strings.Fields(str)
	textImage := ebiten.NewImage(maxWidth, maxHeight)
	var line string
	yOffset := 0
	lineSpacing := 4
	textWidth := 0.0
	f := text.GoTextFace{
		Source: font,
		Size:   float64(fontSize),
	}
	var testW, testH float64
	for _, w := range words {
		// Try adding the next word to the line, do line break if too wide
		testLine := line + w + " "
		testW, testH = text.Measure(testLine, &f, float64(lineSpacing))
		if testW > textWidth {
			textWidth = testW
		}
		if int(testW) > int(maxWidth) {
			line = w + " "
			yOffset += int(testH) + lineSpacing
		} else {
			line = testLine
		}
		op := &text.DrawOptions{}
		op.GeoM.Translate(float64(0), float64(yOffset))
		op.ColorScale.ScaleWithColor(clr)
		text.Draw(textImage, line, &f, op)
	}

	// Center the text block vertically
	textY := float64(y) + float64((maxHeight-(int(testH)+lineSpacing+yOffset))/2)
	centerTextOpt := &ebiten.DrawImageOptions{}
	centerTextOpt.GeoM.Translate(float64(x), textY)
	screen.DrawImage(textImage, centerTextOpt)
}

func init() {
	var err error
	font, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	if err != nil {
		log.Fatal(err)
	}
}
