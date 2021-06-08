package models

import (
	"gorm.io/gorm"
)

type Info struct {
	gorm.Model
	ID             int
	founder        string
	founded        string
	employees      string
	vehicles       string
	launch_sites   string
	test_sites     string
	ceo            string
	cto            string
	coo            string
	cto_propulsion string
	valuation      string
	summary        string
}

//create a user
func CreateInfo(db *gorm.DB, Info *Info) (err error) {
	err = db.Create(Info).Error
	if err != nil {
		return err
	}
	return nil
}

//get users
func GetInfos(db *gorm.DB, Info *[]Info) (err error) {
	err = db.Find(Info).Error
	if err != nil {
		return err
	}
	return nil
}

//get user by id
func GetInfo(db *gorm.DB, Info *Info, id string) (err error) {
	err = db.Where("id = ?", id).First(Info).Error
	if err != nil {
		return err
	}
	return nil
}

//update user
func UpdateInfo(db *gorm.DB, Info *Info) (err error) {
	db.Save(Info)
	return nil
}

//delete user
func DeleteInfo(db *gorm.DB, Info *Info, id string) (err error) {
	db.Where("id = ?", id).Delete(Info)
	return nil
}
