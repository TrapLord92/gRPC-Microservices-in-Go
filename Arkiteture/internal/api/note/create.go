package note

import (
	"context"
	"log"

	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/converter"
	"github.com/TrapLord92/gRPC-Microservices-in-Go/pkg/note_v1"
	desc "github.com/TrapLord92/gRPC-Microservices-in-Go/pkg/note_v1"
)

func (i *Implementation) Create(ctx context.Context, req *note_v1.CreateRequest) (*desc.CreateResponse, error) {
	id, err := i.noteService.Create(ctx, converter.ToNoteInfoFromDesc(req.GetInfo()))
	if err != nil {
		return nil, err
	}

	log.Printf("inserted note with id: %d", id)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}
