package database

func PostAddQuery() (query string) {
	query = `INSERT INTO POSTS (title, description) VALUES (?, ?);`
	return
}
