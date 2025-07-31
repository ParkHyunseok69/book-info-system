package model

type Book struct {
	ID            int32
	Title         string
	Author        string
	Summary       string
	Genre         string
	PublicateYear string
	Pages         int32
	DateAcquired  string
}

type Author struct {
	FirstName   string
	LastName    string
	BirthYear   string
	Nationality string
	CreatedAt   string
}
