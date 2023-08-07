# Flyweight

The Flyweight Design Pattern is a structural design pattern that provides a way to use objects in large numbers when a simple repeated representation would use an unacceptable amount of memory.

Flyweight pattern introduces the concept of sharing common parts of the object state among multiple objects instead of keeping all of the data in each object. The object containing the shared state is usually called a "flyweight". The pattern allows programs to support vast quantities of objects by keeping their memory consumption low.

A typical use case for the Flyweight pattern is a game, for example, where you have numerous objects (trees, buildings, NPCs, etc.) and you can't afford to expend your memory resources to store each object separately with all their properties.

## Example

Let's consider a simple text formatting tool. This tool would need to format different ranges of text within a document with specific styles (bold, italic, underline, etc.). These styles are shared among multiple text ranges, and each text range could be considered a unique object. The Flyweight pattern would allow us to reuse style instances instead of duplicating them for each text range, therefore saving memory.

```go
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
```

In this example, `TextStyle` is a flyweight that represents a text style. `TextStyleFactory` manages the creation and reuse of `TextStyle` instances. `TextRange` represents a range of text within a document and uses a `TextStyle` to format its text. By using the Flyweight pattern, the "Bold" `TextStyle` instance is only created once and shared among multiple `TextRange` instances, thereby saving memory.
