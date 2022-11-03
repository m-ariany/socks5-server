package main

import (
	"os"

	"github.com/armon/go-socks5"
	"github.com/caarlos0/env/v6"
)

type params struct {
	User     string `env:"PROXY_USER" envDefault:""`
	Password string `env:"PROXY_PASSWORD" envDefault:""`
	Port     string `env:"PROXY_PORT" envDefault:"1080"`
}

func main() {
	// Working with app params
	cfg := params{}
	err := env.Parse(&cfg)
	if err != nil {
		println(err)
	}

	//Initialize socks5 config
	socsk5conf := &socks5.Config{}

	if cfg.User+cfg.Password != "" {
		creds := socks5.StaticCredentials{
			os.Getenv("PROXY_USER"): os.Getenv("PROXY_PASSWORD"),
		}
		cator := socks5.UserPassAuthenticator{Credentials: creds}
		socsk5conf.AuthMethods = []socks5.Authenticator{cator}
	}

	server, err := socks5.New(socsk5conf)
	if err != nil {
		panic(err)
	}

	println("Start listening proxy service on port " + cfg.Port)
	if err := server.ListenAndServe("tcp", ":"+cfg.Port); err != nil {
		panic(err)
	}
}
