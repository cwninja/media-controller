package messages

import "io"
import "encoding/xml"
import "fmt"

func SetAVTransportURI(uri string) (io.Reader, string) {
	return newEnvelope(
		setAVTransportURIMessage{
			CurrentURI: uri,
			InstanceID: "0",
			CurrentURIMetaData: fmt.Sprintf(`<?xml version="1.0"?>
<DIDL-Lite xmlns="urn:schemas-upnp-org:metadata-1-0/DIDL-Lite/" xmlns:upnp="urn:schemas-upnp-org:metadata-1-0/upnp/" xmlns:dc="http://purl.org/dc/elements/1.1/" xmlns:sec="http://www.sec.co.kr/">
<item id="f-0" parentID="0" restricted="0">
    <dc:title>Video</dc:title>
    <dc:creator>Anonymous</dc:creator>
    <upnp:class>object.item.videoItem</upnp:class>
    <res protocolInfo="http-get:*:video/mp4:DLNA.ORG_OP=01;DLNA.ORG_CI=0;DLNA.ORG_FLAGS=01700000000000000000000000000000">%s</res>
</item>
</DIDL-Lite>`, uri),
		},
	), NAMESPACE + "#SetAVTransportURI"
}

type setAVTransportURIMessage struct {
	XMLName xml.Name `xml:"urn:schemas-upnp-org:service:AVTransport:1 SetAVTransportURI"`

	InstanceID         string `xml:"urn:schemas-upnp-org:service:AVTransport:1 InstanceID"`
	CurrentURI         string `xml:"urn:schemas-upnp-org:service:AVTransport:1 CurrentURI"`
	CurrentURIMetaData string `xml:"urn:schemas-upnp-org:service:AVTransport:1 CurrentURIMetaData"`
}
