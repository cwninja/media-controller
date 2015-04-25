package messages

import "testing"

func TestPlay(t * testing.T) {
  expected := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><Play xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID><Speed xmlns="urn:schemas-upnp-org:service:AVTransport:1">1</Speed></Play></Body></Envelope>`
  if val := string(Play(1)); val != expected {
    t.Error(val)
  }
}

func TestSetAVTransportURI(t * testing.T) {
  expected := `<Envelope xmlns="http://schemas.xmlsoap.org/soap/envelope/" xmlns:envelope="http://schemas.xmlsoap.org/soap/envelope/" envelope:encodingStyle="http://schemas.xmlsoap.org/soap/encoding/"><Body xmlns="http://schemas.xmlsoap.org/soap/envelope/"><SetAVTransportURI xmlns="urn:schemas-upnp-org:service:AVTransport:1"><InstanceID xmlns="urn:schemas-upnp-org:service:AVTransport:1">0</InstanceID><CurrentURI xmlns="urn:schemas-upnp-org:service:AVTransport:1">http://example.com</CurrentURI><CurrentURIMetaData xmlns="urn:schemas-upnp-org:service:AVTransport:1"></CurrentURIMetaData></SetAVTransportURI></Body></Envelope>`
  if val := string(SetAVTransportURI("http://example.com")); val != expected {
    t.Error(val)
  }
}

