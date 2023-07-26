package main

import "fmt"

// Form - our prototype interface
type Form interface {
	Clone() Form
}

// RegistrationForm is a concrete type implementing Form
type RegistrationForm struct {
	Name  string
	Email string
}

// Clone method for RegistrationForm
func (r *RegistrationForm) Clone() Form {
	return &RegistrationForm{
		Name:  r.Name,
		Email: r.Email,
	}
}

func main() {
	// create a prototype of the RegistrationForm
	prototype := &RegistrationForm{
		Name:  "Default Name",
		Email: "default@example.com",
	}

	// use the prototype to clone a new form
	clonedForm := prototype.Clone()

	// type assertion to RegistrationForm
	regForm, ok := clonedForm.(*RegistrationForm)
	if !ok {
		fmt.Println("Type assertion to RegistrationForm failed")
		return
	}

	fmt.Printf("Name: %s, Email: %s\n", regForm.Name, regForm.Email)
}
