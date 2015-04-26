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

func (tv * TV) Play(url string) {
  tv.sendSoapMessage(messages.Stop())
  tv.sendSoapMessage(messages.SetAVTransportURI(url))
  tv.sendSoapMessage(messages.Play(1))
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
