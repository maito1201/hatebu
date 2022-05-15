package hatebu

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	Href  string
	Title string
}

func ScrapeHotEntry() ([]ScrapeResult, error) {
	var ret []ScrapeResult
	res, err := http.Get("https://b.hatena.ne.jp/hotentry/it")
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	document, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}

	document.Find("div.entrylist-contents-main").Each(func(i int, s *goquery.Selection) {
		result := s.Find("h3.entrylist-contents-title").Find("a")
		href, _ := result.Attr("href")
		title, _ := result.Attr("title")
		ret = append(ret, ScrapeResult{Href: href, Title: title})
	})
	return ret, nil
}
