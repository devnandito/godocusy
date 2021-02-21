package controllers

import (
	"fmt"
	"github.com/revel/revel"
	"errors"
	models "docusys/app/models"
	"strconv"
	"strings"
	"github.com/vcraescu/go-paginator/v2"
	"github.com/vcraescu/go-paginator/v2/adapter"
	"time"
)

type Client struct {
	*revel.Controller
}

func (c Client) Index() revel.Result {
	h2 := "GO CRUD WITH POSTGRES - SEARCH"
	return c.Render(h2)
}

func (c Client) Create(client *models.Client) revel.Result {
	client.Validate(c.Validation)

	first_name := c.Params.Form.Get("client.First_name")
	last_name := c.Params.Form.Get("client.Last_name")
	ci := c.Params.Form.Get("client.Ci")
	birthday := c.Params.Form.Get("client.Birthday")
	sex := c.Params.Form.Get("client.Sex")

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Client.ListPage)
	}

	const SQL_DATE_FORMAT = "2006-01-02"
	t, _ := time.Parse(SQL_DATE_FORMAT, birthday)

	data := models.Client{
		First_name: first_name,
		Last_name: last_name,
		Ci: ci,
		Birthday: t,
		Sex: sex,
	}
	res := DB.Create(&data)

	if res.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + res.Error.Error()))
	}
	return c.Redirect(Client.ListPage)
}

func (c Client) Edit(id int, first_name, last_name, ci, sex string, birthday time.Time) revel.Result {

	client := []models.Client{}
	res := DB.Find(&client, id)
	err := res.Error

	if err != nil {
		return c.RenderError(errors.New("Record not Found"))
	}

	for _, i := range(client){
		first_name = i.First_name
		last_name = i.Last_name
		ci = i.Ci
		sex = i.Sex
		birthday = i.Birthday
	}

	return c.Render(id, first_name, last_name, sex, ci, birthday)
}

func (c Client) Update(client *models.Client) revel.Result {
	client.Validate(c.Validation)
	
	pk := c.Params.Form.Get("client.ID")
	id, _ := strconv.Atoi(pk)
	
	first_name := c.Params.Form.Get("client.First_name")
	last_name := c.Params.Form.Get("client.Last_name")
	ci := c.Params.Form.Get("client.Ci")
	birthday := c.Params.Form.Get("client.Birthday")
	sex := c.Params.Form.Get("client.Sex")
	
	t, _ := time.Parse("2006-01-02", birthday)
	fmt.Println("Time form:", t)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Client.Edit, id, first_name, last_name, sex, ci, birthday)
	}

	res := DB.Model(&client).Where("id = ?", id).Updates(models.Client{First_name: first_name, Last_name: last_name, Ci: ci, Birthday: t, Sex: sex})

	if res.Error != nil {
		return c.RenderError(errors.New("Record Create failure." + res.Error.Error()))
	}
	return c.Redirect(Client.ListPage)

}

func (c Client) Delete(id int) revel.Result {

	client := []models.Client{}
	res := DB.Delete(&client, id)

	if res.Error != nil {
		return c.RenderError(errors.New("Record Delete failure." + res.Error.Error()))
	}

	return c.Redirect(Client.ListPage)
}

func (c Client) Search() revel.Result {
	h2 := "GO CRUD WITH POSTGRES - SEARCH RECORD"
	notFound := "No existe registro"
	clients := []models.Client{}
	first_name := strings.ToUpper(c.Params.Form.Get("first_name"))
	last_name := strings.ToUpper(c.Params.Form.Get("last_name"))
	document := c.Params.Form.Get("document")
	var count int64

	if first_name == "" && last_name == "" && document == ""{
		res := DB.Order("last_name asc").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
		
	} else if first_name == "" && last_name == "" {
		res := DB.Order("ci asc").Where("ci LIKE ?", document+"%").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
	} else if last_name == "" && document == "" {
		res := DB.Where("first_name LIKE ?", first_name+"%").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
	} else if first_name == "" && document == "" {
		res := DB.Order("last_name asc").Where("last_name LIKE ?", last_name+"%").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
	} else if first_name == "" {
		res := DB.Order("last_name asc").Where("last_name LIKE ? OR ci LIKE ?", last_name+"%", document+"%").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
	} else if last_name == "" {
		res := DB.Order("last_name asc").Where("first_name LIKE ? OR ci LIKE ?", first_name+"%", document+"%").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
	} else if document == "" {
		res := DB.Order("last_name asc, first_name asc").Where("first_name LIKE ? OR last_name LIKE ?", first_name+"%", last_name+"%").Find(&clients).Count(&count)
		err := res.Error

		if err != nil {
			return c.RenderError(errors.New("Record not Found"))
		}
	}

	return c.Render(clients, h2, notFound, count)
}

func (c Client) ListPage(page int) revel.Result {
	h2 := "GO CRUD WITH POSTGRES - SHOW RECORDS"
	clients := []models.Client{}
	q := DB.Order("id desc").Find(&clients)
	err := q.Error
	p := paginator.New(adapter.NewGORMAdapter(q), 10)
	if page == 0 {
		p.SetPage(1) // set page actual
		page = 1
	}else{
		p.SetPage(page)
	}
	if err = p.Results(&clients); err != nil {
		panic(err)
	}

	totalPage, _ := p.PageNums()
	
	current, _ := p.Page()
	
	var s []int

	for i:=1; i<= totalPage; i++ {
	 	s = append(s, i)
	}
	
	return c.Render(clients, p, s, current, h2)
}