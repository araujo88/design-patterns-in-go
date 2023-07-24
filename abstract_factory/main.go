package main

import (
	"fmt"
)

type Button interface {
	Render() string
}

type LightButton struct{}

func (b LightButton) Render() string {
	return "Rendering light theme button."
}

type DarkButton struct{}

func (b DarkButton) Render() string {
	return "Rendering dark theme button."
}

type UIFactory interface {
	CreateButton() Button
}

type LightThemeFactory struct{}

func (f LightThemeFactory) CreateButton() Button {
	return LightButton{}
}

type DarkThemeFactory struct{}

func (f DarkThemeFactory) CreateButton() Button {
	return DarkButton{}
}

func main() {
	var factory UIFactory
	var button Button

	// Select the factory based on user preference or config file, for example:
	userPrefersDarkMode := true // Imagine this value is from user settings

	if userPrefersDarkMode {
		factory = DarkThemeFactory{}
	} else {
		factory = LightThemeFactory{}
	}

	// Create a button using the factory
	button = factory.CreateButton()

	// Use the button
	fmt.Println(button.Render())
}
