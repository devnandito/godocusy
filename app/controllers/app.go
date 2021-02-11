package controllers

import (
	"github.com/revel/revel"
)

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	greeting := "Aloha World"
	return c.Render(greeting)
}

func (c App) Hello(name string) revel.Result {
	c.Validation.Required(name).Message("Your name is required!")
	c.Validation.MinSize(name, 3).Message("Your name is not long enough!")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	return c.Render(name)
}

func (c App) Person() revel.Result {
	people, err := allPerson()
	if err != nil {
		panic(err)
	}
	return c.Render(people)
}

func (c App) savePerson() revel.Result {
	var err error
	first_name := c.Params.Get("first_name")
	_, err = insertPerson(first_name)
	if err != nil {
		panic(err)
	}

	return c.Redirect(App.Index)

}