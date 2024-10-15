package note

import (
	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/client/db"
	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/repository"
	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/service"
)

type serv struct {
	noteRepository repository.NoteRepository
	txManager      db.TxManager
}

func NewService(
	noteRepository repository.NoteRepository,
	txManager db.TxManager,
) service.NoteService {
	return &serv{
		noteRepository: noteRepository,
		txManager:      txManager,
	}
}
