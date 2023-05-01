package PTGUdatabase

import "os"

func GenerateTableName(name string) string {
	if os.Getenv("DATABASE_TABLE_PREFIX") == "" {
		return "df_" + name
	}
	return os.Getenv("DATABASE_TABLE_PREFIX") + name
}