package messages

import "testing"
import "bytes"

func TestPlay(t * testing.T) {
  reader, soapAction := Play(1)

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Play xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID><Speed xmlns="urn:schemas-upnp-org:service:AVTransport:1">1</Speed></Play></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#Play`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}

func TestSetAVTransportURI(t * testing.T) {
  reader, soapAction := SetAVTransportURI("http://example.com")

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><SetAVTransportURI xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID><CurrentURI xmlns="urn:schemas-upnp-org:service:AVTransport:1">http://example.com</CurrentURI><CurrentURIMetaData xmlns="urn:schemas-upnp-org:service:AVTransport:1"></CurrentURIMetaData></SetAVTransportURI></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#SetAVTransportURI`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}

func TestStop(t * testing.T) {
  reader, soapAction := Stop()

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Stop xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID></Stop></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#Stop`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}

func TestPause(t * testing.T) {
  reader, soapAction := Pause()

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Pause xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID></Pause></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#Pause`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}

func TestGetTransportInfo(t * testing.T) {
  reader, soapAction := GetTransportInfo()

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><GetTransportInfo xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID></GetTransportInfo></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#GetTransportInfo`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}

func TestGetPositionInfo(t * testing.T) {
  reader, soapAction := GetPositionInfo()

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><GetPositionInfo xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID></GetPositionInfo></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#GetPositionInfo`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}

func TestSeek(t * testing.T) {
  reader, soapAction := Seek(SEEK_REL_TIME, "0:0:10")

  buf := new(bytes.Buffer)
  buf.ReadFrom(reader)
  val := buf.String()

  expectedMessage := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Seek xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID><Target xmlns="urn:schemas-upnp-org:service:AVTransport:1">0:0:10</Target><Unit xmlns="urn:schemas-upnp-org:service:AVTransport:1">REL_TIME</Unit></Seek></Body></Envelope>`
  if val != expectedMessage {
    t.Error(val)
  }

  expectedSoapAction := `urn:schemas-upnp-org:service:AVTransport:1#Seek`
  if soapAction != expectedSoapAction {
    t.Error(soapAction)
  }
}
