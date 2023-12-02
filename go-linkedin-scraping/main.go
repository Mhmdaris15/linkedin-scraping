package main

import (
	"go-linkedin-scraping/utils"
	"log"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {
	// initialize a Chrome browser instance on port 4444
	service, err := selenium.NewChromeDriverService("./ChromeDriver/chromedriver", 4444)
	if err != nil {
		log.Fatal("Error:", err)
	}
	defer service.Stop()

	// proxyServerURL := "36.37.86.60"
	customUserAgent := "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36"

	// configure the browser options
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--user-agent=" + customUserAgent,
		// "--proxy-server=" + proxyServerURL,
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

	// navigate to the login page
	err = driver.Get("https://www.linkedin.com/home")
	if err != nil {
		log.Fatal("Error:", err)
	}

	driver, err = utils.LoginLinkedIn(driver)
	if err != nil {
		log.Fatal("Error:", err)
	}

	// navigate to the search page
	driver, err = utils.SearchJob(driver, "business development")
	if err != nil {
		log.Fatal("Error:", err)
	}

	// Wait untill the last job list is loaded
	err = driver.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		lastJobList, _ := wd.FindElements(selenium.ByCSSSelector, ".jobs-search-results__list-item:nth-child(25)")

		if lastJobList != nil {
			// Scroll container with class "jobs-search-results-list" to the bottom to load all the jobs

			return true, nil
		}

		log.Print("Waiting for the last job list to load...")

		return false, nil
	}, 10*time.Second)
	if err != nil {
		log.Fatal("Error:", err)
	}

	time.Sleep(3 * time.Second)
	// Scroll container with class "jobs-search-results-list" to the bottom to load all the jobs
	err = utils.ScrollToBottom(driver, "jobs-search-results-list")
	if err != nil {
		log.Fatal("Error:", err)
	}

	time.Sleep(15 * time.Second)
}
