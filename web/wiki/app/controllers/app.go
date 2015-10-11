package controllers

import "github.com/revel/revel"
import "fmt"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Login(user, pass string) revel.Result {
	fmt.Println("username :" + user)
	return c.Result
}

func (c App) RegisterPage() revel.Result{
	return c.Render()
}

func(c App) Register(user, pass string) revel.Result{
	return c.Result
}