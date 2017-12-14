package util

/*
#cgo LDFLAGS: -L./cgo/lib/  -ldecrypter
#cgo CFLAGS: -I./cgo/include/
#include "decrypter.h"
*/

import "C"
import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

const (
	ACTIVE  = 1
	PAUSED  = 2
	DELETED = 3
	PENDING = 4
)

const (
	ANDROIDPLATFORM = 1
	IOSPLATFORM     = 2
)

func SliceIndex(limit int, predicate func(i int) bool) int {
	for i := 0; i < limit; i++ {
		if predicate(i) {
			return i
		}
	}
	return -1
}

//对字符串进行SHA1哈希
func Sha1(data string) string {
	t := sha1.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func Md5(data string) string {
	t := md5.New()
	io.WriteString(t, data)
	return fmt.Sprintf("%x", t.Sum(nil))
}

func Base64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func Base64Encode(data string) string {
	base64Data := Base64([]byte(data))
	result := []rune(base64Data)
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	charsto := "vSoajc7dRzpWifGyNxZnV5k+DHLYhJ46lt0U3QrgEuq8sw/XMeBAT2Fb9P1OIKmC"
	charstoArr := []rune(charsto)
	lenData := len(base64Data)
	for i := 0; i < lenData; i++ {
		index := strings.IndexRune(chars, result[i])
		if index >= 0 {
			result[i] = charstoArr[index]
		}
	}
	return string(result)
}

func Base64Decode(data string) string {
	result := []rune(data)
	chars := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	charsto := "vSoajc7dRzpWifGyNxZnV5k+DHLYhJ46lt0U3QrgEuq8sw/XMeBAT2Fb9P1OIKmC"
	charsArr := []rune(chars)
	lenData := len(data)
	for i := 0; i < lenData; i++ {
		index := strings.IndexRune(charsto, result[i])
		if index >= 0 {
			result[i] = charsArr[index]
		}
	}
	decodeStr := string(result)
	de64, err := DeBase64([]byte(decodeStr))
	if err != nil {
		return ""
	}
	if strings.TrimRight(Base64(de64), "=") == strings.TrimRight(decodeStr, "=") {
		return string(de64)
	}
	return ""
}

func DeBase64(src []byte) ([]byte, error) {
	return base64.StdEncoding.DecodeString(string(src))
}

func base64url_encode(data []byte) string {
	ret := base64.StdEncoding.EncodeToString(data)
	return strings.Map(func(r rune) rune {
		switch r {
		case '+':
			return '-'
		case '/':
			return '_'
		}

		return r
	}, ret)
}

func base64url_decode(s string) ([]byte, error) {
	base64Str := strings.Map(func(r rune) rune {
		switch r {
		case '-':
			return '+'
		case '_':
			return '/'
		}

		return r
	}, s)

	if pad := len(base64Str) % 4; pad > 0 {
		base64Str += strings.Repeat("=", 4-pad)
	}

	return base64.StdEncoding.DecodeString(base64Str)
}

//判断是否是字符串还是数字
func IsNum(str string) bool {
	isNum, _ := regexp.MatchString("(^[0-9]+$)", str)
	return isNum
}

func GetStrFromNum(num int64) string {
	return strconv.FormatInt(num, 10)
}

func RequestHeader(req *http.Request, key string) string {
	if values, _ := req.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}

func GetRequestID() string {
	return bson.NewObjectId().Hex()
}

func ClientIP(req *http.Request) string {
	forwardedByClientIP := true
	if forwardedByClientIP {
		clientIP := strings.TrimSpace(RequestHeader(req, "X-Real-Ip"))
		if len(clientIP) > 0 {
			return clientIP
		}
		clientIP = RequestHeader(req, "X-Forwarded-For")
		if index := strings.IndexByte(clientIP, ','); index >= 0 {
			clientIP = clientIP[0:index]
		}
		clientIP = strings.TrimSpace(clientIP)
		if len(clientIP) > 0 {
			return clientIP
		}
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}

func GetServerIP(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// VerCampare 版本比较
func VerCampare(ver1 string, ver2 string) (int, error) {
	IntVer1, err := IntVer(ver1)
	if err != nil {
		return -2, err
	}
	IntVer2, err := IntVer(ver2)
	if err != nil {
		return -2, err
	}
	if IntVer1 == IntVer2 {
		return 0, nil
	}
	if IntVer1 > IntVer2 {
		return 1, nil
	} else {
		return -1, nil
	}
}

// IntVer 把字符串版本转化成一个数字版本
func IntVer(v string) (int64, error) {
	sections := strings.Split(v, ".")
	intVerSection := func(v string, n int) string {
		if n < len(sections) {
			return fmt.Sprintf("%02s", sections[n])
		} else {
			return "00"
		}
	}
	s := ""
	for i := 0; i < 4; i++ {
		s += intVerSection(v, i)
	}
	return strconv.ParseInt(s, 10, 64)
}

func InArray(val int, array []int) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

func InArray64(val int64, array []int64) bool {
	for _, v := range array {
		if val == v {
			return true
		}
	}
	return false
}

func GetInternalIP() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("have get on ip info")
}