package migrator

import (
	"app/internal"
	"app/internal/loader"
)

type MigratorInvoicesToDatabase struct {
	loader     loader.LoaderInvoice
	repository internal.RepositoryInvoice
	path       string
}

func NewMigratorInvoicesToDatabase(loader loader.LoaderInvoice, repository internal.RepositoryInvoice, path string) *MigratorInvoicesToDatabase {
	return &MigratorInvoicesToDatabase{
		loader:     loader,
		repository: repository,
		path:       path,
	}
}

func (m *MigratorInvoicesToDatabase) Migrate() (err error) {
	invoices, err := m.loader.Load()
	if err != nil {
		return err
	}

	for _, invoice := range invoices {
		err = m.repository.Save(&invoice)
		if err != nil {
			return err
		}
	}
	return
}
