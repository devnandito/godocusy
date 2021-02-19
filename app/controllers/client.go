package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"errors"
	models "docusys/app/models"
	"strconv"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
	"time"
)

type Client struct {
	*revel.Controller
}

func (c Client) List(page int) revel.Result {
	h2 := "GO CRUD WITH POSTGRES - SHOW RECORDS"
	clients := []models.Client{}
	q := DB.Order("id desc").Find(&clients)
	err := q.Error
	p := paginator.New(adapter.NewGORMAdapter(q), 3)
	if page == 0 {
		p.SetPage(1) // set page actual
		page = 1
	}else{
		p.SetPage(page)
	}
	if err = p.Results(&clients); err != nil {
		panic(err)
	}

	for _, client := range clients {
		fmt.Println(client.First_name)
	}

	hasPage, _ := p.HasPages()
		fmt.Println("Exists pages:", hasPage)
	
	hasNext, _ := p.HasNext()
		fmt.Println("Exists next:", hasNext)
	
	hasPrev, _ := p.HasPrev()
		fmt.Println("Exists prev:", hasPrev)
	
	next, _ := p.NextPage()
		fmt.Println("Next page:", next)
	
	current, _ := p.Page()
		fmt.Println("Current page:", current)
	
	prev, _ := p.PrevPage()
	 	fmt.Println("Prev page:", prev)
	
	totalPage, _ := p.PageNums()
	 	fmt.Println("Total page:", totalPage)

	nums, _ := p.Nums()
	 	fmt.Println("Total records:", nums)
	
	var s []int

	for i:=1; i<= totalPage; i++ {
	 	s = append(s, i)
	}
	fmt.Println(s)

	// var s []string

	// for i:=1; i<= totalPage; i++ {
	//  	s = append(s, strconv.Itoa(i))
	// }
	// fmt.Println(s)
	
	return c.Render(clients, p, s, current, h2)
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
	birthday := c.Params.Form.Get("birthday")
	sex := c.Params.Form.Get("sex")

	c.Validation.Required(first_name).Message("Nombre es requerido")
	c.Validation.Required(last_name).Message("Apellido es requerido")
	c.Validation.Required(ci).Message("Apellido es requerido")
	c.Validation.Required(ci).Message("Apellido es requerido")
	c.Validation.Required(birthday).Message("Fecha de nacimiento es requerido")
	c.Validation.Required(sex).Message("Sexo es requerido")
	c.Validation.MinSize(sex, 1).Message("Debe ingresar solamente las iniciales")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Client.Index)
	}

	const SQL_DATE_FORMAT = "2006-01-02"
	t, _ := time.Parse(SQL_DATE_FORMAT, birthday)

	client := models.Client{
		First_name: first_name,
		Last_name: last_name,
		Ci: ci,
		Birthday: t,
		Sex: sex,
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
	h2 := "GO CRUD WITH POSTGRES - EDIT RECORD"
	client := []models.Client{}
	res := DB.Find(&client, id)
	err := res.Error

	if err != nil {
		return c.RenderError(errors.New("Record not Found"))
	}
	return c.Render(client, h2)
}

func (c Client) Update() revel.Result {
	
	pk := c.Params.Form.Get("id")
	id, _ := strconv.Atoi(pk)
	
	const SQL_DATE_FORMAT = "2006-01-02"

	first_name := c.Params.Form.Get("first_name")
	last_name := c.Params.Form.Get("last_name")
	ci := c.Params.Form.Get("ci")
	birthday := c.Params.Form.Get("birthday")
	sex := c.Params.Form.Get("sex")

	c.Validation.Required(first_name).Message("Nombre es requerido")
	c.Validation.Required(last_name).Message("Apellido es requerido")
	c.Validation.Required(ci).Message("Documento es requerido")
	c.Validation.Required(birthday).Message("Fecha de nacimiento es requerido")
	c.Validation.Required(sex).Message("Sexo es requerido")
	c.Validation.MinSize(sex, 1).Message("Debe ingresar solamente las iniciales")
	
	t, _ := time.Parse("2006-01-02", birthday)
	fmt.Println("Time form:", t)


	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Client.List)
	}

	client := models.Client{}
	// res := DB.Model(&client).Where("id = ?", id).Update("first_name", first_name)
	// res := DB.Model(&client).Where("id = ?", id).Updates(map[string]interface{}{"first_name": first_name, "last_name": last_name, "ci": ci})
	res := DB.Model(&client).Where("id = ?", id).Updates(models.Client{First_name: first_name, Last_name: last_name, Ci: ci, Birthday: t, Sex: sex})

	if res.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + res.Error.Error()))
	}
	return c.Redirect("/clients")
}

func (c Client) Search() revel.Result {
	h2 := "GO CRUD WITH POSTGRES - SEARCH RECORD"
	notFound := "No existe registro"
	client := []models.Client{}
	first_name := c.Params.Form.Get("first_name")
	last_name := c.Params.Form.Get("last_name")
	document := c.Params.Form.Get("document")

	if first_name == "" && last_name == "" && document == ""{
		res := DB.Find(&client)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
		
	} else if first_name == "" && last_name == "" {
		res := DB.Where("ci LIKE ?", document+"%").Find(&client)
		// res := DB.Find(&client, "ci LIKE ?", document)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
	} else if last_name == "" && document == "" {
		res := DB.Where("first_name LIKE ?", first_name+"%").Find(&client)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
	} else if first_name == "" && document == "" {
		res := DB.Where("last_name LIKE ?", last_name+"%").Find(&client)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
	} else if first_name == "" {
		res := DB.Where("last_name LIKE ? OR ci LIKE ?", last_name+"%", document+"%").Find(&client)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
	} else if last_name == "" {
		res := DB.Where("first_name LIKE ? OR ci LIKE ?", first_name, document).Find(&client)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
	} else if document == "" {
		res := DB.Where("first_name LIKE ? OR last_name LIKE ?", first_name, last_name).Find(&client)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		// return c.Render(client, h2)
	}

	return c.Render(client, h2, notFound)
	// db.Where("updated_at > ?", lastWeek).Find(&users)
}
