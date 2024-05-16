package utils

import (
	"fmt"
	"go-linkedin-scraping/types"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ClickJob(driver *selenium.WebDriver, jobName string) error {
	// find the job
	var jobsThisProfession []types.Job
	var numJobsInt int
	jobCollection := GetCollection(DB, "Jobs")

	// Get number of jobs
	numJobs, err := (*driver).FindElement(selenium.ByCSSSelector, "#main > div > div.scaffold-layout__list-detail-inner.scaffold-layout__list-detail-inner--grow > div.scaffold-layout__list > header > div.jobs-search-results-list__title-heading > small > div > span")
	if err != nil {
		log.Fatal("Error:", err)
	} else {
		if numJobsText, err := numJobs.Text(); err == nil {
			numJobsNum, err := ExtractNumber(numJobsText)
			if err != nil {
				log.Fatal("Error:", err)
			}
			log.Print("Number of jobs: ", numJobsNum)
			numJobsInt = numJobsNum
		}
	}

	numTabs := numJobsInt / 25
	log.Print("Number of tabs: ", numTabs)

	for i := 0; i < numTabs; i++ {
		log.Print("Tab: ", i+1)
		var jobsThisTab []types.Job

		// Scroll container with class "jobs-search-results-list" to the bottom to load all the jobs
		err = WaitForJobsAndClickFirst(driver, jobName)
		if err != nil {
			return err
		}

		paginations, err := (*driver).FindElements(selenium.ByCSSSelector, "li.artdeco-pagination__indicator")
		if err != nil {
			return err
		}

		// Click pagination button
		paginationButton, err := (*driver).FindElement(selenium.ByCSSSelector, fmt.Sprintf("li[data-test-pagination-page-btn='%d'] button", i+1))
		if err != nil {
			// Click the last minus one pagination button
			paginations[len(paginations)-2].Click()
			log.Print("Clicking the last minus one pagination button...")
		} else {
			paginationButton.Click()
			log.Print("Clicking pagination button...")
		}

		time.Sleep(3 * time.Second)

		jobs, err := (*driver).FindElements(selenium.ByCSSSelector, ".jobs-search-results__list-item")
		if err != nil {
			return err
		}
		log.Print("Number of jobs in single pagination: ", len(jobs))

		for j := 0; j < len(jobs); j++ {
			log.Print("Job: ", j)
			var newJob types.Job
			jobId, err := jobs[j].GetAttribute("data-occludable-job-id")
			if err != nil {
				log.Print("Error: ", err)
			} else {
				jobIdInt, err := strconv.Atoi(jobId)
				if err != nil {
					log.Print("Error: ", err)
				}
				newJob.ID = int64(jobIdInt)
			}
			newJob.JobURL = fmt.Sprintf("https://www.linkedin.com/jobs/view/%s", jobId)
			log.Print("Job URL: ", newJob.JobURL)

			jobs[j].LocationInView()

			for i := 0; i < 5; i++ {
				jobs[j].Click()
			}

			err = ScrollToBottom(driver, "div.scaffold-layout__detail.overflow-x-hidden.jobs-search__job-details > div")
			if err != nil {
				return err
			}

			time.Sleep(1 * time.Second)

			for {
				if err := extractAboutTheJob(driver, &newJob); err != nil {
					log.Print("Error: ", err)
					jobs[j].Click()
					time.Sleep(2 * time.Second)
				} else {
					break
				}
			}

			if companyDescription, err := (*driver).FindElement(selenium.ByCSSSelector, "div.jobs-company__box > p > div"); err == nil {
				companyDescription.LocationInView()
				showMore, err := companyDescription.FindElement(selenium.ByCSSSelector, "button.inline-show-more-text__button")
				if err == nil {
					showMore.Click()
					log.Print("Clicking show more button...")
				}
			}

			time.Sleep(500 * time.Millisecond)

			jobDescriptionContainer, err := (*driver).FindElement(selenium.ByCSSSelector, ".job-details-jobs-unified-top-card__primary-description-container")
			if err != nil {
				return err
			}

			if jobDescriptionContainerText, err := jobDescriptionContainer.Text(); err == nil {
				log.Print("Job description container text: ", jobDescriptionContainerText)
			}

			if companyName, err := jobDescriptionContainer.FindElement(selenium.ByCSSSelector, "a.app-aware-link"); err == nil {
				if companyNameText, err := companyName.Text(); err == nil {
					newJob.Company = companyNameText
					log.Print("Company name: ", companyNameText)
				}
			}

			if companyUrl, err := jobDescriptionContainer.FindElement(selenium.ByCSSSelector, "a.app-aware-link"); err == nil {
				if companyUrlText, err := companyUrl.GetAttribute("href"); err == nil {
					newJob.CompanyLink = companyUrlText
					log.Print("Company URL: ", companyUrlText)
				}
			}

			if publishedAt, err := (*driver).FindElement(selenium.ByCSSSelector, "div.job-details-jobs-unified-top-card__content--two-pane > div.job-details-jobs-unified-top-card__primary-description-container > div > span:nth-child(4) > span"); err == nil {
				if publishedAtText, err := publishedAt.Text(); err == nil {
					newJob.PublishedAt = publishedAtText
					log.Print("Published at: ", publishedAtText)
				}
			}

			if numberOfApplicantsApplied, err := (*driver).FindElement(selenium.ByCSSSelector, "div.job-details-jobs-unified-top-card__primary-description-container > div > span:nth-child(6)"); err == nil {
				if numberOfApplicantsAppliedText, err := numberOfApplicantsApplied.Text(); err == nil {
					numberOfApplicantsAppliedNumber, err := ExtractNumber(numberOfApplicantsAppliedText)
					if err != nil {
						return err
					}
					newJob.NumberOfApplicantsApplied = numberOfApplicantsAppliedNumber
					log.Print("Number of applicants applied: ", numberOfApplicantsAppliedText)
				}
			}

			newJob.Location = extractLocation(driver)

			if jobType, err := (*driver).FindElement(selenium.ByCSSSelector, "div.relative.job-details-jobs-unified-top-card__container--two-pane > div.job-details-jobs-unified-top-card__content--two-pane > div.mt3.mb2 > ul > li:nth-child(1) > span > span:nth-child(1)"); err == nil {
				if jobTypeText, err := jobType.Text(); err == nil {
					newJob.WorkLocationType = jobTypeText
					log.Print("Job type: ", jobTypeText)
				}
			}

			if jobTimeType, err := (*driver).FindElement(selenium.ByCSSSelector, "div.relative.job-details-jobs-unified-top-card__container--two-pane > div.job-details-jobs-unified-top-card__content--two-pane > div.mt3.mb2 > ul > li:nth-child(1) > span > span:nth-child(2) > span > span:nth-child(1)"); err == nil {
				if jobTimeTypeText, err := jobTimeType.Text(); err == nil {
					log.Print("Job time type: ", jobTimeTypeText)
				}
			}

			if companyType, err := (*driver).FindElement(selenium.ByCSSSelector, "div.relative.job-details-jobs-unified-top-card__container--two-pane > div.job-details-jobs-unified-top-card__content--two-pane > div.mt3.mb2 > ul > li:nth-child(2) > span"); err == nil {
				if companyTypeText, err := companyType.Text(); err == nil {
					newJob.CompanyType = companyTypeText
					log.Print("Company type: ", companyTypeText)
				}
			}

			if jobTitle, err := (*driver).FindElement(selenium.ByCSSSelector, ".job-details-jobs-unified-top-card__job-title"); err == nil {
				if jobTitleText, err := jobTitle.Text(); err == nil {
					newJob.Title = jobTitleText
					log.Print("Job title: ", jobTitleText)
				}
			}

			if companyLogo, err := (*driver).FindElement(selenium.ByCSSSelector, "div.jobs-company__box > div.display-flex.align-items-center.mt5 > div > div > a > img"); err == nil {
				if companyLogoText, err := companyLogo.GetAttribute("src"); err == nil {
					newJob.CompanyLogo = companyLogoText
					log.Print("Company logo: ", companyLogoText)
				}
			}

			if companyDescription, err := (*driver).FindElement(selenium.ByCSSSelector, "div.jobs-company__box > p > div"); err == nil {
				if companyDescriptionText, err := companyDescription.Text(); err == nil {
					newJob.CompanyDescription = companyDescriptionText
					log.Print("Company description: ", companyDescriptionText)
				}
			}

			jobsThisTab = append(jobsThisTab, newJob)
		}

		if os.Mkdir(fmt.Sprintf("./data/%s", jobName), 0777); err != nil {
			log.Print("Error: ", err)
		}
		// Convert jobsThisTab to []interface{}
		var jobsInterface []interface{}
		for _, job := range jobsThisTab {
			jobsInterface = append(jobsInterface, job)
		}

		if os.Mkdir(fmt.Sprintf("./data/%s/pagination", jobName), 0777); err != nil {
			log.Print("Error: ", err)
		}

		// Save to csv file
		if err := SaveToCSV(&jobsThisTab, fmt.Sprintf("%s-pagination-%d", jobName, i+1), fmt.Sprintf("./data/%s/pagination", jobName)); err != nil {
			return err
		}

		// Insert to database
		if _, err := jobCollection.InsertMany(ctx, jobsInterface); err != nil {
			if writeException, ok := err.(mongo.BulkWriteException); ok {
				for _, writeError := range writeException.WriteErrors {
					index := writeError.Index
					duplicateJob := jobsInterface[index].(types.Job) // Assert the type of duplicateJob to the appropriate struct type
					filter := bson.M{"_id": duplicateJob.ID}

					_, updateErr := jobCollection.ReplaceOne(ctx, filter, duplicateJob)
					if updateErr != nil {
						// Handle the update error appropriately, e.g., return it or log it with more context
						return fmt.Errorf("failed to update duplicate document: %w", updateErr)
					}
				}
			} else {
				// Handle other errors
				log.Print("Error: ", err)
				return err // Return the error for further handling
			}
		}

		jobsThisProfession = append(jobsThisProfession, jobsThisTab...)

	}

	if os.Mkdir(fmt.Sprintf("./data/%s", jobName), 0777); err != nil {
		log.Print("Error: ", err)
	}

	// Save to csv file
	if err := SaveToCSV(&jobsThisProfession, jobName, fmt.Sprintf("./data/%s", jobName)); err != nil {
		return err
	}

	return nil
}

func extractAboutTheJob(driver *selenium.WebDriver, newJob *types.Job) error {
	// Find the element containing the job details
	aboutTheJob, err := (*driver).FindElement(selenium.ByCSSSelector, "#job-details > div")
	if err != nil {
		return err
	}

	// Extract the text content of the element
	aboutTheJobText, err := aboutTheJob.Text()
	if err != nil {
		return err
	}

	// Assign the extracted text to the Job object
	newJob.AboutTheJob = aboutTheJobText

	return nil
}

func WaitForJobsAndClickFirst(driver *selenium.WebDriver, jobName string) error {
	// Wait for the last job list to load with a timeout
	err := (*driver).WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		lastJobList, _ := wd.FindElements(selenium.ByCSSSelector, ".jobs-search-results__list-item:nth-child(25)")

		if lastJobList != nil {
			return true, nil
		}

		log.Print("Waiting for the last job list to load...")
		return false, nil
	}, 10*time.Second)

	if err != nil {
		return err
	}

	// Scroll to the bottom to ensure all jobs are loaded
	err = ScrollToBottom(driver, ".jobs-search-results-list")
	if err != nil {
		return err
	}
	log.Printf("First Scrolling to bottom...")

	err = ScrollToTop(driver, ".jobs-search-results-list")
	if err != nil {
		log.Fatal("Error:", err)
	}
	log.Print("Scrolling to top...")

	return nil
}

func extractLocation(driver *selenium.WebDriver) string {

	topCard, err := (*driver).FindElement(selenium.ByCSSSelector, "div.job-details-jobs-unified-top-card__primary-description-container > div")
	if err != nil {
		log.Fatal("Error:", err)
		return ""
	}

	topCardText, err := topCard.Text()
	if err != nil {
		log.Fatal("Error:", err)
		return ""
	}

	separatedText := strings.Split(topCardText, "·")
	locationText := strings.Trim(separatedText[1], " ")

	return locationText
}
