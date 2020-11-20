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
		ret = append(ret, music_ids[randn[i]])
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
			music_ids := []string{"SX_ViT4Ra7k", "1krJijKL384", "cMRYfNTlpqo", "9qRCARM_LfE", "YapsFDcGe_s"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.
		case "rock":
			music_ids := []string{"O_DLtVuiqhI", "QzmsVn2cHaA", "PCp2iXA1uLE", "PLgYflfgq0M", "5pkBqmX2ymc"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.
		case "edm":
			music_ids := []string{"Ni2PSh0N_58", "YJVmu6yttiw", "3nad7SQhtno", "ALZHF5UqnU4", "r2LpOUwca94"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]
		case "hiphop":
			music_ids := []string{"c0dMiqqve5E", "rmeI_Qk1rrk", "r_0JjYUe5jo", "_dAzUOzWvrk", "tvTRZJ-4EyI"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]
		case "classic":
			music_ids := []string{"9n8R3x68yuI", "CJV4l0cnNO4", "irC7b-SwA8g", "3JZiZcXf12o", "3439BgooWmQ"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]
		case "game":
			music_ids := []string{"7DuUT15c8SE", "KYVNZj9-wZI", "vekg2OXHniU", "eto7Wsv9eqg", "7knlsjItLX8"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.
		case "vocaloid":
			music_ids := []string{"MUahuOoNZNY", "romqp_SB4tU", "KsI_1XelVM8", "UnIhRpIT7nc", "TdegG12IiFo"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.
		case "anime":
			music_ids := []string{"n7VZxg9pxkg", "9liVljr-1cs", "CocEAA4idEU", "3T3ofoKfEoY", "6Sh_ZMXBYG0"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.

		case "all":
			music_ids := []string{"VHYdHIfLgks", "BKl4gZDWP34", "oJlmclcLD74", "F6KgJox-NmM", "Tq49NR_HzfY"}
			selected := GetRandomMusic(music_ids)
			video_id = selected[0]//後に配列になります.
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

		// send_data.Titleの長さが0の時は無効なURLとする
		if(len(send_data.Title) != 0){
			m.BroadcastFilter(send_data_json, func(q *melody.Session) bool {
				return q.Request.URL.Path == s.Request.URL.Path
			})
		}else{
			s.Write(send_data_json)
		}
	})

	router.Run(":8080")
}
