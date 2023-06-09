package controller

import (
	"context"
	"fmt"
	"time"

	"github.com/alobe/seawill/lib"
	"github.com/alobe/seawill/model"
	"github.com/alobe/seawill/util"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
	"golang.org/x/crypto/bcrypt"
)

func login(ctx *fiber.Ctx) error {
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	var searchUser model.User
	if err := lib.DB.Where("name = ?", user.Name).Or("email = ?", user.Email).First(&searchUser).Error; err != nil {
		return err
	}

	if err := bcrypt.CompareHashAndPassword([]byte(searchUser.Password), []byte(user.Password)); err != nil {
		return fiber.ErrUnauthorized
	}

	ctxbg := context.Background()

	key := util.RandomStr(64)

	_, err := lib.Rds.Set(ctxbg, key, searchUser, time.Hour*24).Result()
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "x-seawill",
		Value: key,
	})

	return ctx.Status(200).JSON(fiber.Map{
		"code":    0,
		"user_id": searchUser.ID,
	})
}

func register(ctx *fiber.Ctx) error {
	var user model.User
	if err := ctx.BodyParser(&user); err != nil {
		return err
	}

	// todo 补充校验
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPwd)

	if err := lib.DB.Model(&model.User{}).Create(&user).Error; err != nil {
		return ctx.SendStatus(400)
	}

	return ctx.JSON(fiber.Map{
		"message": "user registered successfully",
	})
}

type GetUserParams struct {
	Limit  int `json:"limit"`
	Offset int `json:"offset"`
}

type User struct {
	ID     uint   `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}

func getUserList(ctx *fiber.Ctx) error {
	var params GetUserParams
	fmt.Println(params)
	if params.Limit == 0 || params.Limit > 50 {
		params.Limit = 50
	}
	var users []User
	if err := lib.DB.Model(&model.User{}).Find(&users).Limit(params.Limit).Offset(params.Offset).Error; err != nil {
		log.Info().Msg("get user list failed")
		return err
	}
	return ctx.JSON(fiber.Map{
		"data": users,
	})
}
