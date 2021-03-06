package migrations

import migrate "github.com/rubenv/sql-migrate"

func init() {
	instance.add(&migrate.Migration{
		Id: "1568808227",
		Up: []string{
			`
            CREATE TABLE account (
			  id bigint(20) NOT NULL AUTO_INCREMENT,
			  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			  deleted_at timestamp NULL DEFAULT NULL,
			  email varchar(255) NOT NULL,
			  account_role varchar(255) NOT NULL,
			  password_hash varchar(255) NOT NULL,
			  status varchar(255),
			  reserved_books int(20) NOT NULL DEFAULT 0,
			  overdue_books int(20) NOT NULL DEFAULT 0,
			  PRIMARY KEY (id),
			  UNIQUE KEY id (id),
			  UNIQUE KEY email (email)
			);
			`,
			`
			CREATE TABLE book (
			  id bigint(20) NOT NULL AUTO_INCREMENT,
			  created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
			  updated_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
			  deleted_at timestamp NULL DEFAULT NULL,
			  name varchar(255) NOT NULL,
			  isbn varchar(255) NOT NULL,
			  stock int(20) NOT NULL,
			  author varchar(255) NOT NULL,
			  year varchar(255) NOT NULL,
			  edition int(20) NOT NULL,
			  cover text NOT NULL,
			  abstract text NOT NULL,
			  category varchar(255) NOT NULL,
			  rating int(20) NOT NULL DEFAULT 5,
			  PRIMARY KEY (id),
			  UNIQUE KEY id (id)
			);
           `,
		},
		//language=SQL
		Down: []string{
			`DROP TABLE account;`,
			`DROP TABLE book;`,
		},
	})
}
