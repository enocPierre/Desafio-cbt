package migration

import (
	"github.com/estudo-go/Desafio-cbt/model"
	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) {
	db.AutoMigrate(model.Transacao{})
	db.AutoMigrate(model.Pagamento{})
}
