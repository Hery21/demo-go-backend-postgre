# Go Backend Postgre

### Prerequisites:
- PostgreSQL Installed
- Basic SQL Knowledge
- Download or clone demo-go-backend-postgre repository (https://git.garena.com/sea-labs-id/batch-01/shared-projects/demo-go-backend-postgre)

### Table of Contents:
- Setup Database
- Create DB Connection
- Refactor

***

## A. Setup Database
1. Open pgAdmin
2. Right-click on the Database and choose Create Database
3. Name `products_db`
4. Expand the new database > Schemas > right-click Tables to create new table
5. Name `products_tab`
6. Go to Columns and create 4 columns with options:
- id bigserial NOT NULL primary key
- name character varying 50 NOT NULL
- description text NOT NULL
- quantity integer NOT NULL
  Query:
```sql
CREATE TABLE public.products_tab
(
    id bigserial NOT NULL,
    name character varying(50) NOT NULL,
    description text NOT NULL,
    quantity integer NOT NULL,
    PRIMARY KEY (id)
);

ALTER TABLE IF EXISTS public.products_tab
    OWNER to postgres;
```

7. Create dummy data by right-click on Tables > Query Tool and insert `query.sql`
8. Try to show the table value by right-click products_tab > View > All rows


## Create Database Connection
In order to do interact with the database, first we need to connect to the database.
1. Create folder `database` and file `setup.go`
2. Install `lib/pq` package (https://github.com/lib/pq)
3. Import `_ "github.com/lib/pq"`
4. Create `const` with configuration of our database
```go
// create const for database configuration
const (
	HOST    = "localhost"
	PORT    = 5432
	DB_USER = "postgres"
	DB_PASS = "123"
	DB_NAME = "products_db"
)
```
3. Create `Init` to open the database connection
```go
func InitDB() *sql.DB {
	// combine string
	// sslmode is whether we use ssl or no (HTTPS)
	dbInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", HOST, PORT, DB_USER, DB_PASS, DB_NAME)

	// open SQL connection
	DB, err := sql.Open("postgres", dbInfo)

	// check error
	if err != nil {
		fmt.Println("Error database connection")
		fmt.Println(err.Error())
	}

	return DB
}
```

## Refactor Current Project
1. Rename project:
- Edit > Find > Replace in Files
- Find previous project name, replace with new project name > Replace all
2.

Note: Choose scheme in Goland (hover in SQL query)
