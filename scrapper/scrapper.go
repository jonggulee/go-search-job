package scrapper

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	company  string
	title    string
	location string
	summary  string
}

// Scrape Indeed by a term
func Scrape(term string) {
	var baseURL string = "https://www.jobkorea.co.kr/Search/?stext=" + term + "&tabType=recruit"
	var jobs []extractedJob
	c := make(chan []extractedJob)
	totalPages := getPages()
	fmt.Println("totalPages:", totalPages)
	for i := 1; i < totalPages+1; i++ {
		go getPage(i, c)
	}

	for i := 1; i < totalPages+1; i++ {
		extractedJob := <-c
		jobs = append(jobs, extractedJob...)
	}

	// fmt.Println(jobs)
	wirteJobs(jobs)
	fmt.Println("Done, extracted", len(jobs))
}

func wirteJobs(jobs []extractedJob) {
	file, err := os.Create("jobs.csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"Link", "Company", "Title", "Location", "summary"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, job := range jobs {
		jobSlice := []string{"https://www.jobkorea.co.kr/Recruit/GI_Read/" + job.id, job.company, job.title, job.location, job.summary}
		jwErr := w.Write(jobSlice)
		checkErr(jwErr)
	}
}

func getPage(page int, mainC chan<- []extractedJob) {
	var jobs []extractedJob
	c := make(chan extractedJob)
	pageURL := baseURL + "&Page_No=" + strconv.Itoa(page)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".list-post")
	searchCards.Each(func(i int, card *goquery.Selection) {
		go extractJob(card, c)
		// jobs = append(jobs, job)
	})

	for i := 0; i < searchCards.Length(); i++ {
		job := <-c
		jobs = append(jobs, job)
	}

	mainC <- jobs
}

func extractJob(card *goquery.Selection, c chan<- extractedJob) {
	id, _ := card.Attr("data-gno")
	company := cleanString(card.Find(".post-list-corp>a").Text())
	title := cleanString(card.Find(".post-list-info>a").Text())
	location := cleanString(card.Find(".loc.short").Text())
	summary := cleanString(card.Find(".etc").Text())
	c <- extractedJob{
		id:       id,
		company:  company,
		title:    title,
		location: location,
		summary:  summary}
	// fmt.Println(id, title, location, summary)
}

func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	// fmt.Println(res)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("div.tplPagination.newVer.wide").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
}
