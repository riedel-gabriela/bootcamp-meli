package migrator

import (
	"app/internal"
	"app/internal/loader"
)

type MigratorSalesToDatabase struct {
	loader     loader.LoaderSale
	repository internal.RepositorySale
	path       string
}

// construtor do migrator de vendas
func NewMigratorSalesToDatabase(loader loader.LoaderSale, repository internal.RepositorySale, path string) *MigratorSalesToDatabase {
	return &MigratorSalesToDatabase{
		loader:     loader,
		repository: repository,
		path:       path,
	}
}

func (m *MigratorSalesToDatabase) Migrate() (err error) {
	sales, err := m.loader.Load()
	if err != nil {
		return err
	}

	for _, sale := range sales {
		err = m.repository.Save(&sale)
		if err != nil {
			return err
		}
	}
	return
}
