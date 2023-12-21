package types

type Job struct {
	ID                        int64  `json:"id" csv:"id" bson:"_id,omitempty"`
	Title                     string `json:"title" csv:"title" bson:"title,omitempty"`
	JobURL                    string `json:"job_url" csv:"job_url" bson:"job_url,omitempty"`
	Company                   string `json:"company" csv:"company" bson:"company,omitempty"`
	CompanyLink               string `json:"company_link" csv:"company_link" bson:"company_link,omitempty"`
	Location                  string `json:"location" csv:"location" bson:"location,omitempty"`
	PublishedAt               string `json:"published_at" csv:"published_at" bson:"published_at,omitempty"`
	NumberOfApplicantsApplied int    `json:"number_of_applicants_applied" csv:"number_of_applicants_applied" bson:"number_of_applicants_applied,omitempty"`
	EmploymentDuration        string `json:"employment_duration" csv:"employment_duration" bson:"employment_duration,omitempty"`
	WorkLocationType          string `json:"work_location_type" csv:"work_location_type" bson:"work_location_type,omitempty"`
	WorkExperienceLevel       string `json:"work_experience_level" csv:"work_experience_level" bson:"work_experience_level,omitempty"`
	CompanyType               string `json:"company_type" csv:"company_type" bson:"company_type,omitempty"`
	AboutTheJob               string `json:"about_the_job" csv:"about_the_job" bson:"about_the_job,omitempty"`
}
