package messages

import "io"
import "encoding/xml"

func GetPositionInfo() (io.Reader, string) {
  return newEnvelope(getPositionInfoMessage{ InstanceID: "0" }), NAMESPACE + "#GetPositionInfo"
}

type getPositionInfoMessage struct {
  XMLName xml.Name                `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetPositionInfo"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}

