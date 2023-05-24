package models

type Baby struct {
	BabyID     rune   `bson:"baby_id" json:"baby_id"`
	Nickname   string `bson:"nickname" json:"nickname"`
	PictureUrl string `bson:"picture_url" json:"picture_url"`
}

type BabyList struct {
	Count   int     `json:"count"`
	Results []*Baby `json:"results"`
}
