package utils

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func SearchJob(driver *selenium.WebDriver, job string) error {
	// Navigate to the search page
	if err := (*driver).Get(fmt.Sprintf("https://www.linkedin.com/jobs/search/?keywords=%s", job)); err != nil {
		return err
	}

	return nil
}
