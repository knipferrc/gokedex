package pokemon

import (
	"fmt"
	"image"
	"strings"

	"github.com/charmbracelet/lipgloss"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
)

type Stat struct {
	Name string `json:"name"`
}

type Stats struct {
	BaseStat int  `json:"base_stat"`
	Effort   int  `json:"effort"`
	Stat     Stat `json:"stat"`
}

type Sprites struct {
	FrontDefault string `json:"front_default"`
	BackDefault  string `json:"back_default"`
}

type PokemonDetails struct {
	ID      string  `json:"id"`
	Name    string  `json:"name"`
	URL     string  `json:"url"`
	Order   int     `json:"order"`
	Sprites Sprites `json:"sprites"`
	Stats   []Stats `json:"stats"`
}

type Pokemon struct {
	Count    int              `json:"count"`
	Next     string           `json:"next"`
	Previous string           `json:"previous"`
	Results  []PokemonDetails `json:"results"`
}

type Model struct {
	Content  Pokemon
	ShowBack bool
}

// SetContent sets the content of the pokemon.
func (m *Model) SetContent(content Pokemon) {
	m.Content = content
}

// ToggleImage toggles between the front and back sprites.
func (m *Model) ToggleImage(showBack bool) {
	m.ShowBack = showBack
}

// ImageToString converts an image to a string.
func ImageToString(width, height uint, img image.Image) (string, error) {
	img = resize.Thumbnail(width, height*2-4, img, resize.Lanczos3)
	b := img.Bounds()
	w := b.Max.X
	h := b.Max.Y
	str := strings.Builder{}

	for y := 0; y < h; y += 2 {
		for x := w; x < int(width); x = x + 2 {
			str.WriteString(" ")
		}

		for x := 0; x < w; x++ {
			c1, _ := colorful.MakeColor(img.At(x, y))
			color1 := lipgloss.Color(c1.Hex())
			c2, _ := colorful.MakeColor(img.At(x, y+1))
			color2 := lipgloss.Color(c2.Hex())
			str.WriteString(lipgloss.NewStyle().Foreground(color1).Background(color2).Render("▀"))
		}

		str.WriteString("\n")
	}

	return str.String(), nil
}

func (m Model) View() string {
	pokemonList := ""

	for _, pokemon := range m.Content.Results {
		image := pokemon.Sprites.FrontDefault

		if m.ShowBack {
			image = pokemon.Sprites.BackDefault
		}

		hp := fmt.Sprintf("HP (%d)", pokemon.Stats[0].BaseStat)
		name := lipgloss.NewStyle().Width(
			lipgloss.Width(image) - lipgloss.Width(hp)).
			Render(strings.Title(pokemon.Name))
		header := lipgloss.JoinHorizontal(lipgloss.Top, name, hp)

		pokemonList += fmt.Sprintf("%s\n%s\n\n", header, image)
	}

	return pokemonList
}
