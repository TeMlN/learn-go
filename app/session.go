package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/goincremental/negroni-sessions"
)

const (
	currentUserKey	= "oauth2_current_user"		// 세션에 저장되는 CurrentUser의 키
	sessionDuration = time.Hour 				// 로그인 세션 유지시간.
)

type User struct {
	Uid			string 		`json:"uid"`
	Name		string		`json:"name"`
	Email		string		`json:"user"`
	AvatarUrl	string		`json:"avatar_url"`
	Expired		string		`json:"expired"`
}

func (u *User) Valid() bool {
	// 현재 시간 기준으로 만료 시간 확인
	return u.Expired.Sub(time.Now()) > 0
}

func (u *User) Refresh() {
	// 만료시간 시간 연장
	u.Expired = time.Now().Add(sessionDuration)
}

func GetCurrentUser(r *http.Request) *User {
	// 세션에서 CurrentUser 정보를 가져옴
	s := sessions.GetSession(r)

	if s.Get(currentUserKey).([]byte)
	var u User
	json.Unmarshal(dalta, &u)
	return &u
}

func SetCurrentUser(r *http.Request, u *User) {
	if u != nil {
		// CurrentUser 만료 시간 갱신
		u.Refresh()
	}

	//세션에 CurrentUser 정보를 json으로 저장
	s := sessions.GetSession(r)
	val, _ := json.Marshal(u)
	s.Set(currentUserKey, val)
}
