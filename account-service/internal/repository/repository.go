package repository

import (
	"github.com/ilhamsyaputra/bank-api-gorm/pkg/logger"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
)

type Repository struct {
	NasabahRepository
	RekeningRepository
	TransaksiRepository
	LoginRepository
	db     *gorm.DB
	log    *logger.Logger
	tracer trace.Tracer
}

func InitRepository(db *gorm.DB, log *logger.Logger, tracer trace.Tracer) *Repository {
	nasabahRepository := InitNasabahRepositoryImpl(db, log, tracer)
	rekeningRepository := InitRekeningRepositoryImpl(db, log, tracer)
	transaksiRepository := InitTransaksiRepositoryImpl(db, log, tracer)
	loginRepository := InitLoginRepositoryImpl(db, log, tracer)

	return &Repository{
		db:     db,
		log:    log,
		tracer: tracer,

		NasabahRepository:   nasabahRepository,
		RekeningRepository:  rekeningRepository,
		TransaksiRepository: transaksiRepository,
		LoginRepository:     loginRepository,
	}
}
