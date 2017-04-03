package main

// Educator defines a database model for an educator user
type Educator struct {
	ID           int    `json:"id"`
	DisplayName  string `json:"displayName"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Active       bool   `json:"active"`
	AvatarURL    string `json:"avatarURL"`
	Description  string `json:"description"`
	AgeRanges    []int  `json:"ageRanges"`
	AgeRangesWgt int    `json:"ageRangesWgt"`
	Cals         []int  `json:"cals"`
	CalsWgt      int    `json:"calsWgt"`
	OrgTypes     []int  `json:"orgTypes"`
	OrgTypesWgt  int    `json:"orgTypesWgt"`
	LocTypes     []int  `json:"locTypes"`
	LocTypesWgt  int    `json:"locTypesWgt"`
	EdTypes      []int  `json:"edTypes"`
	EdTypesWgt   int    `json:"edTypesWgt"`
	Sizes        []int  `json:"sizes"`
	SizesWgt     []int  `json:"sizesWgt"`
	Trainings    []int  `json:"trainings"`
	TrainingsWgt []int  `json:"trainingsWgt"`
	Traits       []int  `json:"traits"`
	TraitsWgt    int    `json:"traitsWgt"`
	States       []int  `json:"states"`
	StatesWgt    int    `json:"statesWgt"`
}

// School defines a database model for a school user
type School struct {
	ID          int    `json:"id"`
	DisplayName string `json:"displayName"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Active      bool   `json:"active"`
	AvatarURL   string `json:"avatarURL"`
	Description string `json:"description"`
}

// SchoolMatchingProfile defines a database model for a school's matching Profile
type SchoolMatchingProfile struct {
	ID           int     `json:"id"`
	School       *School `orm:"rel(fk)"`
	Active       bool    `json:"active"`
	AgeRanges    []int   `json:"ageRanges"`
	AgeRangesWgt int     `json:"ageRangesWgt"`
	Cals         []int   `json:"cals"`
	CalsWgt      int     `json:"calsWgt"`
	OrgTypes     []int   `json:"orgTypes"`
	OrgTypesWgt  int     `json:"orgTypesWgt"`
	LocTypes     []int   `json:"locTypes"`
	LocTypesWgt  int     `json:"locTypesWgt"`
	EdTypes      []int   `json:"edTypes"`
	EdTypesWgt   int     `json:"edTypesWgt"`
	Sizes        []int   `json:"sizes"`
	SizesWgt     []int   `json:"sizesWgt"`
	Trainings    []int   `json:"trainings"`
	TrainingsWgt []int   `json:"trainingsWgt"`
	Traits       []int   `json:"traits"`
	TraitsWgt    int     `json:"traitsWgt"`
	States       []int   `json:"states"`
	StatesWgt    int     `json:"statesWgt"`
}

// Match defines the database model for a match between users
type Match struct {
	ID                      int                   `json:"id"`
	SchoolMatchingProfileID *SchooMatchingProfile `orm:"rel(fk)" json:"schoolMatchingProfileId"`
	EducatorID              *Educator             `orm:"rel(fk)" json:"educatorId"`
	Percentage              int                   `json:"percentage"`
	EducatorConfirmation    bool                  `json:"educatorConfirmation"`
	SchoolConfirmation      bool                  `json:"schoolConfirmation"`
	EducatorDenial          bool                  `json:"educatorDenial"`
	SchoolDenial            bool                  `json:"schoolDenial"`
}
