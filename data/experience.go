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
		Role:      "Software Engineer Intern",
		Location:  "Wayzata, MN",
		Company:   "Cargill",
		StartDate: "05/2024",
		EndDate:   "08/2024",
		KeyPointList: []string{
			"Worked on an agile scrum team building out common services for internal teams, enabling user authentication and authorization, and handling more than four million API calls per month",
			"Built an application to expedite the onboarding process for these services, reducing the onboarding time for each team by at least 8 business days",
		},
	},
	{
		Role:      "Teaching Assistant",
		Location:  "Boston, MA",
		Company:   "Boston University",
		StartDate: "01/2024",
		EndDate:   "Present",
		KeyPointList: []string{
			"Designed and wrote the weekly course material for Web Application Development, teaching students how to create basic websites, progressing up to building and deploying full stack web applications with database and OAuth integration",
			"Managed a team of six undergraduate workers and led weekly lab sections for 60+ students",
			"Covered material on Bash, Git/GitHub, HTML/CSS, JavaScript, Node, React, Next.js, RESTful APIs, HTTP, MongoDB, CI/CD with GitHub Actions, and deployment on Vercel",
		},
	},
}
