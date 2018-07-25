package common

import (
	"github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func InitLogger(){
	zerolog.TimeFieldFormat = ""
}

func Logger(msg string) {
    log.Print(msg)
}

func CheckErr(err error) {
  if err != nil {
	  log.Print(err)
  }
}