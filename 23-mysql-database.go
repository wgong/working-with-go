/**
 * mysql-database.go
 *
 * Working with mysql database
 * See: https://golang.org/pkg/database/sql/
 * See: https://github.com/go-sql-driver/mysql
 */

/*
Install mysql pkg:
$ go get github.com/go-sql-driver/mysql

You can create example table and data with the sql query below

$ mysql -u root -p
mysql> create database godb;
mysql> use godb;
mysql> show tables;

mysql> DROP TABLE IF EXISTS `posts`;

mysql> CREATE TABLE `posts` (
  `id` int(6) unsigned NOT NULL AUTO_INCREMENT,
  `title` varchar(30) NOT NULL,
  `body` text NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

mysql> show tables;
+----------------+
| Tables_in_godb |
+----------------+
| posts          |
+----------------+
1 row in set (0.00 sec)


mysql> INSERT INTO `posts` (`id`, `title`, `body`)
VALUES
	(1,'Hello World','The content of the hello world'),
	(2,'Hello Second World','The content of the hello second world'),
	(3,'Welcome to Golang world','Golang is an interesting programming lang');

mysql> select * from posts;
+----+-------------------------+-------------------------------------------+
| id | title                   | body                                      |
+----+-------------------------+-------------------------------------------+
|  1 | Hello World             | The content of the hello world            |
|  2 | Hello Second World      | The content of the hello second world     |
|  3 | Welcome to Golang world | Golang is an interesting programming lang |
+----+-------------------------+-------------------------------------------+
3 rows in set (0.00 sec)

mysql> exit

*/

// standard main package
package main

// Necessary packages to work with mysql databases
import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// dbConn connects to the database
func dbConn() (db *sql.DB) {

	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "MySQL12345"
	dbName := "godb"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// Post is the key to connect database to the program
type Post struct {
	ID    int
	Title string
	Body  string
}

// getAll gets all the records from database
func getAll() {

	db := dbConn()

	selDB, err := db.Query("SELECT * FROM posts ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	post := Post{}
	posts := []Post{}

	for selDB.Next() {

		var id int
		var title, body string

		err = selDB.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Body = body

		posts = append(posts, post)
	}

	for _, post := range posts {
		fmt.Println(post.Title)
	}

	defer db.Close()
}

// getOne gets only one record from database
func getOne(postID int) {

	db := dbConn()

	selDB, err := db.Query("SELECT * FROM Posts WHERE id=?", postID)
	if err != nil {
		panic(err.Error())
	}

	post := Post{}

	for selDB.Next() {

		var id int
		var title, body string

		err = selDB.Scan(&id, &title, &body)
		if err != nil {
			panic(err.Error())
		}

		post.ID = id
		post.Title = title
		post.Body = body

	}

	fmt.Println("Post Title	: " + post.Title)
	fmt.Println("Post Body	: " + post.Body)

	defer db.Close()
}

// add helps you to add new record to database
func add() {

	db := dbConn()

	title := "Hello Second World"
	body := "The content of the hello second world"
	insertQuery, err := db.Prepare("INSERT INTO posts(title, body) VALUES(?,?)")
	if err != nil {
		panic(err.Error())
	}

	insertQuery.Exec(title, body)

	fmt.Println("ADDED: Title: " + title + " | Body: " + body)

	defer db.Close()

}

// update helps you to update an existing record in the database
func update(postID int) {

	db := dbConn()

	title := "Hello 1 World"
	body := "The content of the hello 1 world"
	updateQuery, err := db.Prepare("UPDATE Posts SET title=?, body=? WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	updateQuery.Exec(title, body, postID)

	fmt.Println("UPDATED: Title: " + title + " | Body: " + body)

	defer db.Close()

}

// delete helps you to delete an existing record in the database
func delete(postID int) {

	db := dbConn()

	deleteQuery, err := db.Prepare("DELETE FROM Posts WHERE id=?")
	if err != nil {
		panic(err.Error())
	}

	deleteQuery.Exec(postID)

	fmt.Println("DELETED")

	defer db.Close()

}

func main() {

	//add()
	//update(1)
	//delete(1)
	//getOne(1)
	getAll()

}
