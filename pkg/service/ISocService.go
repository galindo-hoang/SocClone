package service

type ISocService interface {
	CreateUser()
	AddFollow()
	RemoveFollow()
	GetFollowings()
	GetFollowers()
}
