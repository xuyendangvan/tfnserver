package swagger

import (
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword .. HashPassword
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

//CheckPasswordHash .. CheckPasswordHash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		log.Printf(err.Error())
		return false
	}
	return true
}

func ConvertDateToString(date time.Time, format string) string {
	result := date.Format(format)
	return result
}

//ConvertStringToDate ... ConvertStringToDate
func ConvertStringToDate(date string, format string) (time.Time, error) {
	result, err := time.Parse(format, date)
	if err != nil {
		return time.Now(), err
	}
	return result, nil
}
