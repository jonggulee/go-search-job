package main

import (
	"jonggu/search-job/scrapper"
	"os"
	"strings"

	"github.com/labstack/echo"
)

const fileName string = "jobs.csv"

func main() {
	scrapper.Scrape("term")
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	term := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	userFileName := "jobs-" + term + ".csv"
	defer os.Remove(fileName)
	return c.Attachment(fileName, userFileName)
}
