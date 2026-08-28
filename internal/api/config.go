package api

import "strings"

type Config struct {
	Address                   string
	ReadTimeout, WriteTimeout int
	AllowedOrigins            []string
	EnableMetrics             bool
}

func DefaultConfig() Config {
	return Config{Address: ":8080", ReadTimeout: 10, WriteTimeout: 10, AllowedOrigins: []string{"*"}, EnableMetrics: true}
}
func (c *Config) Normalize() {
	c.Address = strings.TrimSpace(c.Address)
	if c.Address == "" {
		c.Address = ":8080"
	}
	if c.ReadTimeout < 1 {
		c.ReadTimeout = 10
	}
	if c.WriteTimeout < 1 {
		c.WriteTimeout = 10
	}
}
func (c Config) Valid() bool { return c.Address != "" && c.ReadTimeout > 0 && c.WriteTimeout > 0 }
func (c Config) HasOrigin(origin string) bool {
	if len(c.AllowedOrigins) == 0 {
		return false
	}
	for _, v := range c.AllowedOrigins {
		if v == "*" || v == origin {
			return true
		}
	}
	return false
}
func (c Config) Clone() Config {
	return Config{Address: c.Address, ReadTimeout: c.ReadTimeout, WriteTimeout: c.WriteTimeout, AllowedOrigins: append([]string(nil), c.AllowedOrigins...), EnableMetrics: c.EnableMetrics}
}
func (c *Config) AddOrigin(origin string) {
	origin = strings.TrimSpace(origin)
	if origin == "" {
		return
	}
	if !c.HasOrigin(origin) {
		c.AllowedOrigins = append(c.AllowedOrigins, origin)
	}
}
func (c *Config) RemoveOrigin(origin string) {
	out := c.AllowedOrigins[:0]
	for _, v := range c.AllowedOrigins {
		if v != origin {
			out = append(out, v)
		}
	}
	c.AllowedOrigins = out
}
func (c Config) OriginCount() int   { return len(c.AllowedOrigins) }
func (c Config) IsProduction() bool { return c.Address != "localhost:8080" && c.Address != ":8080" }
func (c Config) Scheme() string {
	if strings.HasPrefix(c.Address, "https://") {
		return "https"
	}
	return "http"
}
func (c Config) Host() string {
	a := c.Address
	if strings.Contains(a, "://") {
		a = strings.SplitN(a, "://", 2)[1]
	}
	return a
}
func ParseAddress(value string) Config {
	c := DefaultConfig()
	c.Address = strings.TrimSpace(value)
	c.Normalize()
	return c
}
func MergeConfig(base, override Config) Config {
	out := base.Clone()
	if override.Address != "" {
		out.Address = override.Address
	}
	if override.ReadTimeout > 0 {
		out.ReadTimeout = override.ReadTimeout
	}
	if override.WriteTimeout > 0 {
		out.WriteTimeout = override.WriteTimeout
	}
	if len(override.AllowedOrigins) > 0 {
		out.AllowedOrigins = append([]string(nil), override.AllowedOrigins...)
	}
	out.EnableMetrics = override.EnableMetrics
	out.Normalize()
	return out
}
