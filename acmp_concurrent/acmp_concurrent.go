package acmp_concurrent

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

func Difficulties(urls []string) map[string]float64 {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var cache = make(map[string]float64)
	for _, url := range urls {
		u := url
		wg.Add(1)
		go func() {
			defer wg.Done()
			diff := getRequest(u)
			mu.Lock()
			defer mu.Unlock()
			cache[u] = diff
		}()
	}
	wg.Wait()

	return cache
}

func getRequest(url string) float64 {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return -1
	}

	req.AddCookie(&http.Cookie{Name: "English", Value: "1"})

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
		return -1
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalln(err)
		return -1
	}

	r, _ := regexp.Compile("(\\d+)%\\)")
	res := -1.0
	doc.Find(".nomargin table tbody tr td table tbody tr td center i").
		Each(func(i int, selection *goquery.Selection) {
			str := r.FindString(selection.Text())
			str = strings.TrimRight(str, ")")
			str = strings.TrimRight(str, "%")
			res, err = strconv.ParseFloat(str, 64)
			if err != nil {
				log.Fatalln(err)
			}
		})

	return res
}
