package main

import (
	"cmp"
	"encoding/base64"
	"log"
	"os"
	"time"

	capsolver_go "github.com/capsolver/capsolver-go"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type Product struct {
	name, price string
}

func toBase64(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func main() {

	service, err := selenium.NewChromeDriverService("/usr/local/bin/chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	curDir, _ := os.Getwd()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
			//"--headless", // comment out this line for testing
	}, Prefs: map[string]interface{}{"download.default_directory": curDir}})

	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		log.Fatal("Error:", err)
	}

	apiKey := cmp.Or(os.Getenv("API_KEY_RECAP"), "verySecretKey")
	capSolver := capsolver_go.CapSolver{ApiKey: apiKey}

	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://scrapingclub.com/exercise/basic_captcha/")
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

	img, err := driver.FindElement(selenium.ByCSSSelector, ".captcha")
	if err != nil {
		log.Fatal("Error:", err)
	}

	screenshot, err := img.Screenshot(false)
	if err != nil {
		log.Fatal("Screen shot error:", err)
	}

	solution, err := capSolver.Solve(map[string]any{
		"type":   "ImageToTextTask",
		"module": "common",
		"body":   base64.StdEncoding.EncodeToString(screenshot),
	})
	if err != nil {
		log.Fatal(err)
		return
	}

	capInbox, err := driver.FindElement(selenium.ByCSSSelector, "#id_captcha_1")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = capInbox.SendKeys(solution.Solution.Text)
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
