package config

import "strconv"

const (
  Port = 8080
  Env  = "development"
  Host = "localhost"
)

func PortString() string {
  return ":" + strconv.Itoa(Port)
}
