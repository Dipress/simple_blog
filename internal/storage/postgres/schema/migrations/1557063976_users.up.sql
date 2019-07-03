CREATE TABLE IF NOT EXISTS users (
	id	SERIAL PRIMARY KEY,
	username VARCHAR (50) NOT NULL,
	email	VARCHAR (50) NOT NULL,
	password_hash	VARCHAR (72) NOT NULL,

	/* timestamp */
	created_at	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
	updated_at	TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

/*CREATE INDEX IF NOT EXISTS users_username_idx ON users (username);
CREATE INDEX IF NOT EXISTS users_email_idx ON users (email);*/