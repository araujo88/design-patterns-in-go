package main

import (
	"fmt"
	"sync"
)

// TextStyle is a flyweight that represents a specific text style.
type TextStyle struct {
	name string
}

// Format applies the text style to a text.
func (t *TextStyle) Format(text string) string {
	return fmt.Sprintf("[%s %s]", t.name, text)
}

// TextStyleFactory creates and manages TextStyle instances.
type TextStyleFactory struct {
	styles map[string]*TextStyle
	mu     sync.Mutex
}

// GetTextStyle returns a TextStyle instance. It reuses an existing instance if available.
func (f *TextStyleFactory) GetTextStyle(name string) *TextStyle {
	f.mu.Lock()
	defer f.mu.Unlock()

	if style, ok := f.styles[name]; ok {
		return style
	}

	style := &TextStyle{name: name}
	f.styles[name] = style
	return style
}

// TextRange represents a range of text within a document.
type TextRange struct {
	Start, End int
	Text       string
	Style      *TextStyle
}

// NewTextRange creates a new text range.
func NewTextRange(start, end int, text string, style *TextStyle) *TextRange {
	return &TextRange{
		Start: start,
		End:   end,
		Text:  text,
		Style: style,
	}
}

// Format formats the text range with its text style.
func (t *TextRange) Format() string {
	return t.Style.Format(t.Text)
}

func main() {
	factory := &TextStyleFactory{
		styles: make(map[string]*TextStyle),
	}

	// Both text ranges share the same TextStyle instance.
	textRange1 := NewTextRange(0, 5, "Hello", factory.GetTextStyle("Bold"))
	textRange2 := NewTextRange(6, 11, "World", factory.GetTextStyle("Bold"))

	fmt.Println(textRange1.Format())
	fmt.Println(textRange2.Format())
}
