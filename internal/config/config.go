package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf

	Upload struct {
		Dir      string
		Filename string
	}

	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
}
