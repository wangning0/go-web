

package data

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"crypto/rand"
	"fmt"
	"crypto/sha1"
)

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:wangning@tcp(127.0.0.1:3306)/chitchat?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// copy
func createUUID() (uuid string) {
	u := new([16]byte)
	_, err := rand.Read(u[:])
	if err != nil {
		log.Fatalln("Cannot generate UUID", err)
	}

	// 0x40 is reserved variant from RFC 4122
	u[8] = (u[8] | 0x40) & 0x7F
	// Set the four most significant bits (bits 12 through 15) of the
	// time_hi_and_version field to the 4-bit version number.
	u[6] = (u[6] & 0xF) | (0x4 << 4)
	uuid = fmt.Sprintf("%x-%x-%x-%x-%x", u[0:4], u[4:6], u[6:8], u[8:10], u[10:])
	return
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return
}