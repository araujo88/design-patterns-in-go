# Memento

The Memento design pattern is a behavioral design pattern that provides the capability to restore an object to its previous state. This pattern is particularly useful when you need to implement features like undo or save functionality. The pattern involves three types of participant classes:

- **Originator**: The class that has an internal state we wish to save and restore later.
- **Memento**: The class that stores the internal state of the Originator. The state could be saved outside the Originator but without exposing the details.
- **Caretaker**: The class that keeps track of various states of the Originator. It keeps a list of Memento objects to achieve this.

## Example

Let's implement the Memento pattern to maintain the history of a simple text editor, allowing us to undo changes.

Here is the code in Go:

```go
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
```

In this example:

- `Editor` represents the `Originator`. It maintains some text and can produce snapshots (`Memento`) of its current state.
- `Snapshot` represents the `Memento`. It stores a snapshot of the `Editor`'s state.
- `History` is the `Caretaker`. It maintains a history of states that can be used to restore the `Editor` to previous states.
