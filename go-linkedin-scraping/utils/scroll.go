package utils

import (
	"fmt"

	"github.com/tebeka/selenium"
)

func ScrollToBottom(driver selenium.WebDriver, containerClass string) error {
	// find the container
	// container, err := driver.FindElement(selenium.ByClassName, containerClass)
	// if err != nil {
	// 	return err
	// }

	// Wait untill the last job list is loaded

	scrollingScript := `
		var container = document.getElementsByClassName("%s")[0];
		container.scrollTo(0, container.scrollHeight);
	`

	// scroll to the bottom of the container
	if _, err := driver.ExecuteScript(fmt.Sprintf(scrollingScript, containerClass), nil); err != nil {
		return err
	}

	return nil
}
