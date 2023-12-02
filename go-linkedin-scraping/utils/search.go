package utils

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func SearchJob(driver selenium.WebDriver, job string) (selenium.WebDriver, error) {
	// Navigate to the search page
	if err := driver.Get(fmt.Sprintf("https://www.linkedin.com/jobs/search/?keywords=%s", job)); err != nil {
		return nil, err
	}

	return driver, nil
}
