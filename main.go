package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)
const PORT = "8000"
func main() {
	app := iris.New()

	app.Get("/", func(ctx iris.Context) {
		ctx.ServeFile("view/chat.html", false)
	})

	setupWebsocket(app)
	app.Run(iris.Addr(":"+PORT))
}

func setupWebsocket(app *iris.Application) {
	ws := websocket.New(websocket.Config{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	})
	ws.OnConnection(handleConnection)

	app.Get("/echo", ws.Handler())

	app.Any("/iris-ws.js", func(ctx iris.Context) {
		ctx.Write(websocket.ClientSource)
	})
}

func handleConnection(c websocket.Connection) {
	c.On("chat", func(msg string) {
		c.To(websocket.Broadcast).Emit("chat", msg)
	})
}