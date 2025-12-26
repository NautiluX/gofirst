package game

import (
	"image/color"
	"math/rand/v2"

	"github.com/NautiluX/gofirst/game/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	Width   float32
	Height  float32
	Rules   []Rule
	Card    Rule
	History []Rule
}

type Rule struct {
	Title       string
	Description string
	ImagePath   string
}

func NewGame(width, height int) *Game {
	g := &Game{
		Width:  float32(width),
		Height: float32(height),
		Rules:  GetRules(),
	}
	g.SelectCard()
	return g
}

func (g *Game) Update() error {
	// quit on ESC
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) ||
		inpututil.IsKeyJustPressed(ebiten.KeyQ) {
		return ebiten.Termination
	}

	// select new card on SPACE
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		g.SelectCard()
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		if len(g.History) > 1 {
			g.Card = g.History[len(g.History)-1]
			g.History = g.History[:len(g.History)-1]
		}
	}
	return nil
}

const (
	ImageWidth  float32 = 180
	ImageHeight float32 = 180
)

// Select random card
func (g *Game) SelectCard() {
	g.History = append(g.History, g.Card)
	g.Card = g.Rules[rand.IntN(len(g.Rules))]
	// Only keep last 10 history items
	if len(g.History) > 10 {
		g.History = g.History[len(g.History)-10:]
	}
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{20, 35, 25, 255})

	cardWidth := g.Width - 3
	cardHeight := g.Height - 3

	// Draw shadow
	vector.FillRect(screen, 3, 3, cardWidth, cardHeight, color.RGBA{0, 0, 0, 150}, false)
	// Use an "Off-White" or "Parchment" color
	vector.FillRect(screen, 0, 0, cardWidth, cardHeight, color.RGBA{252, 245, 220, 255}, false)

	// 5. Draw Card Border (Beveled look)
	vector.StrokeRect(screen, 0, 0, cardWidth, cardHeight, 2, color.RGBA{40, 30, 20, 255}, false)

	// 6. Internal Content Layout
	contentX := 10

	textColor := color.RGBA{30, 20, 10, 255}
	util.DrawWrappedText(screen, g.Card.Title, contentX, 10, int(cardWidth), 10, textColor)

	// Image Square (Centered inside card)
	imgX := +(cardWidth - ImageWidth) / 2
	imgY := float32(30)
	vector.FillRect(screen, imgX, imgY, ImageWidth, ImageHeight, color.RGBA{220, 210, 190, 255}, false)

	if g.Card.HasImage() {
		imOp := &ebiten.DrawImageOptions{}
		imOp.GeoM.Translate(float64(imgX), float64(imgY))
		screen.DrawImage(g.Card.GetSprite(), imOp)
	}
	vector.StrokeRect(screen, imgX, imgY, ImageWidth, ImageHeight, 1, color.RGBA{100, 90, 80, 255}, false)

	ruleDescY := imgY + ImageHeight + 15
	util.DrawWrappedText(screen, g.Card.Description, contentX, int(ruleDescY), int(cardWidth), 10, textColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.Width), int(g.Height)
}
