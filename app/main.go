package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"

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

	var Jpop = [...]string{"a", "b", "c", "d", "e"}
	var Jazz = [...]string{"a", "b", "c", "d", "e"}

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

func main() {
	router := gin.Default()
	router.Use(static.Serve("/css", static.LocalFile("css", true)))
	router.Use(static.Serve("/js", static.LocalFile("js", true)))
	layout := "template/layout.tmpl"

	m := melody.New()

	router.GET("/", func(c *gin.Context) {
		router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/index.tmpl")))
		c.HTML(http.StatusOK, "base", gin.H{
		})
	})
	

	router.GET("/music/:genre", func(c *gin.Context) {

		MusicGenre := c.Param("genre")

		router.SetHTMLTemplate(template.Must(template.New("main").ParseFiles(layout, "template/music_page.tmpl")))

		var video_id string

		switch MusicGenre {
		case "jpop":
			video_id = "SX_ViT4Ra7k"
		case "rock":
			video_id = "Xnws-1Oz4kM"
		case "edm":
			video_id = "ZNT_DcTl6s0"
		case "hiphop":
			video_id = "tvTRZJ-4EyI"
		case "classic":
			video_id = "CO7xcXRkyL4"
		case "game":
			video_id = "CrkRWzsmu8E"
		case "vocaloid":
			video_id = "KsI_1XelVM8"
		case "anime":
			video_id = "3T3ofoKfEoY"
		case "all":
			video_id = "VHYdHIfLgks"
		}



		c.HTML(http.StatusOK, "base", gin.H{
			"video_id": video_id,
		})
	})


	//melodyの実装部
	router.GET("/ws/music/:genre", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})


	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.BroadcastFilter(msg, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	router.Run(":8080")
}
