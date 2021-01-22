package action

type SearchRequest struct {
	Searchword string `json:"searchword"`
	Pagination int64  `json:"pagination"`
}

type SearchResponse struct {
	Data       []Movie    `json:"data"`
	Pagination Pagination `json:"pagination"`
}

type DetailRequest struct {
	ImdbID string `json:"ImdbID"`
}

type DetailResponse struct {
	Movie DetailMovie `json:"Movie"`
}

type Movie struct {
	Title  string `json:"title"`
	Year   string `json:"year"`
	ImdbID string `json:"ImdbID"`
	Type   string `json:"type"`
	Poster string `json:"poster"`
}

type DetailMovie struct {
	Title      string   `json:"Title"`
	Year       string   `json:"Year"`
	Rated      string   `json:"Rated"`
	Released   string   `json:"Released"`
	Runtime    string   `json:"Runtime"`
	Genre      string   `json:"Genre"`
	Director   string   `json:"Director"`
	Writer     string   `json:"Writer"`
	Actors     string   `json:"Actors"`
	Plot       string   `json:"Plot"`
	Language   string   `json:"Language"`
	Country    string   `json:"Country"`
	Awards     string   `json:"Awards"`
	Poster     string   `json:"Poster"`
	Ratings    []Rating `json:"Ratings"`
	Metascore  string   `json:"Metascore"`
	ImdbRating string   `json:"imdbRating"`
	ImdbVotes  string   `json:"imdbVotes"`
	ImdbID     string   `json:"imdbID"`
	Type       string   `json:"Type"`
	DVD        string   `json:"DVD"`
	BoxOffice  string   `json:"BoxOffice"`
	Production string   `json:"Production"`
	Website    string   `json:"Website"`
}

type Rating struct {
	Source string `json:"Source"`
	Value  string `json:"Value"`
}

type Pagination struct {
	PageSize    int64 `json:"page_size"`
	CurrentPage int64 `json:"current_page"`
	TotalPage   int64 `json:"total_page"`
	TotalResult int64 `json:"total_result"`
}
