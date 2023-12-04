package utils

import (
	"log"
	"time"

	"github.com/tebeka/selenium"
)

func ClickJob(driver selenium.WebDriver) (selenium.WebDriver, error) {
	// find the job
	jobs, err := driver.FindElements(selenium.ByCSSSelector, ".jobs-search-results__list-item")
	if err != nil {
		return driver, err
	}

	for i := 0; i < len(jobs); i++ {
		log.Print("Job: ", i)

		jobs[i].Click()

		time.Sleep(2 * time.Second)

		jobDescriptionContainer, err := driver.FindElement(selenium.ByCSSSelector, ".job-details-jobs-unified-top-card__primary-description-container")
		if err != nil {
			return driver, err
		}

		if jobDescriptionContainerText, err := jobDescriptionContainer.Text(); err == nil {
			log.Print("Job description container text: ", jobDescriptionContainerText)
		}

		if jobDescriptionContainerInnerHTML, err := jobDescriptionContainer.GetAttribute("innerHTML"); err == nil {
			log.Print("Job description container inner HTML: ", jobDescriptionContainerInnerHTML)
		}

		if companyName, err := jobDescriptionContainer.FindElement(selenium.ByCSSSelector, "a.app-aware-link"); err == nil {
			if companyNameText, err := companyName.Text(); err == nil {
				log.Print("Company name: ", companyNameText)
			}
		}

		if companyUrl, err := jobDescriptionContainer.FindElement(selenium.ByCSSSelector, "a.app-aware-link"); err == nil {
			if companyUrlText, err := companyUrl.GetAttribute("href"); err == nil {
				log.Print("Company URL: ", companyUrlText)
			}
		}

		if jobType, err := driver.FindElement(selenium.ByCSSSelector, "div.relative.job-details-jobs-unified-top-card__container--two-pane > div.job-details-jobs-unified-top-card__content--two-pane > div.mt3.mb2 > ul > li:nth-child(1) > span > span:nth-child(1)"); err == nil {
			if jobTypeText, err := jobType.Text(); err == nil {
				log.Print("Job type: ", jobTypeText)
			}
		}

		if jobTimeType, err := driver.FindElement(selenium.ByCSSSelector, "div.relative.job-details-jobs-unified-top-card__container--two-pane > div.job-details-jobs-unified-top-card__content--two-pane > div.mt3.mb2 > ul > li:nth-child(1) > span > span:nth-child(2) > span > span:nth-child(1)"); err == nil {
			if jobTimeTypeText, err := jobTimeType.Text(); err == nil {
				log.Print("Job time type: ", jobTimeTypeText)
			}
		}

		if companyType, err := driver.FindElement(selenium.ByCSSSelector, "div.relative.job-details-jobs-unified-top-card__container--two-pane > div.job-details-jobs-unified-top-card__content--two-pane > div.mt3.mb2 > ul > li:nth-child(2) > span"); err == nil {
			if companyTypeText, err := companyType.Text(); err == nil {
				log.Print("Company type: ", companyTypeText)
			}
		}

		if aboutTheJob, err := driver.FindElement(selenium.ByCSSSelector, "#job-details > span"); err == nil {
			if aboutTheJobText, err := aboutTheJob.Text(); err == nil {
				log.Print("About the job: ", aboutTheJobText)
			}
		}

		jobTitle, err := driver.FindElement(selenium.ByCSSSelector, ".job-details-jobs-unified-top-card__job-title")
		if err != nil {
			return driver, err
		}

		if jobTitleText, err := jobTitle.Text(); err == nil {
			log.Print("Job title: ", jobTitleText)
		}

	}

	return driver, nil
}
