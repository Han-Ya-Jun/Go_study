package main

/**
 *CreateBy 01305155
 *Date:10:39
 *Description:
*/
import (
	"flag"
	svc "go-restful-1/restful-1"
)
var (
	ArgLogLevel      = flag.Int("verbose", 6, "log level")
	ArgMysqlHost     = flag.String("mysqlhost", "10.118.44.147", "mysql host")
	ArgMysqlPort     = flag.String("mysqlport", "3306", "mysql port")
	ArgMysqlDB       = flag.String("mysqldb", "dubbodemo", "mysql database name")
	ArgMysqlUsername = flag.String("mysqlusername", "root", "mysql username")
	ArgMysqlPassword = flag.String("mysqlpassword", "1234", "mysql password")
	ArgServerPort    = flag.String("serverport", "8080", "server port")
)
func main() {
	flag.Parse()
	dbConf := &svc.DBConfig{
		Host:     *ArgMysqlHost,
		Port:     *ArgMysqlPort,
		Database: *ArgMysqlDB,
		Username: *ArgMysqlUsername,
		Password: *ArgMysqlPassword,
	}
	userMgr := svc.NewUserManager(dbConf)
	controller := svc.NewUserController(userMgr)
	server := svc.NewScheduleServer(*ArgServerPort, controller)
	server.Start()
}
