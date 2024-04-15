package security

import (
	"log/slog"
	"strings"
)

// 对敏感字段进行日志脱敏

// LogValue Ciphertext  脱敏
func (ct Ciphertext) LogValue() slog.Value {
	ict := string(ct)
	if len(ict) == 1 { //  a -> *
		ict = "*"
	} else if len(ict) == 2 { // aa -> a*
		ict = ict[:1] + "*"
	} else if len(ict) > 2 { // aaac -> a**c
		ict = ict[:1] + strings.Repeat("*", len(ict)-2) + ict[len(ict)-1:]
	}
	return slog.StringValue(ict)
}

// 手机号脱敏

type PhoneNo string

func (pNo PhoneNo) LogValue() slog.Value {
	no := string(pNo)
	if len(no) == 11 { //  15345678901 -> 153****8901
		no = no[:3] + strings.Repeat("*", 4) + no[len(no)-5:]
	} else if len(no) > 2 { // 10000 -> 1***0
		no = no[:1] + strings.Repeat("*", len(no)-2) + no[len(no)-1:]
	}
	return slog.StringValue(no)
}

// 邮箱号脱敏

type Email string

func (e Email) LogValue() slog.Value {
	email := string(e)
	if email == "" {
		return slog.StringValue(email)
	}
	idx := strings.Index(email, "@")
	if idx == -1 {
		return slog.StringValue(email)
	}
	ename, domain := email[:idx], email[idx:]
	if len(ename) == 1 { //  a@qq.com -> *@qq.com
		ename = "*"
	} else if len(ename) == 2 { // aa@qq.com -> a*@qq.com
		ename = ename[:1] + "*"
	} else if len(ename) > 2 { // aaac@qq.com -> a**c@qq.com
		ename = ename[:1] + strings.Repeat("*", len(ename)-2) + ename[len(ename)-1:]
	}
	return slog.StringValue(ename + domain)
}

// 身份证号

type IDCard string

func (id IDCard) LogValue() slog.Value {
	idCard := string(id)
	if len(idCard) == 18 || len(idCard) == 15 {
		// 18 bit 522325202403312341 -> 52232520240331***1
		// 15 bit 522325202403311 -> 52232520240***1
		idCard = idCard[:len(idCard)-5] + strings.Repeat("*", 3) + idCard[len(idCard)-1:]
	}
	return slog.StringValue(idCard)
}