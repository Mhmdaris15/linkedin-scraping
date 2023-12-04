package types

type Job struct {
	ID                        int64  `json:"id" csv:"id"`
	Title                     string `json:"title" csv:"title"`
	JobURL                    string `json:"job_url" csv:"job_url"`
	Company                   string `json:"company" csv:"company"`
	CompanyLink               string `json:"company_link" csv:"company_link"`
	Location                  string `json:"location" csv:"location"`
	PublishedAt               string `json:"published_at" csv:"published_at"`
	NumberOfApplicantsApplied int    `json:"number_of_applicants_applied" csv:"number_of_applicants_applied"`
	EmploymentDuration        string `json:"employment_duration" csv:"employment_duration"`
	WorkLocationType          string `json:"work_location_type" csv:"work_location_type"`
	WorkExperienceLevel       string `json:"work_experience_level" csv:"work_experience_level"`
	CompanyType               string `json:"company_type" csv:"company_type"`
	AboutTheJob               string `json:"about_the_job" csv:"about_the_job"`
}
