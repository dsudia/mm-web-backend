package routes

import (
	"encoding/json"

	"golang.org/x/crypto/bcrypt"
)

// CreateEducatorRoute is a route handler for creating a new educator
func CreateEducatorRoute(c *gin.Context) {
	var ne *NewEducator
	c.BindJSON(&ne)

	h := bcrypt.GenerateFromPassword(ne.Password, 10)

	ed := db.Educator{
		DisplayName: ne.DisplayName,
		FirstName:   ne.FirstName,
		LastName:    ne.LastName,
		Email:       ne.Email,
		Password:    h,
		Active:      true,
		AvatarURL:   ne.AvatarURL,
	}

	e, err := db.CreateEducator(ed)
	if err != nil {
		c.JSON(400, gin.H{
			status:  400,
			message: "There was an error inserting the educator into the database.",
		})
	}

	e, _ = json.Marshal(e)
	es := string(e)
	c.JSON(200, gin.H{
		status:  200,
		message: es,
	})
}
