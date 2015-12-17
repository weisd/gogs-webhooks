package main

import (
	"github.com/BurntSushi/toml"
)

type Conf struct {
	Listen string
	Reps   map[string]DataSrc
}

type DataSrc struct {
	Name      string
	Ref       string
	Secret    string
	SrcPath   string
	AllowUser []string
}

var Setting *Conf

func InitConfig(fpath string) error {
	if _, err := toml.DecodeFile(fpath, &Setting); err != nil {
		return err
	}

	return nil
}
