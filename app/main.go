package main

import (
	"encoding/json"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type SendData struct {
	Title string `json:"title"`
	Url string   `json:"url"`
}


func GetRandomMusic(music_ids []string) []string {
	var ret []string
	const number_of_songs = 5
	const select_songs = 1
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

	for i := 0; i < select_songs; i++ {
		ret = append(ret, music_ids[i])
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
			music_ids := []string{"SX_ViT4Ra7k", "SX_ViT4Ra7k", "SX_ViT4Ra7k", "SX_ViT4Ra7k", "SX_ViT4Ra7k"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.
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
		url := string(msg)
		res, err := http.Get(url)
		if err != nil {
			log.Println(err)
		}
		defer res.Body.Close()

		send_data := new(SendData)
		send_data.Url = url

		doc, _ := goquery.NewDocumentFromReader(res.Body)
		doc.Find("title").Each(func(i int, s *goquery.Selection) {
			fmt.Println(s.Text())
			send_data.Title = s.Text()
		})

		send_data_json, _ := json.Marshal(send_data)

		m.BroadcastFilter(send_data_json, func(q *melody.Session) bool {
			return q.Request.URL.Path == s.Request.URL.Path
		})
	})

	router.Run(":8080")
}
