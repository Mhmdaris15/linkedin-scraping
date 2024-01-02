package main

import (
	"fmt"
	"go-linkedin-scraping/utils"
	"log"
	"sync"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

var isLogin bool

func main() {
	// Number of instances/drivers you want to run concurrently
	numInstances := 1

	// Create 3 array of job name to search
	jobNames := [1]string{"Backend Engineer"}

	// Ask driver to login or use existing cookies
	var loginOrUseExistingCookies string
	log.Print("Login or use existing cookies? (y for login, n for use existing cookies): ")
	_, err := fmt.Scanln(&loginOrUseExistingCookies)
	if err != nil {
		log.Fatal("Error:", err)
	}

	if loginOrUseExistingCookies == "y" {
		isLogin = true
	} else {
		// Load cookies from JSON file
		isLogin = false
	}

	// Create a WaitGroup to wait for all Goroutines to finish
	var wg sync.WaitGroup

	// Connect to MongoDB with goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		utils.ConnectDB()
	}()

	// Loop to create and run the specified number of instances
	for i := 0; i < numInstances; i++ {
		wg.Add(1) // Increment the WaitGroup counter for each Goroutine

		// Goroutine to run each instance
		go func(instanceID int) {
			// Defer the WaitGroup Done method to decrement the counter when the Goroutine completes
			defer wg.Done()

			// Your scraping code for each instance
			Scrape(instanceID, jobNames[instanceID])
		}(i)
	}

	// Wait for all Goroutines to finish
	wg.Wait()
}

func Scrape(instanceID int, jobName string) {
	// initialize a Chrome browser instance on port 4444
	port := 4444 + instanceID
	service, err := selenium.NewChromeDriverService("./chromedriver-win64/chromedriver.exe", port)
	if err != nil {
		log.Fatal("Error:", err)
	}
	// defer service.Stop()
	defer EndTheProgram(service)

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
	if err := driver.MaximizeWindow(""); err != nil {
		log.Fatal("Error:", err)
	}

	// navigate to the login page
	if err := driver.Get("https://www.linkedin.com/home"); err != nil {
		log.Fatal("Error:", err)
	}

	if isLogin {
		if err := utils.LoginLinkedIn(&driver); err != nil {
			log.Fatal("Error:", err)
		}
	} else {
		if err := utils.LoadCookiesFromJSON(&driver, "cookies", "./data"); err != nil {
			log.Fatal("Error:", err)
		}
	}

	time.Sleep(3 * time.Second)

	// Save cookies to JSON file
	// if err := utils.SaveCookiesToJSON(&driver, "cookies", "./data"); err != nil {
	// 	log.Fatal("Error:", err)
	// }

	// jobName := "full stack engineer"

	// navigate to the search page
	if err := utils.SearchJob(&driver, jobName); err != nil {
		log.Fatal("Error:", err)
	}

	// Wait untill the last job list is loaded
	err = utils.WaitForJobsAndClickFirst(&driver, jobName)
	if err != nil {
		log.Fatal("Error:", err)
	}

	time.Sleep(3 * time.Second)
	// Scroll container with class "jobs-search-results-list" to the bottom to load all the jobs
	err = utils.ScrollToBottom(&driver, ".jobs-search-results-list")
	if err != nil {
		log.Fatal("Error:", err)
	}

	log.Printf("First Scrolling to bottom...")

	// Click the first job
	if err := utils.ClickJob(&driver, jobName); err != nil {
		log.Fatal("Error:", err)
	}

	// Wait untill the job page is loaded

	time.Sleep(15 * time.Second)
}

func EndTheProgram(service *selenium.Service) {
	if err := service.Stop(); err != nil {
		log.Fatal("Error:", err)
	}
}
