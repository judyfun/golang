package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()

	userParty := app.Party("/users", func(c context.Context) {
		c.Next()
	})

	prodParty := app.Party("/prod", func(c context.Context) {
		c.Next()
	})

	prodParty.Done(func(c context.Context) {
		app.Logger().Info("the request is done")
	})

	userParty.Get("/info/{id}", func(c context.Context) {
		path := c.Path()
		app.Logger().Infof("path:" + path)

		c.Writef("request path: %v", path)
	})

	prodParty.Get("/item/{id}", func(c context.Context) {
		path := c.Path()
		app.Logger().Infof("path:" + path)

		c.Writef("request path: %v", path)
	})

	prodParty.Get("/checkout/{item_id}", func(c context.Context) {
		path := c.Path()
		app.Logger().Infof("request path: %v", path)

		c.Writef("request path: %v", path)
		c.Next()
	})
	app.Run(iris.Addr(":8000"))
}
