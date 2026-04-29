package main

import (
 "fmt"
 "os"
 "time"

 "github.com/labstack/echo/v4"
 "github.com/labstack/gommon/log"

 "webapp/pkg/config"
)

func main() {
 port := config.GetEnv("PORT", "3000")
 logPath := config.GetEnv("LOG_PATH", "/app/log/app.log")

 logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
 if err != nil {
  panic(fmt.Sprintf("failed to open log file %s: %v", logPath, err))
 }
 defer logFile.Close()

 e := echo.New()

 e.Logger.SetOutput(logFile)
 e.Logger.SetLevel(log.ERROR)

 e.Static("/", "public")

 e.GET("/", func(c echo.Context) error {
  return c.File("public/views/webapp.html")
 })

 go func() {
  for {
   e.Logger.Error("GET")
   time.Sleep(time.Second)
  }
 }()

 if err := e.Start(":" + port); err != nil {
  e.Logger.Error(err)
 }
}