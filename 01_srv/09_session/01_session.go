//自建全域管理session
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"sync"
	"time"
)

type Session interface {
	Set(key, value interface{}) error
	Get(key interface{}) error
	Delete(key interface{}) error
}

//---創建session管理器
//管理器provider
type Provider interface {
	SessionInit(id string) (Session, error)
	SessionRead(id string) (Session, error)
	SessionDestory(id string) (Session, error)
	GarbageCollector(maxLifeTime int64)
}

//註冊
var providers = make(map[string]Provider)

//註冊一個能透過名稱獲取的sesion provider
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil.")
	}
	if _, p := providers[name]; p {
		panic("session: Register provider is existed.")
	}
	providers[name] = provider
}

//全域session管理器
//provider封裝(sturct)
type SessionManager struct {
	cookieName  string
	lock        sync.Mutex //保證併發時資料的安全和一致性
	provider    Provider
	maxLifeTime int64
}

func NewSessionManager(providerName, cookieName string, maxLifetime int64) (*SessionManager, error) {
	provider, ok := providers[providerName]
	if !ok {
		return nil, fmt.Errorf("seesion: unknown provider %q", providerName)
	}

	return &SessionManager{
		cookieName:  cookieName,
		provider:    provider,
		maxLifeTime: maxLifetime,
	}, nil
}

var globalSession *SessionManager

func init() {
	globalSession, _ = NewSessionManager("myprovider", "sessionID", 3600)
}

//--獲取session , sessionID must be unique
func (sm *SessionManager) GetSessionId() string {
	buf := make([]byte, 32)
	//to generate cryptographically secure random number.
	if _, err := io.ReadFull(rand.Reader, buf); err != nil {
		return ""
	}
	return base64.URLEncoding.EncodeToString(buf)
}

//--若Session不存在則創建Session_step1_創建Session
func (sm *SessionManager) SessionBegin(w http.ResponseWriter, r *http.Request) (session Session) {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	cookie, err := r.Cookie(sm.cookieName)
	if err != nil || cookie.Value == "" {
		sessionID := sm.GetSessionId()
		session, _ = sm.provider.SessionInit(sessionID)
		//create a cookie
		cookie := http.Cookie{
			Name:     sm.cookieName,
			Value:    url.QueryEscape(sessionID), //escapes the string so it can be safely placed inside a URL query.
			Path:     "/",
			HttpOnly: true,
			MaxAge:   int(sm.maxLifeTime),
		}
		http.SetCookie(w, &cookie)

	} else {
		sessionID, _ := url.QueryUnescape(cookie.Value)
		session, _ = sm.provider.SessionRead(sessionID)
	}
	return session
}

//--若Session不存在則創建Session_setp2
func login(w http.ResponseWriter, r *http.Request) {
	session := globalSession.SessionBegin(w, r)
	r.ParseForm()
	name := session.Get("username")
	if name != nil {
		session.Set("username", r.Form["username"])
	}
}

//登出session
func (sm *SessionManager) SessionLogout(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie(sm.cookieName)
	if err != nil || cookie.Value == "" {
		return
	}
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.provider.SessionDestory(cookie.Value)
	newCookie := http.Cookie{
		Name:     sm.cookieName,
		Path:     "/",
		HttpOnly: true,
		Expires:  time.Now(),
		MaxAge:   -1, //logout session
	}
	http.SetCookie(w, &newCookie)
}

//刪除session
func (sm *SessionManager) DeleteSession() {
	sm.lock.Lock()
	defer sm.lock.Unlock()

	sm.provider.GarbageCollector(sm.maxLifeTime)

	//使用time計時， session逾時後自動呼叫GarbageCollector()
	time.AfterFunc(time.Duration(sm.maxLifeTime), func() {
		sm.DeleteSession()
	})
}
