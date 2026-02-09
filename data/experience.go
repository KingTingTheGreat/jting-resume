package data

type Experience struct {
	Role         string
	Location     string
	Company      string
	StartDate    string
	EndDate      string
	KeyPointList []string
}

var ExperienceData = []Experience{
	{
		Role:      "Software Engineer",
		Location:  "Atlanta, GA",
		Company:   "Cargill",
		StartDate: "06/2025",
		EndDate:   "Present",
		KeyPointList: []string{
			"Developed an application to track the migration of 700+ applications from on-prem data centers to the cloud, saving hundreds of thousands of dollars in operational, development, and maintenance costs",
			"Implemented internal CSS and MUI frontend libraries to ensure accessibility and brand consistency across the organization's thousands of internal and customer-facing applications",
		},
	},
	{
		Role:      "Teaching Assistant",
		Location:  "Boston, MA",
		Company:   "Boston University",
		StartDate: "01/2024",
		EndDate:   "05/2025",
		KeyPointList: []string{
			"Designed and wrote the weekly course material for Web Application Development, teaching students how to create basic websites and progressing up to building and deploying full stack web applications with database and OAuth integration",
			"Managed a team of six undergraduate workers and led weekly lab sections for 60+ students",
		},
	},
	{
		Role:      "Software Engineer Intern",
		Location:  "Wayzata, MN",
		Company:   "Cargill",
		StartDate: "05/2024",
		EndDate:   "08/2024",
		KeyPointList: []string{
			"Built an application to expedite the onboarding process for authentication and authorization services, reducing the onboarding time for each team by at least 8 business days",
		},
	},
}
