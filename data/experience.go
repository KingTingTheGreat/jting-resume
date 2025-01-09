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
			"Led weekly lab sections for 60+ students, held office hours, and managed a team of 6 undergraduate workers",
			"Covered material on Git/GitHub, HTML/CSS, JavaScript, Node, React, NextJS, RESTful APIs, HTTP, MongoDB, CI/CD with GitHub Actions, and deployment on Vercel",
		},
	},
	{
		Role:      "IT Intern",
		Location:  "Garden City, NY",
		Company:   "Lavner Education",
		StartDate: "06/2023",
		EndDate:   "08/2023",
		KeyPointList: []string{
			"Introduced students to various programming languages, such as Python, C++, Java, Bash, HTML/CSS, and JavaScript",
			"Led classes on programming basics, terminal usage, Git/GitHub, web app development, and game development",
			"Adapted curricula and teaching methods to accommodate individual students’ needs",
		},
	},
}
