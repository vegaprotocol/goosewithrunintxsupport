package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"text/template"
)

type TemplateData struct {
	Version   int64
	DBDriver  string
	DBOpen    string
	Direction bool
	Func      string
}

//
// Run a .go migration.
//
// In order to do this, we copy a modified version of the
// original .go migration, and execute it via `go run` along
// with a main() of our own creation.
//
func runGoMigration(conf *DBConf, path string, version int64, direction bool) error {

	// everything gets written to a temp dir, and zapped afterwards
	d, e := ioutil.TempDir("", "goose")
	if e != nil {
		log.Fatal(e)
	}
	defer os.RemoveAll(d)

	directionStr := "Down"
	if direction {
		directionStr = "Up"
	}

	td := &TemplateData{
		Version:   version,
		DBDriver:  conf.Driver,
		DBOpen:    conf.OpenStr,
		Direction: direction,
		Func:      fmt.Sprintf("%v_%v", directionStr, version),
	}
	main, e := writeTemplateToFile(filepath.Join(d, "goose_main.go"), goMigrationTmpl, td)
	if e != nil {
		log.Fatal(e)
	}

	outpath := filepath.Join(d, filepath.Base(path))
	if _, e = copyFile(outpath, path); e != nil {
		log.Fatal(e)
	}

	cmd := exec.Command("go", "run", main, outpath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if e = cmd.Run(); e != nil {
		log.Fatal("`go run` failed: ", e)
	}

	return nil
}

//
// template for the main entry point to a go-based migration.
// this gets linked against the substituted versions of the user-supplied
// scripts in order to execute a migration via `go run`
//
var goMigrationTmpl = template.Must(template.New("driver").Parse(`
package main

import (
	"database/sql"
	_ "github.com/bmizerany/pq"
	"log"
	"fmt"
)

func main() {
	db, err := sql.Open("{{.DBDriver}}", "{{.DBOpen}}")
	if err != nil {
		log.Fatal("failed to open DB:", err)
	}
	defer db.Close()

	txn, err := db.Begin()
	if err != nil {
		log.Fatal("db.Begin:", err)
	}

	{{ .Func }}(txn)

	// XXX: drop goose_db_version table on some minimum version number?
	versionFmt := "INSERT INTO goose_db_version (version_id, is_applied) VALUES (%v, %t);"
	versionStmt := fmt.Sprintf(versionFmt, int64({{ .Version }}), {{ .Direction }})
	if _, err = txn.Exec(versionStmt); err != nil {
		txn.Rollback()
		log.Fatal("failed to write version: ", err)
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal("Commit() failed:", err)
	}
}
`))