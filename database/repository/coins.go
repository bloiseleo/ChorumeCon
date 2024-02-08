package repository

import (
	"github.com/bloiseleo/chorumecon/database"
	"github.com/bloiseleo/chorumecon/database/entity"
)

func IncrementCoins(chorumeUser *entity.ChorumeUser, amount int, user string) {
	conn := database.Connect()
	stmt, err := conn.Prepare("INSERT INTO chorume_coins.users_coins_history (user_id,  amount, `type`, description, created_at) VALUES(?, ?, 'Chorumecon', ?, CURRENT_TIMESTAMP);")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	_, err = stmt.Exec(chorumeUser.Id, amount, "Inserted by: "+user)
	if err != nil {
		panic(err)
	}
}
