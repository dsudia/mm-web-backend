package main

// NewEducator represents JSON post of incoming new educator user
type NewEducator struct {
	DisplayName string `json:"displayName"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
	AvatarURL   string `json:"avatarURL"`
}

// NewSchool represents JSON post of incoming new school user
type NewSchool struct {
	ID          int    `orm:"column(id)" json:"id"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
	AvatarURL   string `json:"avatarURL"`
}
