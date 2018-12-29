
CREATE TABLE IF NOT EXISTS sample_users(
		id serial primary key,
		created_at timestamp NOT NULL,
		updated_at timestamp NOT NULL,
		deleted_at timestamp ,
		first_name varchar(100) NOT NULL,
		last_name varchar(100) NOT NULL,
		username varchar(100) NOT NULL,
		email varchar(254) UNIQUE NOT NULL,
		password_hash char(60) NOT NULL,
        password_updated_at timestamp NOT NULL,
        verified boolean NOT NULL
); 