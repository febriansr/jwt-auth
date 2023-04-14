package repository

import (
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/febriansr/jwt-auth/model"
	"github.com/febriansr/jwt-auth/utils"
	"github.com/jmoiron/sqlx"
)

type UserRepo interface {
	Login(inputedUser *model.User) any
}

type userRepo struct {
	db *sqlx.DB
}

func (u *userRepo) Login(inputedUser *model.User) any {
	var userInDb model.User
	row := u.db.QueryRow(`SELECT username, password FROM users WHERE username = $1`, inputedUser.Username)
	err := row.Scan(&userInDb.Username, &userInDb.Password)

	if userInDb.Username == "" {
		return "user data not found"
	} else if err != nil {
		log.Println(err)
		return "failed to get user data"
	} else if userInDb.Username == inputedUser.Username && userInDb.Password == inputedUser.Password {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		claims["username"] = userInDb.Username
		expiry, _ := strconv.Atoi(utils.DotEnv("EXP"))
		claims["exp"] = time.Now().Add(time.Minute * time.Duration(expiry)).Unix()
		jwtKey := []byte(utils.DotEnv("JWTKEY"))

		tokenString, err := token.SignedString(jwtKey)

		if err != nil {
			return "failed to generate token"
		}

		return tokenString
	} else {
		return "invalid credentials"
	}
}

func NewUserRepo(db *sqlx.DB) UserRepo {
	repo := new(userRepo)
	repo.db = db
	return repo
}
