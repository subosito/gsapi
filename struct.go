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

type Entry struct {
	Name        string `json:"name"`
	Package     string `json:"package"`
	ProjectURL  string `json:"projecturl"`
	Author      string `json:"author"`
	Synopsis    string `json:"synopsis"`
	Description string `json:"description"`
}

type Result struct {
	Query string  `json:"query"`
	Hits  []Entry `json:"hits"`
}
