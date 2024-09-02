package database

// caT QUERIES
// migration or table creation
func catDbMigrartion() string {
	return `CREATE TABLE IF NOT EXISTS categories (
		id UUID PRIMARY KEY,
		name TEXT NOT NULL,
		color TEXT NOT NULL
	);`

}

// insert query,
// takes in (id , name , color)
func catDbInsertQuery() string {
	return `INSERT INTO categories (id, name, color) VALUES ($1, $2, $3)`
}
