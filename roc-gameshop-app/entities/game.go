package entities

type Game struct {
	GameId      int
	Name        string
	Description string
	Genre       string
	SalePrice   float64
	RentalPrice float64
	Studio      string
	Stock       int
	IsDeleted   bool
}
