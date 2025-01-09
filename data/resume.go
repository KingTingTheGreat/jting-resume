package data

import "time"

type Resume struct {
	Date           string
	Header         Header
	ExperienceList []Experience
	ProjectList    []Project
	EducationList  []Education
	Skills         Skills
}

var ResumeData = Resume{
	Date:           time.Now().Format("01-02-2006"),
	Header:         HeaderData,
	ExperienceList: ExperienceData,
	ProjectList:    ProjectData,
	EducationList:  EducationData,
	Skills:         SkillsData,
}
