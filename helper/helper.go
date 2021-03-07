package swagger

import (
	"bytes"
	"encoding/base64"
	"errors"
	"image"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"golang.org/x/crypto/bcrypt"

	"image/gif"
	"image/jpeg"
	"image/png"
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

var (
	ErrBucket       = errors.New("Invalid bucket!")
	ErrSize         = errors.New("Invalid size!")
	ErrInvalidImage = errors.New("Invalid image!")
	FormatJPEG      = "/9j"
	FormatPNG       = "iVB"
	FormatGIF       = "R0l"
)

func SaveImageToDisk(fileNameBase, data string) (string, error) {
	idx := strings.Index(data, ";base64,")
	if idx < 0 {
		return "", ErrInvalidImage
	}
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(data[idx+8:]))
	buff := bytes.Buffer{}
	_, err := buff.ReadFrom(reader)
	if err != nil {
		return "", err
	}
	imgCfg, fm, err := image.DecodeConfig(bytes.NewReader(buff.Bytes()))
	if err != nil {
		return "", err
	}

	if imgCfg.Width != 750 || imgCfg.Height != 685 {
		return "", ErrSize
	}

	fileName := fileNameBase + "." + fm
	ioutil.WriteFile(fileName, buff.Bytes(), 0644)

	return fileName, err
}

func SaveToFile(data string, name string, host string, folder string) *string {
	if len(data) > 3 {
		loc, _ := time.LoadLocation("MST")
		now := time.Now().In(loc)
		fileName := name + now.Format("20060102150405")
		path := "./image/" + fileName
		ImageType := data[0:3]
		log.Println(ImageType)
		unbased, err := base64.StdEncoding.DecodeString(data)
		if err != nil {
			log.Println("Cannot decode b64")
			return nil
		}
		r := bytes.NewReader(unbased)
		switch ImageType {
		case FormatPNG:
			path = path + ".png"
			im, err := png.Decode(r)
			if err != nil {
				log.Println("Bad png")
			}

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				log.Println("Cannot open file")
			}

			png.Encode(f, im)
			fileName += ".png"
		case FormatJPEG:
			path = path + ".jpeg"
			log.Println(path)
			im, err := jpeg.Decode(r)
			if err != nil {
				log.Println("Bad jpeg")
			}

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				log.Println("Cannot open file")
			}
			jpeg.Encode(f, im, nil)
			fileName += ".jpeg"
		case FormatGIF:
			path = path + ".gif"
			im, err := gif.Decode(r)
			if err != nil {
				log.Println("Bad gif")
			}

			f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0777)
			if err != nil {
				log.Println("Cannot open file")
			}

			gif.Encode(f, im, nil)
			fileName += ".gif"
		default:
			fileName = ""
		}
		result := host + folder + fileName
		return &result
	}
	return nil
}
