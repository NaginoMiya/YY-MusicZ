package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"

	"gopkg.in/olahol/melody.v1"
)

func main(){
	router := gin.Default()
	router.Use(static.Serve("/css", static.LocalFile("app/css", true)))
	router.Use(static.Serve("/js", static.LocalFile("app/js", true)))
	layout := "app/template/layout.tmpl"

	m := melody.New()

	router.GET("/", func(c *gin.Context) {

		router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "app/template/index.tmpl")))

		//こいつは消しても大丈夫です
		b_messege := "stringだよ"

		c.HTML(http.StatusOK, "base", gin.H{
			//こんな感じで自由に変えてね
			"a": "生の埋め込みだよ!",
			"b": b_messege,
		})
	})

	//チャット部分
	router.GET("/music", func(c *gin.Context) {

		router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "app/template/music.tmpl")))

		c.HTML(http.StatusOK, "base", gin.H{
		})
	})

	//melodyの実装部
	router.GET("/ws", func(c *gin.Context){
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte){
		m.Broadcast(msg)
	})

	router.Run(":8080")
}