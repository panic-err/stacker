package main

import (
  "os"
  "fmt"
  "strconv"
  "bufio"
)

func main() {
  var boops []string
  num, err := strconv.Atoi(os.Args[1])
  if err != nil {
    panic(err)
  }

  for i := 0;i < num;i++ {
    boop := fmt.Sprint("boop")
    boops = append(boops, boop)
  }
  file, err := os.Create("boops.html")
  if err != nil {
    panic(err)
  }
  defer file.Close()
  writer := bufio.NewWriter(file)

  for i := 0;i < len(boops);i++  {
    writer.WriteString(boops[i])
  }
  writer.Flush()

}
