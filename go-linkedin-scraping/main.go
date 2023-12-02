package main

import (
	"fmt"
	"log"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./ChromeDriver/chromedriver.exe", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless-new", // comment out this line for testing
	}})

	// create a new remote client with the specified options
	driver, err := selenium.NewRemote(caps, "")

	if err != nil {
		log.Fatal("Error:", err)
	}

	// maximize the current window to avoid responsive rendering
	err = driver.MaximizeWindow("")
	if err != nil {
		log.Fatal("Error:", err)
	}

	err = driver.Get("https://www.jobstreet.co.id/id/software-engineer-jobs/")
	if err != nil {
		log.Fatal("Error:", err)
	}

	html, err := driver.PageSource()
	if err != nil {
		log.Fatal("Error:", err)
	}
	fmt.Println(html)

}
