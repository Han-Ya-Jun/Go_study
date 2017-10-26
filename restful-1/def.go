package service

/**
 *CreateBy 01305155
 *Date:10:58
 *Description::
 */
type DBConfig struct {
	Host         string
	Port         string
	Database     string
	Username     string
	Password     string
	MaxIdleConns int
	MaxOpenConns int
}
type User struct {
	Id       int    `json:"id" orm:"column(id),size(11)"`
	Username string `json:"username" orm:"column(username)"`
	Password string `json:"password" orm:"column(password)"`
	Sex      string `json:"sex" orm:"column(sex)"`
	Tel      string `json:"tel" orm:"column(tel)"`
}
type RespWrapper struct {
	RequestId    string      `json:"requestId"`
	Success      bool        `json:"success"`
	Business     string      `json:"business"`
	ErrorCode    string      `json:"errorCode"`
	ErrorMessage string      `json:"errorMessage"`
	Date         string      `json:"date"`
	Version      string      `json:"version"`
	Obj          interface{} `json:"obj"`
}
