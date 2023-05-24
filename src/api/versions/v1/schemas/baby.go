package schemas

type CreateBabySchema struct {
	Nickname string
	Photo    []byte
}

type ClickBabySchema struct {
	BabyID string
}
