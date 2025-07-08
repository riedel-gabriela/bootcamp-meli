package migrator

import (
	"app/internal"
	"app/internal/loader"
)

type MigratorCustomersToDatabase struct {
	loader     loader.LoaderCustomer
	repository internal.RepositoryCustomer
	path       string
}

func NewMigratorCustomersToDatabase(loader loader.LoaderCustomer, repository internal.RepositoryCustomer, path string) *MigratorCustomersToDatabase {
	return &MigratorCustomersToDatabase{
		loader:     loader,
		repository: repository,
		path:       path,
	}
}

func (m *MigratorCustomersToDatabase) Migrate() (err error) {
	customers, err := m.loader.Load()
	if err != nil {
		return err
	}
	for _, customer := range customers {
		err = m.repository.Save(&customer)
		if err != nil {
			return err
		}
	}
	return
}
