package repository

import (
	"context"

	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/repository/note/model"
)

type NoteRepository interface {
	Create(ctx context.Context, info *model.NoteInfo) (int64, error)
	Get(ctx context.Context, id int64) (*model.Note, error)
}
