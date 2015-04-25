package main

import "github.com/cwninja/media-controller/tv"
import "flag"
import "log"

func main() {
  tvUrl := flag.String("tv", "http://192.168.0.120:55000/dmr/control_2", "URL for TV.")
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
    myTv.Pause()
  } else if command == "resume" {
    myTv.Play(1)
  } else if command == "stop" {
    myTv.Stop()
  } else if command == "info" {
    myTv.GetTransportInfo()
  } else if command == "pos" {
    myTv.GetPositionInfo()
  }
}
