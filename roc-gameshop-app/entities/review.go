package entities

type ReviewPerGame struct {
	ReviewId  int
	Rating    float64
	ReviewMsg string
	UserName  string
}

type Review struct {
	ReviewId  int
	UserId    int
	GameId    int
	Rating    float64
	ReviewMsg string
}
