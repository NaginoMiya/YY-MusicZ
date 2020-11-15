package main

import (
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"gopkg.in/olahol/melody.v1"
)

func GetRandomMusic(genre string) []string {
	var ret []string
	const number_of_songs = 5
	const select_songs = 3
	rand.Seed(time.Now().Unix())

	var cnt []int
	for i := 0; i < number_of_songs; i++ {
		cnt = append(cnt, 0)
	}

	var randn []int

	for len(randn) < select_songs {
		x := rand.Intn(number_of_songs)
		if cnt[x] == 0 {
			cnt[x] = 1
			randn = append(randn, x)
		}
	}


	var Jpop = [...] string{"a", "b", "c", "d", "e"}
	var Jazz = [...] string{"a", "b", "c", "d", "e"}

	switch genre {
	case "Jpop":
		for i := 0; i < select_songs; i++ {
			ret = append(ret, Jpop[randn[i]])
		}
	case "Jazz":
		for i := 0; i < select_songs; i++ {
			ret = append(ret, Jazz[randn[i]])
		}
	default:
	}

	return ret
}

func main(){
	router := gin.Default()
	router.Use(static.Serve("/css", static.LocalFile("css", true)))
	router.Use(static.Serve("/js", static.LocalFile("js", true)))
	layout := "template/layout.tmpl"

	m := melody.New()

	router.GET("/", func(c *gin.Context) {

		router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/index.tmpl")))

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

		GetRandomMusic("Jpop")

		router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/music.tmpl")))

		c.HTML(http.StatusOK, "base", gin.H{
		})
	})

	router.GET("/jpop", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/jpop.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/rock", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/rock.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/edm", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/edm.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/hiphop", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/hiphop.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/classic", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/classic.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/game", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/game.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/vocaloid", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/vocaloid.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/anime", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/anime.tmpl")))

        c.HTML(http.StatusOK, "base", gin.H{
        })
	})
	
	router.GET("/all", func(c *gin.Context) {
        router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/all.tmpl")))

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