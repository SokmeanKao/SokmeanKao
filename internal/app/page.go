package app

type PageID int

const (
	PageDashboard PageID = iota
	PageLanguages
	PageFrameworks
	PageTools
	PageDatabases
	PageGitHub
	PageSystem
)

var Menu = []string{
	"Dashboard",
	"Languages",
	"Frameworks",
	"Tools",
	"Databases",
	"GitHub",
	"System",
}

func (id PageID) Valid() bool {
	return id >= PageDashboard && id <= PageSystem
}
