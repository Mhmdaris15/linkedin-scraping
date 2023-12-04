package utils

import (
	"fmt"
	"go-linkedin-scraping/types"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/joho/godotenv"
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
