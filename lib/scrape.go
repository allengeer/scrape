package scrape

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"regexp"
	"bytes"
	"net/http"
)
func fetchContent(url string) string {
	response, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	} else {

		doc, err := goquery.NewDocumentFromResponse(response)

		if err != nil {
			log.Fatal(err)
		}else{
			var buffer bytes.Buffer
			tags := "h1, h2, h3, p, ul, ol, u, b, a, title, meta"

			doc.Find(tags).Each(func(i int, s *goquery.Selection) {

				tagName := s.Get(0).Data
				if tagName == "meta"{
					a, exists := s.Attr("content")
					if exists{
						buffer.WriteString(a + "\n")
					}
				}else{
					buffer.WriteString(s.Text() + "\n")
				}
			})

			return buffer.String()
		}

	}

	return ""
}

func removeTags(content string) string {
	script, _ := regexp.Compile("<script([\\s\\S]*?)</script>")
	clensed := script.ReplaceAllString(content, "")
	reg, _ := regexp.Compile("<[^>]+>")
	reclensed := reg.ReplaceAllString(clensed, "")
	rege, _ := regexp.Compile("\\s+")
	reclensede := rege.ReplaceAllString(reclensed, " ")
	return reclensede
}

func Scrape(url string) string {
	return removeTags(fetchContent(url))
}
