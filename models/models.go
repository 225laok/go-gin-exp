package models

import (
	"fmt"
	"github.com/go-gin-example/pkg/setting"
	"github.com/jinzhu/gorm"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn   int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func ModelInit() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	dbType = setting.DatabaseS.Type
	dbName = setting.DatabaseS.Name

	user = setting.DatabaseS.User
	password = setting.DatabaseS.Password
	host = setting.DatabaseS.Host
	tablePrefix = setting.DatabaseS.TablePrefix
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}
	db.SingularTable(true)
	db.LogMode(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallBack)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
}

func updateTimeStampForCreateCallBack(scope *gorm.Scope)  {
	if !scope.HasError(){
		nowTime:=time.Now().Unix()
		if createTimeField ,ok:=scope.FieldByName("CreatedOn");ok{
			if createTimeField.IsBlank{
				createTimeField.Set(nowTime)
			}
		}
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}

	}
}
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

func CloseDB() {
	defer db.Close()
}
