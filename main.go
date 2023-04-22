package main

import (
	"fmt"
	"github.com/gocolly/colly"
)

const DOMAIN string = "time.is"

type Data struct {
	Scope   string
	Time    string
	DayDate string
}

func main() {
	data, err := getScrapeData()

	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	printData(data)
}

func getScrapeData() (Data, error) {
	data := Data{}

	c := colly.NewCollector(
		colly.AllowedDomains(DOMAIN),
	)

	c.OnHTML("[itemscope]", func(e *colly.HTMLElement) {
		data.Scope = e.Text
	})

	c.OnHTML("#clock0_bg", func(e *colly.HTMLElement) {
		data.Time = e.Text
	})

	c.OnHTML("#dd", func(e *colly.HTMLElement) {
		data.DayDate = e.Text
	})

	err := c.Visit("https://" + DOMAIN)

	return data, err
}

func printData(data Data) {
	fmt.Println("Time in:", data.Scope)
	fmt.Println(data.Time)
	fmt.Println(data.DayDate)
}
