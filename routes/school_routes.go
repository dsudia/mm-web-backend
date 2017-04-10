package routes

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

// CreateSchoolRoute is a route handler for creating a school
func CreateSchoolRoute(c *gin.Context) {
	var ns *NewSchool
	c.BindJSON(&ne)

	h := bcrypt.GenerateFromPassword(ne.Password, 10)

	s := db.School{
		DisplayName: ns.DisplayName,
		Name:        ns.Name,
		Email:       ns.Email,
		Password:    h,
		Active:      true,
		AvatarURL:   ns.AvatarURL,
	}

	s, err := db.CreateSchool(ed)
	if err != nil {
		c.JSON(400, gin.H{
			status:  400,
			message: "There was an error inserting the school into the database.",
		})
	}

	s, _ = json.Marshal(e)
	ss := string(e)
	c.JSON(200, gin.H{
		status:  200,
		message: ss,
	})
}
