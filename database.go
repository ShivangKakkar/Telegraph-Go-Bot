package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var ctx = context.Background()
var _ = godotenv.Load(".env")
var dsn = os.Getenv("DATABASE_URL")
var sqldb = sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
var db = bun.NewDB(sqldb, pgdialect.New())

// Um.. Bad Practice?
var _, _ = db.NewCreateTable().Model((*User)(nil)).Exec(ctx)

// Reset on new features [Not in Production]
// var _ = db.ResetModel(ctx, (*User)(nil))

func GetAllUsers() []User {
	users := make([]User, 0)
	db.NewSelect().
		Model(&users).
		ColumnExpr("id").
		Scan(context.Background())
	return users
}

func AddUser(id int64) error {
	user := User{ID: id}
	_, err := db.NewInsert().
		Model(&user).
		Exec(ctx)
	return err
}

func DeleteUser(id int64) error {
	user := User{ID: id}
	_, err := db.NewDelete().
		Model(&user).
		Where("id = ?", id).
		Exec(ctx)
	return err
}

func GetUser(id int64) *User {
	var token string
	var tokens []string
	// CheckPoint : End
	err := db.NewSelect().
		Model((*User)(nil)).
		Column("tokens", "account").
		Where("id = ?", id).
		Scan(ctx, pgdialect.Array(&tokens), &token)
	if err != nil {
		fmt.Printf("Error: '%v', while calling GetUser", err)
	}
	u := User{ID: id, Account: token, Tokens: tokens}
	return &u
}

func AddAccount(id int64, token string) {
	user := GetUser(id)
	ts := user.Tokens
	u := User{ID: id, Tokens: append(ts, token)}
	db.NewUpdate().
		Model(&u).
		Set("tokens = ?", pgdialect.Array(append(ts, token))).
		Where("id = ?", id).
		Exec(ctx)
}

func SetDefaultAccount(id int64, token string) {
	db.NewUpdate().
		// CheckPoint : End - 1
		Model((*User)(nil)).
		Set("account = ?", token).
		Where("id = ?", id).
		Exec(ctx)
}

// func GetAccount(id int, token string) {
//
// }

func GetAllAccounts(id int64) []string {
	user := GetUser(id)
	tokens := []string{user.Account}
	tokens = append(tokens, user.Tokens...)
	return tokens
}

func RemoveAccount(id int64, token string) {
	user := GetUser(id)
	ts := user.Tokens
	var newTs []string
	for _, t := range ts {
		if t != token {
			newTs = append(newTs, t)
		}
	}
	u := User{ID: id, Tokens: newTs}
	db.NewUpdate().
		Model(&u).
		Set("tokens = ?", pgdialect.Array(newTs)).
		Where("id = ?", id).
		Exec(ctx)
}

func UsersCount() int {
	return len(GetAllUsers())
}

// ToDo: Rename Column ID (id) to UserID (user_id)
type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`
	ID            int64    `bun:",pk"`
	Tokens        []string `bun:",array"`
	Account       string
}
