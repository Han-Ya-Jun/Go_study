package service

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

/**
 *CreateBy 01305155
 *Date:10:57
 *Description::
 */
type UserManager struct {
	DBConf *DBConfig
}

func NewUserManager(dbConf *DBConfig) *UserManager {
	mgr := &UserManager{
		DBConf: dbConf,
	}
	mgr.initDB()
	return mgr
}
func (mgr *UserManager) initDB() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	ds := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", mgr.DBConf.Username, mgr.DBConf.Password, mgr.DBConf.Host, mgr.DBConf.Port, mgr.DBConf.Database)
	log.Infof("datasource=[%s]", ds)
	err := orm.RegisterDataBase("default", "mysql", ds, mgr.DBConf.MaxIdleConns, mgr.DBConf.MaxOpenConns)
	if err != nil {
		panic(err)
	}
	orm.RegisterModel(new(User))
}
func (mgr *UserManager) GetUsers() ([]*User, error) {

	o := orm.NewOrm()
	var users []*User
	_, err := o.QueryTable("user").All(&users)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (mgr *UserManager) GetUsersById(id string) (*[]User, error) {
	orm.Debug = true
	o := orm.NewOrm()
	//id2:=orm.StrTo(id)
	//key,err2:=id2.Int()
	//if err2!=nil{
	//	return nil,err2
	//}
	//var users =new(User)
	//users.Id=key
	//err:=o.Read(users)
	user := new([]User)
	_, err := o.QueryTable("user").Filter("Id", id).All(user)
	//err := o.QueryTable("user").Filter("id",key).One(user)
	//err:=o.Raw("select * from user where Id = ?",id).QueryRow(user)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return user, nil
}
