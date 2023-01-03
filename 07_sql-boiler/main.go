package main

import (
	"context"
	"database/sql"
	"log"
	dbmodels "sqlboiler-test/db/models"

	_ "github.com/lib/pq"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func main() {
	ctx := context.Background()
	db := connectDB()

	boil.SetDB(db)

	author := createAuthor(ctx)
	createArticle(ctx, author)
	createArticle(ctx, author)
}

func createArticle(ctx context.Context, author dbmodels.Author) dbmodels.Article {
	article := dbmodels.Article{
		Title:    "Hello World",
		Body:     null.StringFrom("Hello world, this is an article."),
		AuthorID: author.ID,
	}

	err := article.InsertG(ctx, boil.Infer())
	if err != nil {
		log.Fatal(err)
	}

	return article
}

// 👇 function that creates a new author
func createAuthor(ctx context.Context) dbmodels.Author {
	author := dbmodels.Author{
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	err := author.InsertG(ctx, boil.Infer()) // 👈 insert author
	if err != nil {
		log.Fatal(err)
	}

	return author
}

func connectDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost:2345/postgres?sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
