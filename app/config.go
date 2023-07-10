package app

import "context"


type cfg struct {
	// 记录剪切板最大数量
	MaxHistory int64
}


type Config struct {
	ctx context.Context
	storeDir string
	cfg map[string]interface{}
}

func NewConfig(storeDir string) *Config {
	return &Config{
		storeDir: storeDir,
	}
}

func (c *Config) StartUp(ctx context.Context) {
	c.ctx = ctx
}

func (c *Config) Get(key string) interface{} {
	if val, ok := c.cfg[key]; ok {
		return val
	}
	return nil
}

func (c *Config) Set(key string, val interface{}) {
	c.cfg[key] = val
}

func (c *Config) GetAll() *cfg {
	return nil
}
