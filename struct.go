package gsapi

type Package struct {
	Package     string
	Name        string
	StarCount   int
	Synopsis    string
	Description string
	Imported    []string
	Imports     []string
	ProjectURL  string
	StaticRank  int
}

type Item struct {
	Index   int
	Name    string
	Package string
	Link    string
	Info    string
}

type Top struct {
	Name  string
	Info  string
	Items []Item
}
