package main

import "fmt"

// Originator
type Editor struct {
	text string
}

func (e *Editor) SetText(text string) {
	e.text = text
}

func (e *Editor) CreateSnapshot() *Snapshot {
	return &Snapshot{e.text}
}

func (e *Editor) RestoreSnapshot(s *Snapshot) {
	e.text = s.GetText()
}

// Memento
type Snapshot struct {
	text string
}

func (s *Snapshot) GetText() string {
	return s.text
}

// Caretaker
type History struct {
	snapshots []*Snapshot
}

func (h *History) AddSnapshot(s *Snapshot) {
	h.snapshots = append(h.snapshots, s)
}

func (h *History) Undo() *Snapshot {
	length := len(h.snapshots)

	if length < 2 {
		return nil // Cannot undo; not enough snapshots
	}

	lastSnapshot := h.snapshots[length-2] // Get the second last snapshot
	h.snapshots = h.snapshots[:length-1]  // Remove the last snapshot

	return lastSnapshot // Return the second last snapshot
}

func main() {
	editor := &Editor{}
	history := &History{}

	editor.SetText("First line.")
	history.AddSnapshot(editor.CreateSnapshot())

	editor.SetText("Second line.")
	history.AddSnapshot(editor.CreateSnapshot())

	editor.SetText("Third line.")
	history.AddSnapshot(editor.CreateSnapshot())

	fmt.Println("Current Text:", editor.text)

	// Undo to previous state
	editor.RestoreSnapshot(history.Undo())
	fmt.Println("Undo 1: ", editor.text)

	// Undo to previous state
	editor.RestoreSnapshot(history.Undo())
	fmt.Println("Undo 2: ", editor.text)

	// Undo to previous state; As this is the last state, it should remain the same.
	snapshot := history.Undo()
	if snapshot != nil {
		editor.RestoreSnapshot(snapshot)
	}
	fmt.Println("Undo 3: ", editor.text)
}
