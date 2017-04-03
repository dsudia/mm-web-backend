package main

// NewEducator represents JSON post of incoming new user
type NewEducator struct {
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
	AvatarURL   string `json:"avatarURL"`
}
