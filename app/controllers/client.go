package controllers

import (
	"github.com/revel/revel"
	"errors"
	models "docusys/app/models"
	// "strconv"
)

type Client struct {
	*revel.Controller
}

func (c Client) Index() revel.Result {
	clients := []models.Client{}
	res := DB.Order("id desc").Find(&clients)
	err := res.Error
	if err != nil {
		return c.RenderError(errors.New("Record not Found"))
	}
	return c.Render(clients)
}

func (c Client) Create() revel.Result {

	first_name := c.Params.Form.Get("first_name")
	last_name := c.Params.Form.Get("last_name")
	ci := c.Params.Form.Get("ci")

	c.Validation.Required(first_name).Message("Nombre es requerido")
	c.Validation.Required(last_name).Message("Apellido es requerido")
	c.Validation.Required(ci).Message("Apellido es requerido")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Client.Index)
	}

	client := models.Client{
		First_name: first_name,
		Last_name: last_name,
		Ci: ci,
	}
	res := DB.Create(&client)

	if res.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + res.Error.Error()))
	}
	return c.Redirect("/clients")
}

func (c Client) Delete() revel.Result {

	id := c.Params.Route.Get("id")
	client := []models.Client{}
	res := DB.Delete(&client, id)
	// res := DB.Unscoped().Delete(&clients, id) // delete permanently

	if res.Error != nil {
		return c.RenderError(errors.New("Record Delete failure." + res.Error.Error()))
	}

	return c.Redirect("/clients")
}

func (c Client) Edit(id int) revel.Result {

	client := []models.Client{}
	res := DB.Find(&client, id)
	err := res.Error

	if err != nil {
		return c.RenderError(errors.New("Record not Found"))
	}
	return c.Render(client)
}

func (c Client) Update(id int) revel.Result {
	
	// pk := c.Params.Form.Get("id")
	// id, _ := strconv.Atoi(pk)
	
	first_name := c.Params.Form.Get("first_name")
	last_name := c.Params.Form.Get("last_name")
	ci := c.Params.Form.Get("ci")

	c.Validation.Required(first_name).Message("Nombre es requerido")
	c.Validation.Required(last_name).Message("Apellido es requerido")
	c.Validation.Required(ci).Message("Apellido es requerido")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Client.Index)
	}

	client := models.Client{}
	// res := DB.Model(&client).Where("id = ?", id).Update("first_name", first_name)
	// res := DB.Model(&client).Where("id = ?", id).Updates(map[string]interface{}{"first_name": first_name, "last_name": last_name, "ci": ci})
	res := DB.Model(&client).Where("id = ?", id).Updates(models.Client{First_name: first_name, Last_name: last_name, Ci: ci})

	if res.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + res.Error.Error()))
	}
	return c.Redirect("/clients")
}

// func (c Client) Update() revel.Result {
	
// 	pk := c.Params.Form.Get("id")
// 	id, _ := strconv.Atoi(pk)
	
// 	first_name := c.Params.Form.Get("first_name")
// 	last_name := c.Params.Form.Get("last_name")
// 	ci := c.Params.Form.Get("ci")

// 	c.Validation.Required(first_name).Message("Nombre es requerido")
// 	c.Validation.Required(last_name).Message("Apellido es requerido")
// 	c.Validation.Required(ci).Message("Apellido es requerido")

// 	if c.Validation.HasErrors() {
// 		c.Validation.Keep()
// 		c.FlashParams()
// 		return c.Redirect(Client.Index)
// 	}

// 	client := models.Client{}
// 	// res := DB.Model(&client).Where("id = ?", id).Update("first_name", first_name)
// 	// res := DB.Model(&client).Where("id = ?", id).Updates(map[string]interface{}{"first_name": first_name, "last_name": last_name, "ci": ci})
// 	res := DB.Model(&client).Where("id = ?", id).Updates(models.Client{First_name: first_name, Last_name: last_name, Ci: ci})

// 	if res.Error != nil {
// 		return c.RenderError(errors.New("Record Create failure." + res.Error.Error()))
// 	}
// 	return c.Redirect("/clients")
// }