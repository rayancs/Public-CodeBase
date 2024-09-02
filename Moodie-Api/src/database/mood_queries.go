package database

// category table migration is meant to be made first as there is a refrence here
func moodDbMigrationQuery() string {
	return `
	CREATE TABLE IF NOT EXISTS moods (
    id UUID PRIMARY KEY,
    name TEXT NOT NULL,
    description TEXT NOT NULL,
    category_id UUID NOT NULL,
    emoji TEXT NOT NULL,
    FOREIGN KEY (category_id) REFERENCES categories(id)
	);
	`
}

// id , name , description , category_id , emoji
func moodInsertQuery() string {
	return `INSERT INTO moods (id, name, description, category_id, emoji) VALUES ($1, $2, $3, $4, $5)`
}
