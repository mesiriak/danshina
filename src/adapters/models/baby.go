package models

type BabyModel struct {
	Uuid       string `bson:"uuid" json:"uuid"`
	Nickname   string `bson:"nickname" json:"nickname"`
	PictureUrl string `bson:"picture_url" json:"picture_url"`
	Counter    uint   `bson:"counter" json:"counter"`
}

type BabyListModel struct {
	Count   int          `json:"count"`
	Results []*BabyModel `json:"results"`
}

type BabyRandomListModel struct {
	Size    int          `json:"size"`
	Results []*BabyModel `json:"results"`
}

type CreateBabyModel struct {
	Nickname string  `json:"nickname" validate:"required,min=3,max=32"`
	Picture  []*byte `json:"picture"`
}
