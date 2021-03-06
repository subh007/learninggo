package controllers

import "github.com/revel/revel"
import (
	"fmt"
	"github.com/subh007/goodl/web/wiki/app/models"
	"github.com/subh007/goodl/web/wiki/app/routes"
)

type App struct {
	GorpController
}

func (c App) Index() revel.Result {
	return c.Render()
}

func (c App) Login(user, pass string) revel.Result {
	fmt.Println("username :" + user)

	// read the user info from db to authenticate
	users, err := c.Txn.Select(models.WikiUser{},
		"select * from WikiUser where UserName = ? and Password = ?",
		user,
		pass)

	fmt.Println("query is done")
	if err != nil {
		panic(err)
	}
	if len(users) == 0 {
		// user is not authenticated
		fmt.Println("redirected !!")
		return c.Redirect(routes.App.Index())
	} else {
		// Authentication successfull.
		fmt.Println("user Authenticated")
		// Redirec to wiki User home
		return c.Result
	}
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
	userModel := models.WikiUser{
		UserName: user,
		Password: pass,
		Nick:     "temp",
	}
	if !checkExistingUser(user, c.GorpController) {
		// user does not exist insert the entry
		c.Txn.Insert(&userModel)
		fmt.Println("added the new user")
	} else {
		fmt.Println("existing user.")
	}

	// after user auth redirect to index.
	return c.Redirect(routes.App.Index())
}
