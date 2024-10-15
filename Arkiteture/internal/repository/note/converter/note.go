package converter

import (
	"github.com/TrapLord92/gRPC-Microservices-in-Go/internal/repository/note/model"
	modelRepo "github.com/TrapLord92/gRPC-Microservices-in-Go/internal/repository/note/model"
)

func ToNoteFromRepo(note *modelRepo.Note) *model.Note {
	return &model.Note{
		ID:        note.ID,
		Info:      ToNoteInfoFromRepo(note.Info),
		CreatedAt: note.CreatedAt,
		UpdatedAt: note.UpdatedAt,
	}
}

func ToNoteInfoFromRepo(info modelRepo.NoteInfo) model.NoteInfo {
	return model.NoteInfo{
		Title:   info.Title,
		Content: info.Content,
	}
}
