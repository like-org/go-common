package store

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/like-org/common/app/service/account/model"
)

func hmacSha256(key, message string) string {
	hash := hmac.New(sha256.New, []byte(key))
	hash.Write([]byte(message))
	return hex.EncodeToString(hash.Sum([]byte("")))
}

func GenerateToken(uid int, password string) string {
	exp := time.Now().Add(time.Minute * 100).Unix()
	msg := fmt.Sprintf("%d:%d", uid, exp)

	sign := hmacSha256(password, msg)

	signMsg := fmt.Sprintf("%s$%s", msg, sign)
	token := base64.StdEncoding.EncodeToString([]byte(signMsg))
	return token
}

func (store *SQLStore) GetMemberByToken(ctx context.Context, token string) (*model.Member, error) {
	tokenStr, _ := base64.StdEncoding.DecodeString(token)
	fmt.Println("tokenStr", string(tokenStr))

	msgD := strings.Split(string(tokenStr), "$")
	if len(msgD) < 2 {
		return nil, fmt.Errorf("无效的Token")
	}
	msg := msgD[0]
	sign := msgD[1]

	msgarr := strings.Split(msg, ":")
	if len(msgarr) < 2 {
		return nil, fmt.Errorf("无效的Token")
	}
	uid := msgarr[0]
	exp := msgarr[1]

	expInt, _ := strconv.ParseInt(exp, 10, 64)
	expTime := time.Unix(expInt, 0)
	if expTime.Before(time.Now()) {
		return nil, fmt.Errorf("无效的Token")
	}

	uidInt, err := strconv.Atoi(uid)
	if err != nil {
		return nil, fmt.Errorf("无效的Token")
	}

	u, err := store.GetMemberByID(ctx, uidInt)
	if err != nil {
		return nil, fmt.Errorf("无效的Token，获取用户失败")
	}

	if sign != hmacSha256(u.Password, msg) {
		return nil, fmt.Errorf("无效的Token")
	}

	return u, nil
}
