package dao

import (
	"INServer/src/common/dbobj"
	"INServer/src/common/logger"
	"INServer/src/proto/db"
)

func AllRoleSummaryDataQuery(DB *dbobj.DBObject) []*db.DBRole {
	rows, err := DB.DB().Query("select UUID, SummaryData from roles")
	if err != nil {
		logger.Fatal(err)
	}
	roles := []*db.DBRole{}
	for rows.Next() {
		role := &db.DBRole{}
		rows.Scan(&role.UUID, &role.SummaryData)
	}
	return roles
}

func RoleOnlineDataQuery(DB *dbobj.DBObject, uuid string) ([]byte, error) {
	onlineData := make([]byte, 0)
	row := DB.DB().QueryRow("select OnlineData from roles where UUID=?", uuid)
	if err := row.Scan(&onlineData); err != nil {
		logger.Debug(err)
		return nil, err
	}
	return onlineData, nil
}

func RoleInsert(DB *dbobj.DBObject, role *db.DBRole) error {
	_, err := DB.DB().Exec("insert INTO roles(UUID,SummaryData,OnlineData) values(?,?,?)", role.UUID, role.SummaryData, role.OnlineData)
	if err != nil {
		logger.Debug(err)
		return err
	}
	return nil
}

func RoleUpdate(DB *dbobj.DBObject, role *db.DBRole) error {
	_, err := DB.DB().Exec("UPDATE roles set SummaryData=? OnlineData=? where UUID=?", role.SummaryData, role.OnlineData, role.UUID)
	if err != nil {
		logger.Debug(err)
		return err
	}
	return nil
}
