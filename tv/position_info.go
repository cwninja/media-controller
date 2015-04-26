package tv

import "github.com/cwninja/media-controller/messages"
import "launchpad.net/xmlpath"
import "log"
import "bytes"

type PositionInfo struct {
  URI string
  Position int
  Duration int
}

func (tv * TV) GetPositionInfo() (info PositionInfo) {
  response, err := tv.sendSoapMessage(messages.GetPositionInfo())
  if err != nil {
    log.Fatal(err)
  }

  trackUriPath := xmlpath.MustCompile("//TrackURI")
  trackLengthPath := xmlpath.MustCompile("//TrackDuration")
  trackPos := xmlpath.MustCompile("//AbsTime")

  root, err := xmlpath.Parse(bytes.NewBuffer(response))
  if err != nil {
    log.Fatal(err)
  }

  if value, ok := trackUriPath.String(root); ok {
    info.URI = value
  }

  if value, ok := trackLengthPath.String(root); ok {
    info.Duration = parseTime(value)
  }

  if value, ok := trackPos.String(root); ok {
    info.Position = parseTime(value)
  }

  return
}
