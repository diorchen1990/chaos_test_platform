package database

import (
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
)

var DB *gorm.DB

func Init() error {
    var err error
    DB, err = gorm.Open(mysql.Open("user:pass@tcp(127.0.0.1:3306)/chaos_platform?charset=utf8mb4"))
    return err
} 