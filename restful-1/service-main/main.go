package main

/**
 *CreateBy 01305155
 *Date:10:39
 *Description:
 */
import (
	svc "Go_study/restful-1"
	"flag"
)

var (
	ArgLogLevel      = flag.Int("verbose", 6, "log level")
	ArgMysqlHost     = flag.String("mysqlhost", "localhost", "mysql host")
	ArgMysqlPort     = flag.String("mysqlport", "3306", "mysql port")
	ArgMysqlDB       = flag.String("mysqldb", "user", "mysql database name")
	ArgMysqlUsername = flag.String("mysqlusername", "root", "mysql username")
	ArgMysqlPassword = flag.String("mysqlpassword", "123", "mysql password")
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
