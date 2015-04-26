package server

import "net"
import "github.com/cwninja/media-controller/tv"
import "github.com/yasuyuky/jsonpath"
import "bufio"
import "log"
import "time"
import jsonEncoding "encoding/json"


type Server struct {
  Listener * net.TCPListener
  TV * tv.TV
}

func New(tv * tv.TV, listenAddress string) (* Server, error) {
  tcpAddress, err := net.ResolveTCPAddr("tcp", listenAddress)
  if err != nil {
    return nil, err
  }

  listener, err := net.ListenTCP("tcp", tcpAddress)
  if err != nil {
    return nil, err
  }

  server := Server{ Listener: listener, TV: tv }
  return &server, nil
}

func (s * Server) Start() {
  for {
    if c, err := s.Listener.AcceptTCP(); err == nil {
      go s.HandleConnection(c)
    }
  }
}

func (s * Server) HandleConnection(c * net.TCPConn) {
  messageScanner := bufio.NewScanner(c)
  for messageScanner.Scan() {
    if json, err := jsonpath.DecodeString(messageScanner.Text()); err == nil {
      command, _ := jsonpath.GetString(json, []interface{}{"command"}, "")
      if command == "stop" || command == "exit" {
        s.TV.Stop()
      } else if command == "pause" {
        s.TV.Pause()
      } else if command == "play" {
        url, _ := jsonpath.GetString(json, []interface{}{"url"}, "")
        position, _ := jsonpath.GetNumber(json, []interface{}{"position"}, 0)
        log.Printf("No seek support implemnted: %f", position)
        s.TV.Play(url)
      } else if command == "status" {
        if data, err := jsonEncoding.Marshal(s.TV.Status()); err == nil {
          c.Write(data)
          c.Write([]byte{'\n'})
        } else {
          log.Fatal(err)
        }
      } else if command == "monitor" {
        for {
          if data, err := jsonEncoding.Marshal(s.TV.Status()); err == nil {
            c.Write(data)
            c.Write([]byte{'\n'})
          } else {
            log.Fatal(err)
          }
          time.Sleep(time.Second)
        }
      }
    }
  }
}
