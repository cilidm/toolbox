package session

import (
	"errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-contrib/sessions/memstore"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
)

const (
	NotFound = "session_not_found"
)

type CookieSessionReq struct {
	Key    []byte
	Path   string
	MaxAge int
	Name   string
}

func NewCookieSessionReq() *CookieSessionReq {
	return &CookieSessionReq{Key: []byte("_SESSION"), Path: "/", MaxAge: 24 * 3600, Name: "_SESSION"}
}

func (c *CookieSessionReq) SetKey(key []byte) *CookieSessionReq {
	c.Key = key
	return c
}

func (c *CookieSessionReq) SetPath(path string) *CookieSessionReq {
	c.Path = path
	return c
}

func (c *CookieSessionReq) SetMaxAge(maxAge int) *CookieSessionReq {
	c.MaxAge = maxAge
	return c
}

func (c *CookieSessionReq) SetSessionName(name string) *CookieSessionReq {
	c.Name = name
	return c
}

// EnableCookieSession 使用 Cookie 保存 session
func EnableCookieSession(req *CookieSessionReq) gin.HandlerFunc {
	store := cookie.NewStore(req.Key)
	store.Options(sessions.Options{Path: req.Path, MaxAge: req.MaxAge})
	return sessions.Sessions(req.Name, store)
}

type RedisSessionReq struct {
	Size     int
	Address  string
	Password string
	Key      []byte
	Path     string
	MaxAge   int
	Name     string
}

func (r *RedisSessionReq) SetSize(size int) *RedisSessionReq {
	r.Size = size
	return r
}

func (r *RedisSessionReq) SetAddress(address string) *RedisSessionReq {
	r.Address = address
	return r
}

func (r *RedisSessionReq) SetPassword(pwd string) *RedisSessionReq {
	r.Password = pwd
	return r
}

func (r *RedisSessionReq) SetKey(key []byte) *RedisSessionReq {
	r.Key = key
	return r
}

func (r *RedisSessionReq) SetPath(path string) *RedisSessionReq {
	r.Path = path
	return r
}

func (r *RedisSessionReq) SetMaxAge(maxAge int) *RedisSessionReq {
	r.MaxAge = maxAge
	return r
}

func (r *RedisSessionReq) SetSessionName(name string) *RedisSessionReq {
	r.Name = name
	return r
}

func NewRedisSessionReq() *RedisSessionReq {
	return &RedisSessionReq{Size: 10000, Address: "127.0.0.1:6379", Password: "",
		Key: []byte("_SESSION"), Path: "/", MaxAge: 24 * 3600, Name: "_SESSION"}
}

// EnableRedisSession 使用 Redis 保存 session
func EnableRedisSession(req *RedisSessionReq) gin.HandlerFunc {
	store, _ := redis.NewStore(req.Size, "tcp", req.Address, req.Password, req.Key)
	store.Options(sessions.Options{Path: req.Path, MaxAge: req.MaxAge})
	return sessions.Sessions(req.Name, store)
}

// EnableMemorySession 使用 内存 保存 session
func EnableMemorySession(key string) gin.HandlerFunc {
	store := memstore.NewStore([]byte(key))
	store.Options(sessions.Options{Path: "/", MaxAge: 6 * 3600})
	return sessions.Sessions("_SESSION", store)
}

func Set(c *gin.Context, key string, value interface{}) (err error) {
	s := sessions.Default(c)
	s.Set(key, value)
	err = s.Save()
	if err != nil {
		return err
	}
	return nil
}

func Get(c *gin.Context, key string) interface{} {
	s := sessions.Default(c)
	return s.Get(key)
}

func Del(c *gin.Context, key string) error {
	s := sessions.Default(c)
	s.Delete(key)
	return s.Save()
}

func GetSessionId(c *gin.Context) (int64, error) {
	s := sessions.Default(c)
	auth, ok := s.Get("uid").(uint)
	if !ok {
		return 0, errors.New(NotFound)
	}
	return int64(auth), nil
}
