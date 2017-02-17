package users

import gin "gopkg.in/gin-gonic/gin.v1"

type postRequest struct {
	EmailID   string `json:"emailId" binding:"required"`
	AvatarURL string `json:"avatarUrl"`
}

func Post(c *gin.Context) {
	var req postRequest
	if c.BindJSON(&req) == nil {

		// create user in db
		var user = User{
			EmailID:   req.EmailID,
			AvatarURL: req.AvatarURL,
		}

	}
}
