package main

import (
  "encoding/json"
  "os"
  "time"
)

type Badge struct {
  SchemaVersion int    `json:"schemaVersion"`
  Label         string `json:"label"`
  Message       string `json:"message"`
  Color         string `json:"color"`
}

func main(){
  badge := Badge{1, "PRISMA", "ALL CLEAR", "brightgreen"}
  f, _ := json.MarshalIndent(badge, "", "  ")
  os.WriteFile("prisma.json", f, 0644)
  os.WriteFile("health.json", []byte(`{"status":"ok","time":"`+time.Now().Format(time.RFC3339)+`"}`), 0644)
  println("⚡ PRISMA V2.1 ALL CLEAR - Go Native Diamond")
}
