package utils

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func ScrollToBottom(driver *selenium.WebDriver, containerClass string) error {
	scrollingScript := `
		var container = document.getElementsByClassName("%s")[0];
		options = {
			left: 0,
			top: container.scrollHeight,
			behavior: "smooth",
		}
		
		// container.scrollTo(0, container.scrollHeight);
		container.scrollTo(options);
	`

	// scroll to the bottom of the container
	if _, err := (*driver).ExecuteScript(fmt.Sprintf(scrollingScript, containerClass), nil); err != nil {
		return err
	}

	return nil
}

func ScrollToTop(driver *selenium.WebDriver, containerClass string) error {
	scrollingScript := `
		var container = document.getElementsByClassName("%s")[0];
		options = {
			left: 0,
			top: 0,
			behavior: "instant",
		}
		container.scrollTo(options);
		// container.scrollTo(0, 0);
	`

	// scroll to the bottom of the container
	if _, err := (*driver).ExecuteScript(fmt.Sprintf(scrollingScript, containerClass), nil); err != nil {
		return err
	}

	return nil
}
