package dbtestutil

import (
	"fmt"
	"os"
	"github.com/olekukonko/tablewriter"
	"github.com/jmoiron/sqlx"
)

type DBOptions struct {
	Host          string
	Port          int
	Name          string
	User          string
	Password      string
	MigrationsDir string
}

type DBHelper struct {
	DB            *sqlx.DB
	migrationsDir string
}

func New(opt DBOptions) *DBHelper {
	if opt.User == "" {
		opt.User = "postgres"
	}
	if opt.Password == "" {
		opt.Password = "postgres"
	}
	if opt.Host == "" {
		opt.Host = "localhost"
	}
	if opt.Port == 0 {
		opt.Port = 5432
	}

	return &DBHelper{
		DB:            createConnection(opt),
		migrationsDir: opt.MigrationsDir,
	}
}

func (d *DBHelper) CreateSchema() {
	err := MigrateDatabaseUp(d.DB, d.migrationsDir)
	if err != nil {
		if err.Error() == "no change" {
			return
		}
		panic(err)
	}
}

func (d *DBHelper) DropSchema() {
	err := DropDatabase(d.DB, d.migrationsDir)
	if err != nil {
		panic(err)
	}
}

func (d *DBHelper) SpoilConnection() {
	err := d.DB.Close()
	if err != nil {
		panic(err)
	}
}

func (d *DBHelper) PreviewTable(tableName string) {
	rows, err := d.DB.Query(fmt.Sprintf(`SELECT * FROM %s`, tableName))
	if err != nil {
		panic(err)
	}

	table := tablewriter.NewWriter(os.Stdout)
	cols, _ := rows.Columns()
	columns := make([]interface{}, len(cols))

	for rows.Next() {
		columnPointers := make([]interface{}, len(cols))
		for i, _ := range columns {
			columnPointers[i] = &columns[i]
		}

		if err := rows.Scan(columnPointers...); err != nil {
			panic(err)
		}

		m := make(map[string]interface{})
		s := make([]string, len(cols))
		for i, colName := range cols {
			val := columnPointers[i].(*interface{})
			m[colName] = *val
			s[i] = fmt.Sprint(*val)
		}
		table.Append(s)
	}
	table.SetHeader(cols)
	table.Render()
}

func createConnection(db DBOptions) *sqlx.DB {
	conn := fmt.Sprintf(
		"host=%s port=%d dbname=%s user=%s password=%s sslmode=disable",
		db.Host, db.Port, db.Name, db.User, db.Password,
	)
	return sqlx.MustConnect("postgres", conn)
}
