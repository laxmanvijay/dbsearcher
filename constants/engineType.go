package constants

// EngineType refers the type of db
type EngineType string

const (
	// MySQL refers mysql engine type
	MySQL EngineType = "mysql"

	// PostgreSQL refers postgres engine type
	PostgreSQL EngineType = "postgres"
)

// String is defined to implement flag.Value
func (engine *EngineType) String() string {
	return string(*engine)
}

// Set is defined to implement flag.Value
func (engine *EngineType) Set(val string) error {
	if val == "mysql" {
		*engine = MySQL
	} else {
		*engine = PostgreSQL
	}
	return nil
}
