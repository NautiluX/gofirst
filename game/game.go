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
	Width        float32
	Height       float32
	Rules        []Rule
	Card         Rule
	SwapCard     bool
	SwapDistance float32
	History      []Rule
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

	// next card when screen is touched
	touchIds := []ebiten.TouchID{}
	touchIds = inpututil.AppendJustReleasedTouchIDs(touchIds)
	if len(touchIds) > 0 {
		g.SelectCard()
	}

	// next card when screen is tapped/clicked
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.SelectCard()
	}
	return nil
}

const (
	ImageWidth  float32 = 180
	ImageHeight float32 = 180
)

// Select random card
func (g *Game) SelectCard() {
	if g.SwapCard {
		return
	}

	// Append current card to history
	g.History = append(g.History, g.Card)

	g.Card = g.Rules[rand.IntN(len(g.Rules))]
	for g.Card.Title == g.History[len(g.History)-1].Title {
		g.Card = g.Rules[rand.IntN(len(g.Rules))]
	}
	if len(g.History) > 1 {
		g.SwapCard = true
	}
	g.SwapDistance = 0
	// Only keep last 10 history items
	if len(g.History) > 10 {
		g.History = g.History[len(g.History)-10:]
	}
}

func (g *Game) Draw(screen *ebiten.Image) {

	currentCard := ebiten.NewImage(int(g.Width), int(g.Height))

	g.drawCard(currentCard, g.Card)

	screen.DrawImage(currentCard, &ebiten.DrawImageOptions{})

	if g.SwapCard {
		flipAnimation := ebiten.DrawImageOptions{}
		flipAnimation.GeoM.Translate(-float64(g.SwapDistance), 0)
		oldCard := ebiten.NewImage(int(g.Width), int(g.Height))
		g.drawCard(oldCard, g.History[len(g.History)-1])
		screen.DrawImage(oldCard, &flipAnimation)
		if g.SwapDistance < g.Width {
			g.SwapDistance += 7
		}
		if g.SwapDistance >= g.Width {
			g.SwapCard = false
		}
	}
}

func (g *Game) drawCard(image *ebiten.Image, card Rule) {
	//image.Fill(color.RGBA{20, 35, 25, 255})

	cardWidth := g.Width - 3
	cardHeight := g.Height - 3

	// Draw shadow
	vector.FillRect(image, 3, 3, cardWidth, cardHeight, color.RGBA{0, 0, 0, 20}, false)
	// Use an "Off-White" or "Parchment" color
	vector.FillRect(image, 0, 0, cardWidth, cardHeight, color.RGBA{252, 245, 220, 255}, false)

	// 5. Draw Card Border (Beveled look)
	vector.StrokeRect(image, 0, 0, cardWidth, cardHeight, 2, color.RGBA{40, 30, 20, 20}, false)

	// 6. Internal Content Layout
	contentX := 10

	textColor := color.RGBA{30, 20, 10, 255}
	util.DrawWrappedText(image, card.Title, contentX, 0, int(cardWidth), 40, 10, textColor)

	// Image Square (Centered inside card)
	imgX := +(cardWidth - ImageWidth) / 2
	imgY := float32(40)
	vector.FillRect(image, imgX, imgY, ImageWidth, ImageHeight, color.RGBA{220, 210, 190, 255}, false)

	if card.HasImage() {
		imOp := &ebiten.DrawImageOptions{}
		imOp.GeoM.Translate(float64(imgX), float64(imgY))
		image.DrawImage(card.GetSprite(), imOp)
	}
	vector.StrokeRect(image, imgX, imgY, ImageWidth, ImageHeight, 1, color.RGBA{100, 90, 80, 255}, false)

	ruleDescY := imgY + ImageHeight
	util.DrawWrappedText(image, card.Description, contentX, int(ruleDescY), int(cardWidth), int(cardHeight-ruleDescY), 10, textColor)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(g.Width), int(g.Height)
}
