package database

func CreateUserMoodQuery() string {
	return `
	CREATE TABLE IF NOT EXISTS user_mood(
	id UUID PRIMARY KEY ,
	user_id UUID NOT NULL ,
	note TEXT NOT NULL ,
	mood_id UUID NOT NULL ,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP , 
	FOREIGN KEY (user_id) REFERENCES users(id),
	FOREIGN KEY (mood_id) REFERENCES moods(id)
	)
	`
}

// pK(id uuid), fk(user_id uuid) ,note string ,fk(mood_id uuid)
func InsertUserMoodQuery() string {
	return ` 
		INSERT INTO user_mood(
		id , user_id , note ,mood_id
		)VALUES($1,$2,$3,$4)
	`
}
