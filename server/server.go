package server

import "net"
import "github.com/cwninja/media-controller/tv"
import "github.com/yasuyuky/jsonpath"
import "bufio"
import "log"


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
      v, _ := jsonpath.GetString(json, []interface{}{"type"}, "")
      log.Printf(v)
    }
  }
}
