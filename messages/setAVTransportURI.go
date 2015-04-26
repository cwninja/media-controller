package messages

import "io"
import "encoding/xml"

func SetAVTransportURI(uri string) (io.Reader, string) {
  return newEnvelope(
    setAVTransportURIMessage{
      CurrentURI: uri,
      InstanceID: "0",
    },
  ), NAMESPACE + "#SetAVTransportURI"
}

type setAVTransportURIMessage struct {
  XMLName xml.Name               `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetAVTransportURI"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
  CurrentURI  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 CurrentURI"`
  CurrentURIMetaData  string     `xml:"urn:schemas-upnp-org:service:AVTransport:1 CurrentURIMetaData"`
}

