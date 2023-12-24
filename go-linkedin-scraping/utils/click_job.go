package utils

import (
	"fmt"
	"go-linkedin-scraping/types"
	"log"
	"os"
	"strconv"
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
	numJobs, err := (*driver).FindElement(selenium.ByCSSSelector, "#main > div > div.scaffold-layout__list > header > div.jobs-search-results-list__title-heading > small > div > span")
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
		// Click pagination button
		paginationButton, err := (*driver).FindElement(selenium.ByCSSSelector, fmt.Sprintf("li[data-test-pagination-page-btn='%d'] button", i+1))
		if err != nil {
			return err
		}

		paginationButton.Click()
		log.Print("Clicking pagination button...")

		time.Sleep(3 * time.Second)
		// Scroll container with class "jobs-search-results-list" to the bottom to load all the jobs
		err = ScrollToBottom(driver, "jobs-search-results-list")
		if err != nil {
			log.Fatal("Error:", err)
		}
		log.Print("Scrolling to bottom...")

		err = ScrollToTop(driver, "jobs-search-results-list")
		if err != nil {
			log.Fatal("Error:", err)
		}
		log.Print("Scrolling to top...")

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
			newJob.JobURL = fmt.Sprintf("htthttps://www.linkedin.com/jobs/view/%s", jobId)
			log.Print("Job URL: ", newJob.JobURL)

			jobs[j].LocationInView()

			jobs[j].Click()

			time.Sleep(2 * time.Second)

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

			if location, err := (*driver).FindElement(selenium.ByCSSSelector, ".mb2 .app-aware-link + span.white-space-pre"); err == nil {
				if locationText, err := location.Text(); err == nil {
					newJob.Location = locationText
					log.Print("Location: ", locationText)
				}
			}

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

			if aboutTheJob, err := (*driver).FindElement(selenium.ByCSSSelector, "#job-details > span"); err == nil {
				if aboutTheJobText, err := aboutTheJob.Text(); err == nil {
					newJob.AboutTheJob = aboutTheJobText
					// log.Print("About the job: ", aboutTheJobText)
				}
			}

			if jobTitle, err := (*driver).FindElement(selenium.ByCSSSelector, ".job-details-jobs-unified-top-card__job-title"); err == nil {
				if jobTitleText, err := jobTitle.Text(); err == nil {
					newJob.Title = jobTitleText
					log.Print("Job title: ", jobTitleText)
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
