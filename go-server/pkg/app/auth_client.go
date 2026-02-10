package app

import (
	"database/sql"
	"goAuthExample/pkg/database"
	"log"

	"github.com/markbates/goth"
)

type AuthClient struct {
	db database.Service
}

func NewAuthClient(database database.Service) *AuthClient {
	return &AuthClient{
		db: database,
	}
}

// GetOrCreateUser will retrieve the id of user with a specific email. If they don't
// exist the method will provision a new profile and return the newly created id.
func (a *AuthClient) GetOrCreateUser(gothUser goth.User) (int, error) {
	var user User

	log.Printf("Getting user by email: %s", gothUser.Email)
	getUserErr := a.db.GetConn().Get(&user, getUserByEmailQuery, gothUser.Email)
	if getUserErr != nil && getUserErr != sql.ErrNoRows {
		log.Printf("Error fetching user by email: %v", getUserErr)
		return user.Id, getUserErr
	}

	if getUserErr == sql.ErrNoRows {
		log.Printf("User didn't exist. Creating new user with email: %s", gothUser.Email)
		res, insertUserErr := a.db.GetConn().Exec(insertUserQuery, gothUser.UserID, gothUser.Email, gothUser.AvatarURL)
		if insertUserErr != nil {
			log.Printf("Error inserting user: %v", insertUserErr)
			return 0, insertUserErr
		}

		lastId, getLastIdErr := res.LastInsertId()
		if getLastIdErr != nil {
			log.Printf("Error retrieving last insert id: %v", getLastIdErr)
			return 0, getLastIdErr
		}

		log.Printf("Created new user with id: %d", lastId)
		return int(lastId), nil
	}

	return user.Id, nil
}

func (a *AuthClient) UpdateUserDetails(googleId string, pictureURL string, userId int) error {
	_, updateErr := a.db.GetConn().Exec(updateUserQuery, googleId, pictureURL, userId)
	if updateErr != nil {
		log.Printf("Error updating user: %v", updateErr)
		return updateErr
	}
	return nil
}
