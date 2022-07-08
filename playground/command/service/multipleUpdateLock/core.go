package multipleUpdateLock

import (
	"go-leetcode/pkg/mysql"
	"log"
)

func Launcher() {

	db := mysql.GetDB()
	log.Print(db)
}
