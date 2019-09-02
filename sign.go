package wxsign

import (
	"crypto/sha1"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

var (
	chars = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

// RandString
func RandString(l int) string {
	bs := []byte{}
	for i := 0; i < l; i++ {
		bs = append(bs, chars[rand.Intn(len(chars))])
	}
	return string(bs)
}

// GetJsSign GetJsSign
func (wSign *WxSign) GetJsSign(url string, proxy_flag bool, proxy_url string) (*WxJsSign, error) {
	jsTicket, err := wSign.GetTicket(proxy_flag, proxy_url)
	if err != nil {
		return nil, err
	}
	// splite url
	urlSlice := strings.Split(url, "#")
	jsSign := &WxJsSign{
		Appid:     wSign.Appid,
		Noncestr:  RandString(16),
		Timestamp: strconv.FormatInt(time.Now().UTC().Unix(), 10),
		Url:       urlSlice[0],
	}
	jsSign.Signature = Signature(jsTicket, jsSign.Noncestr, jsSign.Timestamp, jsSign.Url)
	return jsSign, nil
}

// Signature
func Signature(jsTicket, noncestr, timestamp, url string) string {
	h := sha1.New()
	h.Write([]byte(fmt.Sprintf("jsapi_ticket=%s&noncestr=%s&timestamp=%s&url=%s", jsTicket, noncestr, timestamp, url)))
	return fmt.Sprintf("%x", h.Sum(nil))
}
