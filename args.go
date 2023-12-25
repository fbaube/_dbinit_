package dbinit

import(
	D "github.com/fbaube/dsmnd"
	RU "github.com/fbaube/repoutils"
)

type DBargs struct {
     D.DB_type // DB_SQLite = "sqlite"
     BaseFilename string // default to "mmmc.db"
     Dir string 
     DoImport bool
     DoZeroOut bool
     TableDetails []RU.TableDescriptor
}

