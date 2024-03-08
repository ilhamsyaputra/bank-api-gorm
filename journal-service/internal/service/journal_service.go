package service

import (
	"journal-service/internal/data/request"
)

type JournalService interface {
	CreateJournal(nasabah request.CreateJournal) error
}
