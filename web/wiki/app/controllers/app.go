package controllers

import "github.com/revel/revel"
import (
	"fmt"
	"github.com/subh007/goodl/web/wiki/app/models"
)

type App struct {
	GorpController
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

// regiser the user with {username, password, nick}
func(c App) Register(user, pass string) revel.Result{
	// insert the entry into table
	userModel := models.WikiUser{UserName: user,
							Password: pass,
							Nick: "temp",
	}

	c.Txn.Insert(&userModel);

	return c.Result
}