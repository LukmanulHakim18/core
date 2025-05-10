# Package Name: database

## Usage

Below are examples of how to use the functions available in the <b>database</b> package:

1. Function NewQBClient

   ```go
   package main

   import (
       "context"
       "fmt"
       "time"
       "log"

       "github.com/LukmanulHakim18/core/database"
       commonMicrosvc "github.com/LukmanulHakim18/core/microservice"
   )

   type databaseConf struct {
       Host         string
       Username     string
       Password     string
       SslMode      string
       Port         string
       DatabaseName string
   }

   type user struct {
       Id              int       `db:"id"`
       Username        string    `db:"username"`
       Email           string    `db:"email"`
       Password        string    `db:"password"`
   }

   func main() {
       dbConf := getDbConf()
       connectionStr := connectionString(dbConf)

       qbConn, err := database.NewQBClient(
           dbConf.Host,
           dbConf.Username,
           dbConf.Password,
           dbConf.SslMode,
           dbConf.Port,
           dbConf.DatabaseName)

       if err != nil {
           log.Fatalf(err.Error())
       }

       res := user{}
       err = qbConn.Select("users").
           Where("user_id", "=", 97).FindOne(context.Background(), &res)
       if err != nil {
           log.Fatalf(err.Error())
       }

       fmt.Printf("%+v\n", res)
   }

   func getDbConf() *databaseConf {
       return &databaseConf{
           Host:         commonMicrosvc.OsGetString("DB_HOST", ""),
           Username:     commonMicrosvc.OsGetString("DB_USERNAME", ""),
           Password:     commonMicrosvc.OsGetString("DB_PASSWORD", ""),
           SslMode:      commonMicrosvc.OsGetString("DB_SSL_MODE", ""),
           Port:         commonMicrosvc.OsGetString("DB_PORT", ""),
           DatabaseName: commonMicrosvc.OsGetString("DB_NAME", ""),
       }
   }

   func connectionString(dc *databaseConf) string {
       connectionStringTemplate := "host=%s port=%s sslmode=%s user=%s password='%s' dbname=%s "
       return fmt.Sprintf(connectionStringTemplate, dc.Host, dc.Port, dc.sslMode, dc.Username, dc.Password, dc.DatabaseName)
   }


   ```
