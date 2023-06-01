package controller

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/alobe/seawill/lib"
	"github.com/alobe/seawill/model"
	"github.com/gofiber/fiber/v2"
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

	_, err := lib.Rds.Set(ctxbg, fmt.Sprintf("sealwill-user-%d", searchUser.ID), searchUser, time.Hour*24).Result()
	if err != nil {
		return err
	}

	ctx.Cookie(&fiber.Cookie{
		Name:  "x-seawill",
		Value: strconv.FormatUint(uint64(user.ID), 10),
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
		return err
	}

	return ctx.JSON(fiber.Map{
		"message": "user registered successfully",
	})
}
