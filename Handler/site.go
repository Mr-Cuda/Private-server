package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/Mr-Cuda/Private-server"
)

type Link struct {
	Title       string
	Url         string
	Description string
}

type Object struct {
	Title       string
	Slug        string
	Description string
	Items       []Link
}

func main() {
	GenerateHTML()
	input, err := ioutil.ReadFile("./site/index.html")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(input)
	query, err := goquery.NewDocumentFromReader(buf)
	if err != nil {
		panic(err)
	}

	objs := []Object{}
	query.Find("body #content ul ul").First().Each(func(_ int, s *goquery.Selection) {

		s.Find("li a").Each(func(_ int, s *goquery.Selection) {
			selector, _ := s.Attr("href")
			obj := makeObjById(selector, query.Find("body"))
			objs = append(objs, obj)
		})
	})

	makeSiteStruct(objs)
	makeSitemap(objs)
	changeLinksInIndex(string(input), query)
}

func makeSiteStruct(objs []Object) {
	for _, obj := range objs {
		folder := fmt.Sprintf("site/%s", obj.Slug)
		err := os.Mkdir(folder, 0755)
		if err != nil {
			log.Println(err)
		}

		t := template.Must(template.ParseFiles("site/cat-tmpl.html"))
		f, _ := os.Create(fmt.Sprintf("%s/index.html", folder))
		t.Execute(f, obj)
	}
}

func makeSitemap(objs []Object) {
	t := template.Must(template.ParseFiles("tmpl/sitemap-tmpl.xml"))
	f, _ := os.Create("site/sitemap.xml")
	t.Execute(f, objs)
}

func makeObjById(selector string, s *goquery.Selection) (obj Object) {
	s.Find(selector).Each(func(_ int, s *goquery.Selection) {
		desc := s.NextFiltered("p")
		ul := desc.NextFiltered("ul")

		links := []Link{}
		ul.Find("li").Each(func(_ int, s *goquery.Selection) {
			url, _ := s.Find("a").Attr("href")
			link := Link{
				Title:       s.Find("a").Text(),
				Description: s.Text(),
				Url:         url,
			}
			links = append(links, link)
		})
		obj = Object{
			Slug:        slugify.Slugify(s.Text()),
			Title:       s.Text(),
			Description: desc.Text(),
			Items:       links,
		}
	})
	return
}

func changeLinksInIndex(html string, query *goquery.Document) {
	query.Find("body #content ul li ul li a").Each(func(_ int, s *goquery.Selection) {

		href, exists := s.Attr("href")
		if exists {
			uri := strings.SplitAfter(href, "#")
			if len(uri) >= 2 && uri[1] != "contents" {
				html = strings.ReplaceAll(
					html, fmt.Sprintf(`href="%s"`, href), fmt.Sprintf(`href="%s"`, uri[1]))
			}
		}
	})

	os.WriteFile("./site/index.html", []byte(html), 0644)
}
