package messages

import "io"
import "encoding/xml"

func Seek(mode seekMode, target string) (io.Reader, string) {
	return newEnvelope(seekMessage{
		InstanceID: "0",
		Target:     target,
		Unit:       mode,
	}), NAMESPACE + "#Seek"
}

const (
	SEEK_REL_TIME = seekMode("REL_TIME")
)

type seekMode string

type seekMessage struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 Seek"`

	InstanceID string   `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
	Target     string   `xml:"urn:schemas-upnp-org:service:AVTransport:1 Target"`
	Unit       seekMode `xml:"urn:schemas-upnp-org:service:AVTransport:1 Unit"`
}
