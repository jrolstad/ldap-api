package repositories

import (
	"database/sql"
	"fmt"
	"github.com/jrolstad/ldap-api/internal/pkg/models"
	"log"
	"os"
	"strings"
)

type DatabaseDirectoryObjectRepository struct {
	connection *sql.DB
}

func (r *DatabaseDirectoryObjectRepository) init() {

	host := os.Getenv("directoryobject_db_host")
	port := os.Getenv("directoryobject_db_port")
	user := os.Getenv("directoryobject_db_username")
	password := os.Getenv("directoryobject_db_userpassword")
	dbName := os.Getenv("directoryobject_db_database")
	enableSSL := true

	var err error
	r.connection, err = r.buildConnection(host, port, user, password, dbName, enableSSL)
	if err != nil {
		log.Fatal("Unable to open a new database connection")
	}
}

func (r *DatabaseDirectoryObjectRepository) Destroy() {
	if r.connection != nil {
		r.connection.Close()
	}
}

func (r *DatabaseDirectoryObjectRepository) buildConnection(host string,
	port string,
	user string,
	password string,
	databaseName string,
	enableSSL bool) (*sql.DB, error) {

	connectionString := r.buildConnectionString(host, port, user, password, databaseName, enableSSL)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return db, nil
}

func (r *DatabaseDirectoryObjectRepository) buildConnectionString(host string,
	port string,
	user string,
	password string,
	databaseName string,
	enableSSL bool) string {
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s",
		host, port, user, password, databaseName)

	if !enableSSL {
		connectionString += " sslmode=disable"
	}

	return connectionString
}

func (r *DatabaseDirectoryObjectRepository) Save(item *models.DirectoryObject) error {
	if strings.EqualFold("user", item.ObjectType) {
		return r.saveUser(item)
	}

	return nil
}

func (r *DatabaseDirectoryObjectRepository) saveUser(item *models.DirectoryObject) error {
	command := "INSERT INTO user (id,object_type,location,upn,name,email,given_name,surname,manager,type,company,department,status,title) values (1,2,3,4,5,6,7,8,9,10,11,12,13,14)"
	values := []string{""}
	_, err := r.connection.Exec(command, values)
	return err
}
