package database

// user schemas & query are here nothing else these are called during  a migraiton
// they are kept lower case as they are only accessible to the db layer its not out pointing
func userTable() string {
	return `CREATE TABLE IF NOT EXISTS users (id UUID PRIMARY KEY , email TEXT UNIQUE NOT NULL) ;`
}

func userInsertQuery() string   { return `INSERT INTO  users(id,email)VALUES($1,$2);` }
func findUserByIdQuery() string { return `SELECT * FROM users WHERE email = $1` }
