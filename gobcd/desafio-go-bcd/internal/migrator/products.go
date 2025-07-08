package migrator

import (
	"app/internal"
	"app/internal/loader"
)

type MigratorProductsToDatabase struct {
	loader     loader.LoaderProduct
	repository internal.RepositoryProduct
	path       string
}

// construtor do migrator de produtos
func NewMigratorProductsToDatabase(loader loader.LoaderProduct, repository internal.RepositoryProduct, path string) *MigratorProductsToDatabase {
	return &MigratorProductsToDatabase{
		loader:     loader,
		repository: repository,
		path:       path,
	}
}

func (m *MigratorProductsToDatabase) Migrate() (err error) {
	products, err := m.loader.Load()
	if err != nil {
		return err
	}

	for _, product := range products {
		err = m.repository.Save(&product)
		if err != nil {
			return err
		}
	}
	return
}
