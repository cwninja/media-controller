package main

import "github.com/cwninja/media-controller/tv"
import "flag"
import "log"
import "fmt"
import "sync"
import "os"
import urlParser "net/url"
import "net/http"
import "net"
import "time"

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
    var url string
    var wg sync.WaitGroup

    file := flag.Arg(1)
    if _, err := os.Stat(file); err != nil {
      url = file
    }

    myTv.Stop()

    if url == "" {
      http.HandleFunc("/file", func(w http.ResponseWriter, r *http.Request){
        http.ServeFile(w, r, file)
      })

      parsedTvUrl, err := urlParser.Parse(myTv.Url)
      if err != nil {
        log.Fatal("Bad URL for TV")
      }

      connection, err := net.Dial("tcp", parsedTvUrl.Host)
      if err != nil {
        log.Fatal("Could not connect to TV")
      }
      localAddress := connection.LocalAddr()
      connection.Close()

      url = fmt.Sprintf("http://%s/file", localAddress.String())
      log.Printf("Serving from URL: %s", url)
      wg.Add(1)
      go http.ListenAndServe(localAddress.String(), nil)
      go func(){
        for {
          time.Sleep(time.Second)
          if ( myTv.GetTransportInfo() == tv.STATUS_STOPPED ) {
            os.Exit(0)
          }
        }
      }()
    }

    myTv.LoadMedia(url)
    myTv.Play(1)
    wg.Wait()
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
