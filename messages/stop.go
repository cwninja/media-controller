package messages

import "io"
import "encoding/xml"

func Stop() (io.Reader, string) {
  return newEnvelope(stopMessage{ InstanceID: "0" }), NAMESPACE + "#Stop"
}

type stopMessage struct {
  XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Stop"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}

