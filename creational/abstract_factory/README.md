# Abstract factory

The Abstract Factory Design Pattern is a creational pattern that provides an interface for creating families of related or dependent objects without specifying their concrete classes.

In other words, an abstract factory is a factory of factories. With this pattern, a high level interface is provided that encapsulates a group of factories which share common themes without specifying concrete classes. These factories produce families of related products.

This pattern is typically implemented via a master factory (the abstract factory) which is capable of creating other factories, which in turn, produce products within the same family.

# Example

Imagine we have a UI library, and we want it to support multiple themes (like Light and Dark). Different themes have different styles for the same type of components (like buttons, menus, etc.).

```go
// Abstract product
type Button interface {
    Render() string
}

// Concrete product: LightButton
type LightButton struct{}

func (b LightButton) Render() string {
    return "Rendering light theme button."
}

// Concrete product: DarkButton
type DarkButton struct{}

func (b DarkButton) Render() string {
    return "Rendering dark theme button."
}

// Abstract factory
type UIFactory interface {
    CreateButton() Button
}

// Concrete factory: LightThemeFactory
type LightThemeFactory struct{}

func (f LightThemeFactory) CreateButton() Button {
    return LightButton{}
}

// Concrete factory: DarkThemeFactory
type DarkThemeFactory struct{}

func (f DarkThemeFactory) CreateButton() Button {
    return DarkButton{}
}
```

Here, `UIFactory` is the abstract factory. `LightThemeFactory` and `DarkThemeFactory` are its concrete factories. `Button` is the abstract product. `LightButton` and `DarkButton` are its concrete products.

We can add more methods to the UIFactory to create other types of UI components (like `CreateMenu`, `CreateToolbar`, etc.), and more themes by adding more concrete factories.

This pattern is especially useful when a system should be independent of how its products are created, composed, and represented, and when a family of related product objects is designed to be used together.
