package querytesting

import (
	"context"
	"database/sql"
	"log"
	"testing"
	"time"

	db "github.com/aniket0951/video_status/sqlc_lib"
	"github.com/aniket0951/video_status/utils"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func hasAndSalt(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if err != nil {
		log.Println(err)
		panic("Failed to hash a password")
	}
	return string(hash)
}
func comparePassword(hashedPwd string, plainPassword []byte) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, plainPassword)

	if err != nil {
		log.Println(err)
		return false
	}
	return true
}

func createRandomAdmin(t *testing.T) db.Users {
	args := db.CreateAdminUserParams{
		Name:     utils.RandomString(6),
		Email:    utils.RandomEmail(),
		Contact:  utils.RandomMobileNumber(),
		Password: hasAndSalt([]byte("user@123")),
		UserType: "Admin",
		IsAccountActive: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	user, err := testQueries.CreateAdminUser(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Name, args.Name)
	require.Equal(t, user.Email, args.Email)
	require.Equal(t, user.Password, args.Password)
	require.Equal(t, user.Contact, args.Contact)
	require.Equal(t, user.UserType, args.UserType)
	require.Equal(t, user.IsAccountActive, args.IsAccountActive)

	require.True(t, comparePassword(user.Password, []byte("user@123")))
	return user
}

func TestCreateAdminUser(t *testing.T) {
	createRandomAdmin(t)
}

func TestGetUser(t *testing.T) {
	first_user := createRandomAdmin(t)

	sec_user, err := testQueries.GetUser(context.Background(), first_user.ID)

	require.NoError(t, err)
	require.NotEmpty(t, sec_user)

	require.Equal(t, sec_user.Name, first_user.Name)
	require.Equal(t, sec_user.Email, first_user.Email)
	require.Equal(t, sec_user.Password, first_user.Password)
	require.Equal(t, sec_user.Contact, first_user.Contact)
	require.Equal(t, sec_user.UserType, first_user.UserType)
	require.Equal(t, sec_user.IsAccountActive, first_user.IsAccountActive)

	require.WithinDuration(t, sec_user.CreatedAt, first_user.CreatedAt, time.Second)
}

func TestGetUsers(t *testing.T) {
	args := db.GetUsersParams{
		Limit:  1,
		Offset: 2,
	}

	users, err := testQueries.GetUsers(context.Background(), args)

	require.NoError(t, err)
	require.NotEmpty(t, users)
}

func TestDeleteUser(t *testing.T) {

	user_id, err := uuid.Parse("9478f0c6-ef00-4426-823e-c3e4935ba638")

	require.NoError(t, err)

	result, err := testQueries.DeleteUser(context.Background(), user_id)
	require.NoError(t, err)
	rows_effect, _ := result.RowsAffected()
	require.NotZero(t, rows_effect)
}

func TestUserLogin(t *testing.T) {
	const password = "user@123"
	const email = "PfpxtkPoekAlpM@gmail.com"

	// fetch user by mail first
	user, err := testQueries.GetUserByEmail(context.Background(), email)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, user.Email, email)

	// match the user password
	is_password := comparePassword(user.Password, []byte(password))
	require.True(t, is_password)
}
