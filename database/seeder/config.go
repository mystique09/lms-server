package seeder

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	database "server/database/sqlc"
	"server/utils"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

type Seeder struct {
	Table  string
	Amount int
	Conn   *database.Queries
}

func New(table string, amount int, cfg *utils.Config) *Seeder {
	conn, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("Cannot connect to DB: ", err)
	}

	queries := database.New(conn)

	return &Seeder{
		Table:  table,
		Amount: amount,
		Conn:   queries,
	}
}

func (s *Seeder) Run() {
	switch s.Table {
	case "users":
		seedUser(s.Amount, s.Conn)
		return
	case "classrooms":
		seedClassroom(s.Amount, s.Conn)
		return
	case "classworks":
		seedClasswork(s.Amount, s.Conn)
		return
	default:
		log.Fatal("Table not found")
		return
	}
}

func seedUser(amount int, q *database.Queries) {
	ch := make(chan string)

	log.Printf("Seeding users table with %d random data", amount)

	for i := 0; i < amount; i++ {
		go seedUserRunner(q, ch)
	}

	for status := range ch {
		log.Println(status)
	}

	close(ch)
	os.Exit(0)
}

func seedUserRunner(q *database.Queries, r chan string) {
	encryptedPassword, err := utils.Encrypt(utils.RandomString(12))
	if err != nil {
		r <- err.Error()
		return
	}

	randomUser := database.CreateUserParams{
		ID:         uuid.New(),
		Username:   utils.RandomString(12),
		Password:   encryptedPassword,
		Email:      utils.RandomEmail(),
		UserRole:   database.RoleSTUDENT,
		Visibility: database.VisibilityPUBLIC,
	}

	userId, err := q.CreateUser(context.TODO(), randomUser)
	if err != nil {
		r <- err.Error()
		return
	}

	r <- userId
}

func seedClassroom(amount int, q *database.Queries) {
	ch := make(chan string)

	log.Printf("Seeding classrooms table with %d random data", amount)

	for i := 0; i < amount; i++ {
		go seedClassroomRunner(q, ch, amount)
	}

	for status := range ch {
		log.Println(status)
	}
}

func seedClassroomRunner(q *database.Queries, r chan string, amount int) {
	users, err := q.GetUsers(context.TODO(), 0)

	if err != nil {
		r <- err.Error()
		return
	}

	for _, user := range users {
		r <- fmt.Sprintf("Seeding user %s with %d classrooms", user.ID.String(), amount)
		for i := 0; i < amount; i++ {
			randomClassroom := database.CreateClassParams{
				ID:          uuid.New(),
				AdminID:     user.ID,
				Name:        utils.RandomString(12),
				Description: utils.RandomString(24),
				Subject:     utils.RandomString(8),
				Section:     utils.RandomString(12),
				Room:        utils.RandomString(12),
				InviteCode:  uuid.New(),
				Visibility:  database.VisibilityPUBLIC,
			}

			room, err := q.CreateClass(context.TODO(), randomClassroom)
			if err != nil {
				r <- err.Error()
				return
			}
			r <- fmt.Sprintf("Room %s has been added to user %s", room.ID.String(), user.ID.String())
		}
		r <- fmt.Sprintf("Done seeding user %s with %d classrooms", user.ID.String(), amount)
	}
}

func seedClasswork(amount int, q *database.Queries) {
	for i := 0; i < amount; i++ {

	}
}

func seedClassworkRunner(q *database.Queries, r chan string) {}
