package service

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/gocolly/colly"

	"github.com/chijiokejoseph/gowayback-beats/consts"
)

type Song struct {
	Title  string
	Artist string
}

func NewSong(title, artist string) *Song {
	return &(Song{title, artist})
}

func Scrape(pastDate time.Time) ([]*Song, error) {
	dateStr := pastDate.Format(consts.DATE_FORMAT)
	url := strings.Join([]string{consts.URL, dateStr, "/"}, "")
	fmt.Printf("url = %s\n", url)
	coll := colly.NewCollector()
	songs := make([]*Song, 0, 100)
	selector := ".o-chart-results-list__item"

	coll.OnResponse(func(r *colly.Response) {
		if r.StatusCode != 200 {
			fmt.Println("Request unsuccessful")
			return
		}
		fmt.Printf("Success with the status code: %d\n", r.StatusCode)
	})

	count := 0

	coll.OnHTML(selector, func(h *colly.HTMLElement) {
		titleQuery := "#title-of-a-story"
		artistQuery := ".c-label"
		var (
			titleList  []string
			artistList []string
		)
		count++
		appendText := func(store *[]string, elem *colly.HTMLElement) {
			text := elem.Text
			text = strings.ReplaceAll(text, "\t", "")
			text = strings.ReplaceAll(text, "\n", "")
			text = strings.Trim(text, " ")
			*store = append(*store, text)
		}

		h.ForEach(titleQuery, func(i int, elem *colly.HTMLElement) {
			appendText(&titleList, elem)
		})

		h.ForEach(artistQuery, func(i int, elem *colly.HTMLElement) {
			appendText(&artistList, elem)
		})

		if len(titleList) == 0 || len(artistList) == 0 {
			return
		}

		title := strings.Join(titleList, "")
		artist := strings.Join(artistList, "")

		songs = append(songs, NewSong(title, artist))
	})

	coll.Visit(url)

	if songs[0] == nil {
		return songs, errors.New("Html Parsing unsuccessful")
	}
	return songs, nil
}
