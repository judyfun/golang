package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()

	// Handler
	app.Handle("GET", "/user", func(c context.Context) {
		path := c.Path()
		app.Logger().Infof("request coming in..." + path)
		name := c.URLParam("name")
		//name := c.Params().Get("name")
		age := c.Params().Get("age")
		c.Writef("Request path: %v, name: %v, age:%v", path, name, age)
	})

	app.Get("/weather/{date}/{city}", func(c context.Context) {
		path := c.Path()
		date := c.Params().GetString("date")
		city := c.Params().GetString("city")

		c.Writef("path: %v, date: %v, city: %v", path, date, city)
	})

	app.Get("/user/{name:string}/{age:int}", func(c context.Context) {
		path := c.Path()
		name := c.Params().GetString("name")
		age := c.Params().GetString("age")

		c.Writef("path: %v, name: %v, age: %v", path, name, age)
	})

	app.Get("user/{isLogin:bool}", func(c context.Context) {
		isLogin, _ := c.Params().GetBool("isLogin")

		if isLogin {
			c.Writef("You are already login.")
		} else {
			c.Writef("You are not login.")
		}
	})

	app.Run(iris.Addr(":8000"))
}
