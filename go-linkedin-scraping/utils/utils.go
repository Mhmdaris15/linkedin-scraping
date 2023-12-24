package utils

import (
	"crypto/aes"
	"encoding/json"
	"fmt"
	"go-linkedin-scraping/types"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
	"github.com/tebeka/selenium"
)

func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func ExtractNumber(text string) (int, error) {
	text = strings.TrimSpace(text)
	pattern := `\d+`
	r := regexp.MustCompile(pattern)
	matches := r.FindAllString(text, -1)

	if len(matches) > 0 {
		number, err := strconv.Atoi(matches[0])
		if err != nil {
			return 0, err
		}
		return number, nil
	} else {
		fmt.Println("No numbers found")
		return 0, nil
	}
}

func SaveToCSV(jobs *[]types.Job, filename string, directory string) error {
	file, err := os.Create(fmt.Sprintf("%s/%s.csv", directory, filename))
	if err != nil {
		return err
	}

	if err := gocsv.MarshalFile(jobs, file); err != nil {
		return err
	}

	return nil
}

func SaveCookiesToJSON(driver *selenium.WebDriver, filename string, directory string) error {
	cookies, err := (*driver).GetCookies()
	if err != nil {
		return err
	}

	file, err := os.Create(fmt.Sprintf("%s/%s.json", directory, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cookies)
	if err != nil {
		return err
	}

	// Encrypted cookies
	block, err := aes.NewCipher([]byte(GoDotEnvVariable("AES_KEY")))
	if err != nil {
		return err
	}

	encryptedCookies := make([]selenium.Cookie, len(cookies))
	for i, cookie := range cookies {
		encryptedCookie := cookie
		encryptedCookie.Value = encrypt(block, cookie.Value)
		encryptedCookies[i] = encryptedCookie
	}

	file, err = os.Create(fmt.Sprintf("%s/%s_encrypted.json", directory, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	encoder = json.NewEncoder(file)
	err = encoder.Encode(encryptedCookies)
	if err != nil {
		return err
	}

	return nil
}

func LoadCookiesFromJSON(driver *selenium.WebDriver, filename string, directory string) error {
	// Load cookies from JSON file
	file, err := os.Open(fmt.Sprintf("%s/%s.json", directory, filename))
	if err != nil {
		return err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	var cookies []selenium.Cookie
	err = decoder.Decode(&cookies)
	if err != nil {
		return err
	}

	// Add cookies to driver
	for _, cookie := range cookies {
		(*driver).AddCookie(&cookie)
	}

	return nil
}
