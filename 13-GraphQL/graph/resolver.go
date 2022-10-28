package graph

import "github.com/devfullcycle/13-GraphQL/internal/database"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
    CategoryDB *database.Category
    CourseDB *database.Course
}
