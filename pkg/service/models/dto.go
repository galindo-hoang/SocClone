package models

type PersonDto struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

type FollowDto struct {
	From string `json:"from"`
	To   string `json:"to"`
}
