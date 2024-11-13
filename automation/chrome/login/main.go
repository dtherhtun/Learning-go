package main

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

type Product struct {
	name, price string
}

func main() {

	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	//caps.AddChrome(chrome.Capabilities{Args: []string{
	//	"--headless", // comment out this line for testing
	//}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://scrapingclub.com/exercise/basic_login/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	name, err := driver.FindElement(selenium.ByCSSSelector, "#id_name")
	if err != nil {
		log.Fatal("Error:", err)
	}
	if err := name.Clear(); err != nil {
		log.Fatal("Error:", err)
	}
	err = name.SendKeys("scrapingclub")
	if err != nil {
		log.Fatal("Error:", err)
	}

	pass, err := driver.FindElement(selenium.ByCSSSelector, "#id_password")
	if err != nil {
		log.Fatal("Error:", err)
	}

	if err := pass.Clear(); err != nil {
		log.Fatal("Error:", err)
	}

	err = pass.SendKeys("scrapingclub")
	if err != nil {
		log.Fatal("Error:", err)
	}

	btn, err := driver.FindElement(selenium.ByCSSSelector, "button[type=\"submit\"]")
	if err != nil {
		log.Fatal("Error:", err)
	}

	if err = btn.Click(); err != nil {
		log.Fatal("Error:", err)
	}

	time.Sleep(5 * time.Minute)
}
