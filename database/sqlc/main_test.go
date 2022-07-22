package database

import (
	"context"
	"database/sql"
	"log"
	"os"
	"server/config"
	"testing"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/random"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var charsets string = "abcdefghijklmnopqrstuvwxyz"

func precreateUser(username, password, email string) (User, error) {
	user, err := testQueries.CreateUser(context.Background(), CreateUserParams{
		ID:         uuid.New(),
		Username:   username,
		Password:   password,
		Email:      email,
		UserRole:   RoleSTUDENT,
		Visibility: VisibilityPUBLIC,
	})
	return user, err
}

func postDeleteUser(id uuid.UUID) (User, error) {
	deleted, err := testQueries.DeleteUser(context.Background(), id)
	return deleted, err
}

func preCreateClassroom(admin_id uuid.UUID, name string) (Classroom, error) {
	class, err := testQueries.CreateClass(context.Background(), CreateClassParams{
		ID:          uuid.New(),
		AdminID:     admin_id,
		Name:        name,
		Description: random.New().String(10, charsets),
		Room:        random.New().String(5, charsets),
		Subject:     random.New().String(8, charsets),
		Section:     random.New().String(10, charsets),
		InviteCode:  uuid.New(),
		Visibility:  VisibilityPUBLIC,
	})
	return class, err
}

func postDeleteClassroom(id uuid.UUID) (Classroom, error) {
	deleted, err := testQueries.DeleteClass(context.Background(), id)
	return deleted, err
}

func TestMain(m *testing.M) {
	godotenv.Load("./.development.env")
	cfg := config.Init()
	conn, err := sql.Open("postgres", cfg.DATABASE_URL)

	if err != nil {
		log.Fatal(err.Error())
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
