package main

// Educator defines a database model for an educator user
type Educator struct {
	ID           int
	DisplayName  string
	FirstName    string
	LastName     string
	Email        string
	Password     string
	Active       bool
	AvatarURL    string
	Description  string
	AgeRanges    []int
	AgeRangesWgt int
	Cals         []int
	CalsWgt      int
	OrgTypes     []int
	OrgTypesWgt  int
	LocTypes     []int
	LocTypesWgt  int
	EdTypes      []int
	EdTypesWgt   int
	Sizes        []int
	SizesWgt     []int
	Trainings    []int
	TrainingsWgt []int
	Traits       []int
	TraitsWgt    int
	States       []int
	StatesWgt    int
}

// School defines a database model for a school user
type School struct {
	ID          int
	DisplayName string
	Name        string
	Email       string
	Password    string
	Active      bool
	AvatarURL   string
	Description string
}

// SchoolMatchingProfile defines a database model for a school's matching Profile
type SchoolMatchingProfile struct {
	ID           int
	School       *School `orm:"rel(fk)"`
	Active       bool
	AgeRanges    []int
	AgeRangesWgt int
	Cals         []int
	CalsWgt      int
	OrgTypes     []int
	OrgTypesWgt  int
	LocTypes     []int
	LocTypesWgt  int
	EdTypes      []int
	EdTypesWgt   int
	Sizes        []int
	SizesWgt     []int
	Trainings    []int
	TrainingsWgt []int
	Traits       []int
	TraitsWgt    int
	States       []int
	StatesWgt    int
}

// Match defines the database model for a match between users
type Match struct {
	ID                      int
	SchoolMatchingProfileID *SchooMatchingProfile `orm:"rel(fk)"`
	EducatorID              *Educator             `orm:"rel(fk)"`
	Percentage              int
	EducatorConfirmation    bool
	SchoolConfirmation      bool
	EducatorDenial          bool
	SchoolDenial            bool
}
