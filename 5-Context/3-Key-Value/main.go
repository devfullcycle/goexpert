package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotel(ctx, "Hotel")
}

func bookHotel(ctx context.Context, name string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
