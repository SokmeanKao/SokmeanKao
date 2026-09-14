package profile

type Skill struct {
	Name    string
	Status  string
	Percent int
}

type NamedStatus struct {
	Name   string
	Status string
}

type Profile struct {
	Name       string
	Role       string
	Location   string
	Status     string
	GitHubUser string
	Host       string
	Languages  []Skill
	Frameworks []NamedStatus
	Tools      []string
	Databases  []NamedStatus
	Editors    []string
	Browsers   []string
	Modes      []string
}

func Default() Profile {
	return Profile{
		Name:       "Sokmean",
		Role:       "Full Stack Developer",
		Location:   "Cambodia",
		Status:     "ONLINE",
		GitHubUser: "sokmeankao",
		Host:       "MSI Laptop",
		Languages: []Skill{
			{Name: "Java", Status: "Active", Percent: 90},
			{Name: "JavaScript", Status: "Active", Percent: 80},
			{Name: "HTML/CSS", Status: "Active", Percent: 85},
		},
		Frameworks: []NamedStatus{
			{Name: "Spring", Status: "Ready"},
			{Name: "Spring Boot", Status: "Ready"},
			{Name: "Next.js", Status: "Ready"},
			{Name: "Docker", Status: "Ready"},
			{Name: "JWT", Status: "Ready"},
			{Name: "Yarn", Status: "Ready"},
		},
		Tools: []string{
			"GitHub", "Notion", "DigitalOcean", "Google Cloud",
			"Vercel", "Hostinger", "RabbitMQ", "Jenkins",
		},
		Databases: []NamedStatus{
			{Name: "PostgreSQL", Status: "READY"},
			{Name: "MySQL", Status: "READY"},
		},
		Editors:  []string{"IntelliJ IDEA", "VS Code", "CodeSandbox"},
		Browsers: []string{"Brave", "Chrome"},
		Modes:    []string{"Building systems", "Learning", "Shipping software"},
	}
}
