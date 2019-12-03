package models

import (
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// 获取所有 mysql 实例的函数
func GetAllMysqlClusters() ([]MysqlCluster, int) {
	var mysqls []MysqlCluster
	var count int
	db.Find(&mysqls)
	db.Model(&MysqlCluster{}).Count(&count)
	return mysqls, count
}

// 判断是否存在该 mysql 实例
func ExistMysqlCluster(db_instance_name string) bool {
	var mysql MysqlCluster
	db.Select("id").Where("cluster_name=?", db_instance_name).First(&mysql)
	if mysql.ID > 0 {
		return true
	}

	return false
}

// 创建数据库实例
func AddMysqlCluster(mysqlcluster *MysqlCluster) bool {
	db.Create(&mysqlcluster)
	return true
}

func CloseDB() {
	defer db.Close()
}
