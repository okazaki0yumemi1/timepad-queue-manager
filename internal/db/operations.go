package operations

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	FirstName   string
	FamilyName  string
	LastName    string
	PhoneNumber string
}

func getAllUsers(dbpool *pgxpool.Pool) []User {
	var users []User
	rows, err := dbpool.Query(context.Background(), "SELECT * FROM users") //.Scan()(&users)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		//os.Exit(1)
	}

	defer rows.Close()

	// Iterate through the result set
	for rows.Next() {
		var user User
		err = rows.Scan(&user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		}
		users = append(users, user)
	}

	return users
}
