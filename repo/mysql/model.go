package mysql

import "time"

type Config struct {
	Read struct {
		Addr string `toml:"addr"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		Name string `toml:"name"`
	} `toml:"read"`
	Write struct {
		Addr string `toml:"addr"`
		User string `toml:"user"`
		Pass string `toml:"pass"`
		Name string `toml:"name"`
	} `toml:"write"`
	Base struct {
		MaxOpenConn     int           `toml:"maxOpenConn"`
		MaxIdleConn     int           `toml:"maxIdleConn"`
		ConnMaxLifeTime time.Duration `toml:"connMaxLifeTime"`
	} `toml:"base"`
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) SetBase(maxOpenConn, maxIdleConn int, maxLifeTime time.Duration) *Config {
	c.Base.MaxOpenConn = maxOpenConn
	c.Base.MaxIdleConn = maxIdleConn
	c.Base.ConnMaxLifeTime = maxLifeTime
	return c
}

func (c *Config) SetWrite(addr, name, user, pwd string) *Config {
	c.Write.Name = name
	c.Write.Addr = addr
	c.Write.User = user
	c.Write.Pass = pwd
	return c
}

func (c *Config) SetRead(addr, name, user, pwd string) *Config {
	c.Read.Name = name
	c.Read.Addr = addr
	c.Read.User = user
	c.Read.Pass = pwd
	return c
}
