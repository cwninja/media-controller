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

func (tv * TV) LoadMedia(url string) {
  tv.sendSoapMessage(messages.SetAVTransportURI(url))
}

func (tv * TV) Play(speed int) {
  tv.sendSoapMessage(messages.Play(speed))
}

func (tv * TV) Stop() {
  tv.sendSoapMessage(messages.Stop())
}

func (tv * TV) Pause() {
  tv.sendSoapMessage(messages.Pause())
}
