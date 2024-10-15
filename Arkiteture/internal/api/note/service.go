package note

import (
	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/service"
	desc "github.com/TrapLord92/gRPC-Microservices-in-Go/pkg/note_v1"
)

type Implementation struct {
	desc.UnimplementedNoteV1Server
	noteService service.NoteService
}

func NewImplementation(noteService service.NoteService) *Implementation {
	return &Implementation{
		noteService: noteService,
	}
}
