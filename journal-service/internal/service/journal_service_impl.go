package service

import (
	"journal-service/internal/data/request"
	"journal-service/internal/entity"
	"journal-service/internal/repository"
	Errors "journal-service/pkg/errors"
	"journal-service/pkg/helper"
	"journal-service/pkg/logger"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type JournalServiceImpl struct {
	journalRepository repository.JournalRepository
	validate          *validator.Validate
	log               *logger.Logger
	db                *gorm.DB
}

func InitJournalRepositoryImpl(db *gorm.DB, repo repository.JournalRepository, validator *validator.Validate, logger *logger.Logger) JournalService {
	return &JournalServiceImpl{
		journalRepository: repo,
		validate:          validator,
		log:               logger,
		db:                db,
	}
}

func (service *JournalServiceImpl) CreateJournal(journal request.CreateJournal) (err error) {
	service.log.Info(logrus.Fields{}, journal, "CREATEJOURNAL START")

	err = service.validate.Struct(journal)
	if err != nil {
		helper.ServiceError(&err, service.log)
		_, errMsg := helper.RequestValidation(err)
		err = Errors.NewError(errMsg)
		helper.ServiceError(&err, service.log)
		return
	}

	tx := service.db.Begin()

	defer helper.TransactionStatusHandler(tx, &err, service.log)

	tanggalTransaksi, err := time.Parse("02-01-2006", journal.TanggalTransaksi)

	journal_ := entity.Journal{
		Id:               uuid.New(),
		NoRekeningKredit: journal.NoRekeningKredit,
		NoRekeningDebit:  journal.NoRekeningDebit,
		NominalKredit:    journal.NominalKredit,
		NominalDebit:     journal.NominalDebit,
		TanggalTransaksi: tanggalTransaksi,
	}

	err = service.journalRepository.CreateJournal(tx, journal_)
	if err != nil {
		helper.ServiceError(&err, service.log)
		return
	}

	defer service.log.Info(logrus.Fields{}, journal, "CREATEJOURNAL END")

	return
}
