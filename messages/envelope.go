package messages

import "encoding/xml"
import "io"
import "bytes"

func newEnvelope(payload interface{}) (out io.Reader) {
  byteArray, _ := xml.Marshal(
    envelope{
      EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
      Body: envelopeBody{ Payload: payload },
    },
  )
  return bytes.NewReader(byteArray)
}

type envelope struct {
    XMLName       xml.Name       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
    EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
    Body          envelopeBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type envelopeBody struct {
  Payload interface{}
}

const (
  NAMESPACE = "urn:schemas-upnp-org:service:AVTransport:1"
)

