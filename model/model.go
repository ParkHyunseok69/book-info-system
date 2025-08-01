package model

type Book struct {
	ID            int32  `json:"id"`
	Title         string `json:"title"`
	Author        string `json:"author"`
	Summary       string `json:"summary"`
	Genre         string `json:"genre"`
	PublicateYear string `json:"publicate_year"`
	Pages         int32  `json:"pages"`
	DateAcquired  string `json:"date_acquired"`
}

type Author struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	BirthYear   string `json:"birth_year"`
	Nationality string `json:"nationality"`
	CreatedAt   string `json:"created_at"`
}
