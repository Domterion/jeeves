package startup

import (
	"database/sql"

	"github.com/domterion/jeeves/internal/models"
	"github.com/domterion/jeeves/internal/utils"
	"github.com/sarulabs/di/v2"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	_ "github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func InitDatabase(container di.Container) (*bun.DB, error) {
	config := container.Get(utils.DIConfig).(*models.Config)
	opened, err := sql.Open("pg", config.DatabaseUri)

	if err != nil {
		return nil, err
	}

	db := bun.NewDB(opened, pgdialect.New())

	// TODO: We need to remove this query hook when we are confident
	db.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
		bundebug.FromEnv("BUNDEBUG"),
	))

	return db, err
}
