package prompt

import "fmt"

// Prompt contains what needs to be displayed before asking for user input
type Prompt struct {
	Items []Item
}

// Print prints the prompt before waiting for user input
func (p *Prompt) Print() {
	for i := 0; i < len(p.Items); i++ {
		if i != 0 {
			p.Items[i].Prefix()
		}
		fmt.Print(p.Items[i])
	}
}

// Generate initializes all prompt items
func (p *Prompt) Generate() {
	for i := 0; i < len(p.Items); i++ {
		p.Items[i] = p.Items[i].Generate()
	}
}
