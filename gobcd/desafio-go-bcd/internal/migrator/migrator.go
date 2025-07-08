package migrator

type Migrator interface {
	Migrate() error
}
