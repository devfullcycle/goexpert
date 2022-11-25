package main

import (
	"context"
	"database/sql"

	"github.com/devfullcycle/17-SQLC/internal/db"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    ctx := context.Background()
    dbConn, err := sql.Open("mysql","root:root@tcp(localhost:3306)/courses")
    if err != nil {
        panic(err)
    }
    defer dbConn.Close()

    queries := db.New(dbConn)

    // err = queries.CreateCategory(ctx, db.CreateCategoryParams{
    //     ID: uuid.New().String(),
    //     Name: "Backend",
    //     Description: sql.NullString{String: "Backend description", Valid: true},
    // })

    // if err != nil {
    //     panic(err)
    // }

    // categories, err := queries.ListCategories(ctx)
    // if err != nil {
    //     panic(err)
    // }

    // for _, category := range categories {
    //     println(category.ID, category.Name, category.Description.String)
    // }

    // err = queries.UpdateCategory(ctx, db.UpdateCategoryParams{
    //     ID: "361d935e-4ed1-49c3-afad-bdf5da95bbc9",
    //     Name: "Backend updated",
    //     Description: sql.NullString{String: "Backend description updated", Valid: true},
    // })

    

    err = queries.DeleteCategory(ctx, "361d935e-4ed1-49c3-afad-bdf5da95bbc9")
    if err != nil {
        panic(err)
    }

    categories, err := queries.ListCategories(ctx)
    if err != nil {
        panic(err)
    }

    for _, category := range categories {
        println(category.ID, category.Name, category.Description.String)
    }
}