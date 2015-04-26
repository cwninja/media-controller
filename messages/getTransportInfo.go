package messages

import "io"
import "encoding/xml"

func GetTransportInfo() (io.Reader, string) {
  return newEnvelope(getTransportInfoMessage{ InstanceID: "0" }), NAMESPACE + "#GetTransportInfo"
}

type getTransportInfoMessage struct {
  XMLName xml.Name                `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetTransportInfo"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}
