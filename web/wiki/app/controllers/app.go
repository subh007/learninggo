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

func (c App) RegisterPage() revel.Result {
	return c.Render()
}

// this function checks whether the user
// name is already taken.
func checkExistingUser(user string, c GorpController) bool {
	users, err := c.Txn.Select(models.WikiUser{}, "select * from WikiUser where UserName = ?", user)

	if err != nil {
		panic(err)
	}

	if len(users) == 0 {
		return false
	}
	fmt.Print(users)
	return true
}

// regiser the user with {username, password, nick}
func (c App) Register(user, pass string) revel.Result {
	// insert the entry into table
	userModel := models.WikiUser{UserName: user,
		Password: pass,
		Nick:     "temp",
	}
	checkExistingUser(user, c.GorpController)
	c.Txn.Insert(&userModel)

	c.Flash.Error("lgoin failed")
	return c.Result
}
