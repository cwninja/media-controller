package messages

import "io"
import "encoding/xml"

func Play(speed int) (io.Reader, string) {
  return newEnvelope(
    playMessage{
      Speed: speed,
      InstanceID: "0",
    },
  ), NAMESPACE + "#Play"
}

type playMessage struct {
  XMLName xml.Name               `xml:"urn:schemas-upnp-org:service:AVTransport:1 Play"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
  Speed  int                     `xml:"urn:schemas-upnp-org:service:AVTransport:1 Speed"`
}

