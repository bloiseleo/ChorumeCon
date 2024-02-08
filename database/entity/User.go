package entity

import (
	"database/sql"

	"github.com/bloiseleo/chorumecon/database"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string
	Username string
	Password string
	Exchange float32
}

func (user *User) ComparePassword(pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(pass))
	return err == nil
}

func AdaptUser(rows *sql.Rows) *User {
	var user User
	rows.Scan(&user.Id, &user.Username, &user.Password, &user.Exchange)
	return &user
}

func AdaptUserWithoutPassword(rows *sql.Rows) *User {
	var user User
	rows.Scan(&user.Id, &user.Username, &user.Exchange)
	return &user
}

func AdaptFromToken(token *jwt.Token) *User {
	sub, err := token.Claims.GetSubject()
	if err != nil {
		panic(err)
	}

	conn := database.Connect()

	rows, err := conn.Query("SELECT id, username, exchange FROM apiusers WHERE id = " + sub)

	if err != nil {
		panic(err)
	}

	defer rows.Close()

	if !rows.Next() {
		return nil
	}

	return AdaptUserWithoutPassword(rows)
}
