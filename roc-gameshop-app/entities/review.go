package entities

type Review struct {
	ReviewId int     //Key untuk review
	GameId   int     // Id Game yang akan direview
	UserId   int     // Id User yang melakukan review
	Score    float32 //Nilai 1 1,5 2 2,5 3 3,5 4 4,5 5
}
