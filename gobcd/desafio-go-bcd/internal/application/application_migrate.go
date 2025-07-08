package application

import (
	"app/internal/loader"
	"app/internal/migrator"
	"app/internal/repository"
	"database/sql"

	"github.com/go-sql-driver/mysql"
)

// ConfigApplicationMigrate is the struct that contains the paths to the files that will be loaded
type ConfigApplicationMigrate struct {
	Db               *mysql.Config
	FilePathCustomer string
	FilePathProduct  string
	FilePathInvoice  string
	FilePathSale     string
}

// NewApplicationMigrate returns a new ApplicationMigrate
func NewApplicationMigrate(config *ConfigApplicationMigrate) (a *ApplicationMigrate) {
	a = &ApplicationMigrate{
		config: config,
	}
	return
}

// ApplicationMigrate is the implementation of the interface ApplicationMigrate
type ApplicationMigrate struct {
	// config is the configuration of the application
	config *ConfigApplicationMigrate
	// database is the database to load the data
	database *sql.DB
	// Migrators
	migrators []migrator.Migrator
}

// SetUp is the method to set up the application migrate
func (a *ApplicationMigrate) SetUp() (err error) {
	// dependencies
	// - db: init
	a.database, err = sql.Open("mysql", a.config.Db.FormatDSN())
	if err != nil {
		return
	}
	// - db: ping
	err = a.database.Ping()
	if err != nil {
		return
	}
	// - migrators
	rpCustomer := repository.NewCustomersMySQL(a.database)
	ldCustomer := loader.NewCustomersJSON(&a.config.FilePathCustomer)
	mgCustomer := migrator.NewMigratorCustomersToDatabase(ldCustomer, rpCustomer, a.config.FilePathCustomer)

	rpProduct := repository.NewProductsMySQL(a.database)
	ldProduct := loader.NewProductsJSON(&a.config.FilePathProduct)
	mgProduct := migrator.NewMigratorProductsToDatabase(ldProduct, rpProduct, a.config.FilePathProduct)

	rpInvoice := repository.NewInvoicesMySQL(a.database)
	ldInvoice := loader.NewInvoicesJSON(&a.config.FilePathInvoice)
	mgInvoice := migrator.NewMigratorInvoicesToDatabase(ldInvoice, rpInvoice, a.config.FilePathInvoice)

	rpSale := repository.NewSalesMySQL(a.database)
	ldSale := loader.NewSalesJSON(&a.config.FilePathSale)
	mgSale := migrator.NewMigratorSalesToDatabase(ldSale, rpSale, a.config.FilePathSale)

	a.migrators = []migrator.Migrator{
		mgCustomer,
		mgInvoice,
		mgProduct,
		mgSale,
	}

	return
}

// Run is the method to run the application migrate
func (a *ApplicationMigrate) Run() (err error) {
	for _, v := range a.migrators {
		err = v.Migrate()
		if err != nil {
			return
		}
	}

	return
}
