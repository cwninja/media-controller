package messages

import "io"
import "encoding/xml"

func Pause() (io.Reader, string) {
  return newEnvelope(pauseMessage{ InstanceID: "0" }), NAMESPACE + "#Pause"
}

type pauseMessage struct {
  XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Pause"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}
