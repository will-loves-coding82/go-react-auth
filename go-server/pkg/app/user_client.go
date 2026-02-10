package app

import (
	"database/sql"
	"fmt"
	"goAuthExample/pkg/database"
	"log"
)

type User struct {
	Id         int    `json:"id" db:"id"`
	GoogleId   string `json:"google_id" db:"google_id"`
	Email      string `json:"email" db:"email"`
	PictureURL string `json:"picture_url" db:"picture_url"`
}

type UserClient struct {
	db database.Service
}

func NewUserClient(database database.Service) *UserClient {
	return &UserClient{
		db: database,
	}
}

func (u *UserClient) CreateNewUser(model User) (int64, error) {
	res, insertUserErr := u.db.GetConn().Exec(fmt.Sprintf(`
		INSERT INTO users(google_id, email, picture_url)  VALUES (%s, %s, %s)`),
		model.GoogleId, model.Email, model.PictureURL,
	)

	if insertUserErr != nil {
		log.Printf("Error inserting user: %v", insertUserErr)
		return 0, insertUserErr
	}

	lastId, getLastIdErr := res.LastInsertId()
	if getLastIdErr != nil {
		log.Printf("Error retrieving last insert id: %v", getLastIdErr)
		return 0, getLastIdErr
	}

	return lastId, nil
}

// GetUserByEmail retrieves a user by their email. If they don't exist the method will provision a new account
func (u *UserClient) GetUserById(id int) (User, error) {
	var user User
	log.Printf("Getting user by id: %d", id)
	getUserErr := u.db.GetConn().Get(&user, fmt.Sprintf(`
		SELECT * FROM users WHERE id = %d`, id),
	)

	if getUserErr != nil && getUserErr != sql.ErrNoRows {
		log.Printf("Error fetching user: %v", getUserErr)
		return user, getUserErr
	}

	return user, nil
}

// GetOrCreateUser retrieves a user by their email. If they don't exist the method will provision a new account
func (u *UserClient) GetUserByEmail(email string) (User, error) {
	var user User
	getUserErr := u.db.GetConn().Get(&user, fmt.Sprintf(`
		SELECT id, google_id, email, picture_url FROM users WHERE email = %s`, email),
	)

	if getUserErr != nil && getUserErr != sql.ErrNoRows {
		log.Printf("Error fetching user: %v", getUserErr)
		return user, getUserErr
	}

	return user, nil
}

// UpdateUserInfo updates an existing user's information in the database
func (u *UserClient) UpdateUserInfo(user User) error {
	_, updateErr := u.db.GetConn().Exec(fmt.Sprintf(`
		UPDATE user SET google_id = %s, picture_url = %s WHERE id = %d`, user.GoogleId, user.PictureURL, user.Id),
	)
	if updateErr != nil {
		log.Printf("Error updating user: %v", updateErr)
		return updateErr
	}
	return nil
}
