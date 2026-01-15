package game

import "github.com/hajimehoshi/ebiten/v2"

func GetRules() []Rule {
	return []Rule{
		{
			Title:       "The Rainbow Walker",
			Description: "The player wearing the most colorful socks starts the game.",
			ImagePath:   "assets/rainbow_walker.png",
		},
		{
			Title:       "The Time Traveler",
			Description: "The player who most recently visited another country starts.",
			ImagePath:   "assets/time_traveler.png",
		},
		{
			Title:       "The Early Riser",
			Description: "The player who woke up the earliest this morning starts.",
			ImagePath:   "assets/early_riser.png",
		},
		{
			Title:       "The Bread Winner",
			Description: "The player who most recently ate a piece of bread starts.",
			ImagePath:   "assets/bread_winner.png",
		},
		{
			Title:       "The Birthday VIP",
			Description: "The player whose birthday is closest to today starts.",
			ImagePath:   "assets/birthday_vip.png",
		},
		{
			Title:       "The Tallest Tower",
			Description: "The player whose head is furthest from the floor starts.",
			ImagePath:   "assets/tallest_tower.png",
		},
		{
			Title:       "The Fashion Icon",
			Description: "The player wearing the most buttons on their clothes starts.",
			ImagePath:   "assets/fashion_icon.png",
		},
		{
			Title:       "The Local Hero",
			Description: "The player who lives closest to the current location starts.",
			ImagePath:   "assets/local_hero.png",
		},
		{
			Title:       "The Master Chef",
			Description: "The player who most recently cooked a meal starts.",
			ImagePath:   "assets/master_chef.png",
		},
		{
			Title:       "The Game Buyer",
			Description: "The player who most recently bought a new board game starts.",
			ImagePath:   "assets/game_buyer.png",
		},
		{
			Title:       "The Fresh Recruit",
			Description: "The youngest player at the table starts the game.",
			ImagePath:   "assets/fresh_recruit.png",
		},
		{
			Title:       "The Ancient One",
			Description: "The oldest player at the table starts the game.",
			ImagePath:   "assets/ancient_one.png",
		},
		{
			Title:       "The Trimmed Traveler",
			Description: "The player who most recently visited a barber starts.",
			ImagePath:   "assets/trimmed_traveler.png",
		},
		{
			Title:       "The Hydrated Hero",
			Description: "The player who most recently finished a glass of water starts.",
			ImagePath:   "assets/hydrated_hero.png",
		},
		{
			Title:       "The Key Master",
			Description: "The player with the most physical keys on them starts.",
			ImagePath:   "assets/key_master.png",
		},
		{
			Title:       "The Screen Addict",
			Description: "The player with the lowest phone battery percentage starts.",
			ImagePath:   "assets/screen_addict.png",
		},
		{
			Title:       "The Pet Parent",
			Description: "The player who last showed a photo of their pet starts.",
			ImagePath:   "assets/pet_parent.png",
		},
		{
			Title:       "The Ink Bearer",
			Description: "The player with the newest (or only) tattoo starts.",
			ImagePath:   "assets/ink_bearer.png",
		},
	}
}

func (r *Rule) GetSprite() *ebiten.Image {
	img, err := GetSprite(r.ImagePath)
	if err != nil {
		panic(err)
	}
	return img
}

// Check if rule has an associated image and file exists
func (r *Rule) HasImage() bool {
	return r.ImagePath != "" && FileExists(r.ImagePath)
}
