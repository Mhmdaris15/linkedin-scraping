package utils

import (
	"log"

	"github.com/tebeka/selenium"
)

func LoginLinkedIn(driver *selenium.WebDriver) error {
	// navigate to the login page
	err := (*driver).Get("https://www.linkedin.com/home")
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	// find the username and password fields
	username, err := (*driver).FindElement(selenium.ByID, "session_key")
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}
	password, err := (*driver).FindElement(selenium.ByID, "session_password")
	if err != nil {
		log.Fatal("Error:", err)
		return err

	}

	// enter the username and password
	username.SendKeys(GoDotEnvVariable("EMAIL_SYSTEM"))
	password.SendKeys(GoDotEnvVariable("PASSWORD_SYSTEM"))

	// find the login button
	login, err := (*driver).FindElement(selenium.ByCSSSelector, ".sign-in-form__submit-btn--full-width")
	if err != nil {
		log.Fatal("Error:", err)
		return err
	}

	// click the login button
	login.Click()

	return nil
}
