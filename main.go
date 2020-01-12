package main

import (
	"context"
	"github.com/google/uuid"
	"github.com/udayangaac/mobile-api/internal/boot"
)

func main() {
	boot.Init(context.WithValue(context.Background(), "uuid_str", uuid.New().String()))
}
