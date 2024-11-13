package main

import (
	"encoding/csv"
	"log"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Product struct {
	name, price string
}

func main() {
	var products []Product

	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
	}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://scrapingclub.com/exercise/list_infinite_scroll/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	productElements, err := driver.FindElements(selenium.ByCSSSelector, ".post")
	if err != nil {
		log.Fatal("Error:", err)
	}

	for _, productElement := range productElements {
		// select the name and price nodes
		nameElement, err := productElement.FindElement(selenium.ByCSSSelector, "h4")
		priceElement, err := productElement.FindElement(selenium.ByCSSSelector, "h5")
		// extract the data of interest
		name, err := nameElement.Text()
		price, err := priceElement.Text()
		if err != nil {
			log.Fatal("Error:", err)
		}

		// add the scraped data to the list
		product := Product{}
		product.name = name
		product.price = price
		products = append(products, product)
	}

	file, err := os.Create("products.csv")
	if err != nil {
		log.Fatal("Error:", err)
	}

	defer file.Close()

	writer := csv.NewWriter(file)

	headers := []string{
		"name",
		"price",
	}

	writer.Write(headers)

	for _, product := range products {

		// converting a Product to an array of strings
		record := []string{
			product.name,
			product.price,
		}

		// writing a new CSV record
		writer.Write(record)
	}

	defer writer.Flush()
}
