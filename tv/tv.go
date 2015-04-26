package tv

import "github.com/cwninja/media-controller/messages"
import "net/http"

type TV struct {
  Url string
  Client * http.Client
}

func NewTV(url string) TV {
  return TV{
    Url: url,
    Client: &http.Client{},
  }
}

func (tv * TV) PlayFrom(url string, seconds int) {
  tv.sendSoapMessage(messages.Stop())
  tv.sendSoapMessage(messages.SetAVTransportURI(url))
  tv.sendSoapMessage(messages.Play(1))
  tv.sendSoapMessage(messages.Seek(messages.SEEK_REL_TIME, formatTime(seconds)))
}

func (tv * TV) Play(url string) {
  tv.sendSoapMessage(messages.Stop())
  tv.sendSoapMessage(messages.SetAVTransportURI(url))
  tv.sendSoapMessage(messages.Play(1))
}

func (tv * TV) SeekTo(seconds int) {
  tv.sendSoapMessage(messages.Seek(messages.SEEK_REL_TIME, formatTime(seconds)))
}

func (tv * TV) SeekBy(seconds int) {
  seconds = seconds + tv.Status().Position
  tv.sendSoapMessage(messages.Seek(messages.SEEK_REL_TIME, formatTime(seconds)))
}

func (tv * TV) Stop() {
  tv.sendSoapMessage(messages.Stop())
}

func (tv * TV) Pause() {
  status := tv.GetTransportInfo()
  if status == STATUS_PAUSED {
    tv.sendSoapMessage(messages.Play(1))
  } else if status == STATUS_PLAYING {
    tv.sendSoapMessage(messages.Pause())
  }
}

func (tv * TV) Status() Status {
  positionInfo := tv.GetPositionInfo()
  transportInfo := tv.GetTransportInfo()

  if transportInfo == STATUS_STOPPED {
    return Status{}
  } else {
    return Status{
      Paused: transportInfo == STATUS_PAUSED,
      Length: positionInfo.Duration,
      Position: positionInfo.Position,
      Url: positionInfo.URI,
    }
  }
}

type Status struct {
  Paused bool `json:"paused"`
  Length int `json:"length"`
  Position int `json:"position"`
  Url string `json:"url"`
}
