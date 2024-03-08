package service

import (
	"journal-service/internal/data/request"
)

type JournalService interface {
	CreateJournal(journal request.CreateJournal) error
}
