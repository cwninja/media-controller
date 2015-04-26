package tv

import "github.com/cwninja/media-controller/messages"
import "launchpad.net/xmlpath"
import "net/http"
import "io"
import "os"
import "io/ioutil"
import "fmt"
import "log"
import "errors"
import "bytes"

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

func (tv * TV) sendSoapMessage(body io.Reader, action string) (responseBody []byte, err error) {
  request, err := http.NewRequest("POST", tv.Url, body)
  if err != nil {
    log.Fatal(err)
  }

  request.Header.Set("Content-Type", "text/xml")
  request.Header.Set("SOAPACTION", fmt.Sprintf(`"%s"`, action))

  response, err := tv.Client.Do(request)
  if err != nil {
    log.Fatal(err)
  }

  if response.StatusCode != 200 {
    io.Copy(os.Stderr, response.Body)
    log.Fatal(response.Status)
    return responseBody, errors.New(response.Status)
  }

  defer response.Body.Close()

  responseBody, err = ioutil.ReadAll(response.Body)
  if err != nil {
    return responseBody, err
  }

  return responseBody, nil
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

type PlaybackStatus string

const (
  STATUS_PAUSED = PlaybackStatus("PAUSED_PLAYBACK")
  STATUS_PLAYING = PlaybackStatus("PLAYING")
  STATUS_STOPPED = PlaybackStatus("STOPPED")
  STATUS_UNKNOWN = PlaybackStatus("UNKNOWN")
)

func (tv * TV) GetTransportInfo() PlaybackStatus {
  response, err := tv.sendSoapMessage(messages.GetTransportInfo())
  if err != nil {
    log.Fatal(err)
  }

  currentTransportStatePath := xmlpath.MustCompile("//CurrentTransportState")

  root, err := xmlpath.Parse(bytes.NewBuffer(response))
  if err != nil {
    log.Fatal(err)
  }

  if value, ok := currentTransportStatePath.String(root); ok {
    return PlaybackStatus(value)
  }

  return STATUS_UNKNOWN
}

func parseTime(str string) int {
  var h, m, s int
  _, err := fmt.Sscanf(str, "%d:%d:%d", &h, &m, &s)
  if err != nil {
    log.Fatal(err)
  }
  return h * 60 * 60 + m * 60 + s
}

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
