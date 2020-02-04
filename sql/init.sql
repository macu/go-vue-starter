-- Drop tables if desired to enable re-initializing database
DROP TABLE IF EXISTS session;
DROP TABLE IF EXISTS user;

-- Create minimal tables for user authentication and session management
CREATE TABLE user (
	id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
	username VARCHAR(25) UNIQUE NOT NULL,
	email VARCHAR(50) UNIQUE NOT NULL,
	auth_hash VARCHAR(60) NOT NULL,
	created_at DATETIME NOT NULL
);
CREATE TABLE session (
	token VARCHAR(30) PRIMARY KEY,
	user_id INT UNSIGNED NOT NULL,
	expires DATETIME NOT NULL,
	FOREIGN KEY (user_id) REFERENCES user(id)
);
