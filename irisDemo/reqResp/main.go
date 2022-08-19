package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()

	app.Logger().Info("---start server---")
	app.Get("/getReq", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
	})

	app.Get("/userPath", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)
		c.WriteString("The request path from client: " + path)
	})

	app.Get("/userInfo", func(c context.Context) {
		path := c.Path()
		c.WriteString("The request path from client: " + path)
		name := c.URLParam("name")
		pass := c.URLParam("pass")
		app.Logger().Infof("req path: %v, name: %v, pass: %v", path, name, pass)
	})

	app.Get("/userInfo2", func(c context.Context) {
		path := c.Path()
		name := c.URLParam("name")
		pass := c.URLParam("pass")
		c.HTML("<h1>" + path + "," + name + ":" + pass + "</h1>")
	})

	app.Post("postLogin", func(c context.Context) {
		path := c.Path()
		app.Logger().Info(path)

		name := c.PostValue("name")
		pass := c.PostValue("pass")

		c.HTML("<h1>name:" + name + ", pass:" + pass + "</h1>")
	})

	app.Post("/postParam", func(c context.Context) {
		name := c.PostValue("name")
		company := c.URLParam("company")
		c.WriteString("name:" + name + ", company:" + company)
	})

	app.Post("/postJson", func(c context.Context) {
		path := c.Path()
		var person person
		c.ReadJSON(&person)
		c.JSON(iris.Map{"message": "hello world", "code": 200})
		c.Writef("path: %v, name: %v, age: %v", path, person.Name, person.Age)
	})

	app.Post("/postXML", func(c context.Context) {
		path := c.Path()
		var student student
		c.ReadXML(&student)
		c.Writef("path: %v, name: %v, class: %v", path, student.Name, student.Class)
	})
	app.Run(iris.Addr(":8000"))
}

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type student struct {
	Name  string `xml:"stu_name"`
	Class string `xml:"stu_class"`
}
