package services

import (
	"errors"

	"maspulung/gotodo/app/entities"
	"maspulung/gotodo/app/repositories"
	"maspulung/gotodo/app/types"
	"maspulung/gotodo/config/database"
	"maspulung/gotodo/utils"
	"maspulung/gotodo/utils/jwt"
	"maspulung/gotodo/utils/password"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var userRepository = repositories.NewUserRepository(database.DB)

// Login service logs in a user
func Login(ctx *fiber.Ctx) error {
	body := new(types.LoginDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	userResponse := &types.UserResponse{}

	err := userRepository.FindUserByEmail(userResponse, body.Email).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	if err := password.Verify(userResponse.Password, body.Password); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid email or password")
	}

	token := jwt.Generate(&jwt.TokenPayload{
		ID: userResponse.ID,
	})

	return ctx.JSON(&types.AuthResponse{
		User: userResponse,
		Auth: &types.AccessResponse{
			Token: token,
		},
	})
}

// Signup service creates a user
func Signup(ctx *fiber.Ctx) error {
	body := new(types.SignupDTO)

	if err := utils.ParseBodyAndValidate(ctx, body); err != nil {
		return err
	}

	userResponse := &types.UserResponse{}

	err := userRepository.FindUserByEmail(&userResponse, body.Email).Error

	// If email already exists, return
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return fiber.NewError(fiber.StatusConflict, "Email already exists")
	}

	user := &entities.User{
		Name:     body.Name,
		Password: password.Generate(body.Password),
		Email:    body.Email,
	}

	// Create a user, if error return
	if err := userRepository.CreateUser(user); err.Error != nil {
		return fiber.NewError(fiber.StatusConflict, err.Error.Error())
	}

	// generate access token
	token := jwt.Generate(&jwt.TokenPayload{
		ID: user.ID,
	})

	return ctx.JSON(&types.AuthResponse{
		User: &types.UserResponse{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
		},
		Auth: &types.AccessResponse{
			Token: token,
		},
	})
}
