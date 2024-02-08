package repository

import (
	"github.com/bloiseleo/chorumecon/database"
	"github.com/bloiseleo/chorumecon/database/entity"
)

func FindUserByDiscordId(discordId string) *entity.ChorumeUser {
	conn := database.Connect()
	stmt, err := conn.Prepare("SELECT id, discord_user_id FROM users WHERE discord_user_id = ? LIMIT 1")
	if err != nil {
		panic(err)
	}
	rows, err := stmt.Query(discordId)
	if err != nil {
		panic(err)
	}
	if !rows.Next() {
		return nil
	}
	var user entity.ChorumeUser
	err = rows.Scan(&user.Id, &user.Discord_user_id)
	if err != nil {
		panic(err)
	}
	return &user
}

func FindApiUserByName(username string) *entity.User {
	conn := database.Connect()
	stmt, err := conn.Prepare("SELECT id, username, password, exchange FROM apiusers WHERE username LIKE ? LIMIT 1")
	if err != nil {
		panic(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(username)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	if !rows.Next() {
		return nil
	}
	return entity.AdaptUser(rows)
}
