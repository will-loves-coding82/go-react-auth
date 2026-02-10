package app

const insertUserQuery = `
INSERT INTO users(google_id, email, picture_url) 
VALUES (?, ?, ?);
`

const updateUserQuery = `
UPDATE users
SET google_id = ?, picture_url = ?
WHERE id = ?;
`

const getUserByIdQuery = `
SELECT id, google_id, email, picture_url
FROM users
WHERE id = ?;
`

const getUserByEmailQuery = `
SELECT *
FROM users
WHERE email = ?;
`
