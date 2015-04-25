package messages

import "encoding/xml"
import "io"
import "bytes"

func SetAVTransportURI(uri string) (io.Reader, string) {
  return newEnvelope(
    setAVTransportURI{
      CurrentURI: uri,
      InstanceID: "0",
    },
  ), NAMESPACE + "#SetAVTransportURI"
}

func Play(speed int) (io.Reader, string) {
  return newEnvelope(
    play{
      Speed: speed,
      InstanceID: "0",
    },
  ), NAMESPACE + "#Play"
}

func Stop() (io.Reader, string) {
  return newEnvelope(stop{ InstanceID: "0" }), NAMESPACE + "#Stop"
}

func Pause() (io.Reader, string) {
  return newEnvelope(pause{ InstanceID: "0" }), NAMESPACE + "#Pause"
}

func GetTransportInfo() (io.Reader, string) {
  return newEnvelope(getTransportInfo{ InstanceID: "0" }), NAMESPACE + "#GetTransportInfo"
}

func GetPositionInfo() (io.Reader, string) {
  return newEnvelope(getPositionInfo{ InstanceID: "0" }), NAMESPACE + "#GetPositionInfo"
}

const (
  NAMESPACE = "urn:schemas-upnp-org:service:AVTransport:1"
)

type envelope struct {
    XMLName       xml.Name       `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
    EncodingStyle string         `xml:"http://schemas.xmlsoap.org/soap/envelope/ encodingStyle,attr"`
    Body          envelopeBody   `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
}

type envelopeBody struct {
  Payload interface{}
}

func newEnvelope(payload interface{}) (out io.Reader) {
  byteArray, _ := xml.Marshal(
    envelope{
      EncodingStyle: "http://schemas.xmlsoap.org/soap/encoding/",
      Body: envelopeBody{ Payload: payload },
    },
  )
  return bytes.NewReader(byteArray)
}

type play struct {
  XMLName xml.Name               `xml:"urn:schemas-upnp-org:service:AVTransport:1 Play"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
  Speed  int                     `xml:"urn:schemas-upnp-org:service:AVTransport:1 Speed"`
}

type pause struct {
  XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Pause"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}

type stop struct {
  XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Stop"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}

type getTransportInfo struct {
  XMLName xml.Name                `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetTransportInfo"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}

type getPositionInfo struct {
  XMLName xml.Name                `xml:"urn:schemas-upnp-org:service:AVTransport:1 GetPositionInfo"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
}

type setAVTransportURI struct {
  XMLName xml.Name               `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetAVTransportURI"`

  InstanceID  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
  CurrentURI  string             `xml:"urn:schemas-upnp-org:service:AVTransport:1 CurrentURI"`
  CurrentURIMetaData  string     `xml:"urn:schemas-upnp-org:service:AVTransport:1 CurrentURIMetaData"`
}

