import (
	"database/sql"
	_ "github.com/lib/pq"
)

func dbExample() {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	var name string
	err = db.QueryRow("SELECT name FROM users WHERE id = $1", 1).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
