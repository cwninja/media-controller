package main

import "github.com/cwninja/media-controller/tv"
import "flag"
import "log"
import "fmt"
import "os"

func main() {
  log.SetFlags(0)
  tvUrl := flag.String("tv", os.Getenv("TV_CONTROL_URL"), "URL for TV.")
  flag.Parse()

  if flag.NArg() < 1 {
    log.Fatal("Please provide at least a command")
  }

  myTv := tv.NewTV(*tvUrl)

  command := flag.Arg(0)
  if command == "play" {
    myTv.Stop()
    myTv.LoadMedia(flag.Arg(1))
    myTv.Play(1)
  } else if command == "pause" {
    status := myTv.GetTransportInfo()
    if status == tv.STATUS_PAUSED {
      myTv.Play(1)
    } else if status == tv.STATUS_PLAYING {
      myTv.Pause()
    } else {
      log.Fatal("Not playing")
    }
  } else if command == "stop" {
    myTv.Stop()
  } else if command == "info" {
    status := myTv.GetTransportInfo()
    posInfo := myTv.GetPositionInfo()
    fmt.Printf("Url: %s\n%s  -  Progress: %d/%d\n", posInfo.URI, status, posInfo.Position, posInfo.Duration)
  } else {
    log.Fatal("Unknown command")
  }
}
