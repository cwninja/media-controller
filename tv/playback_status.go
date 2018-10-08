package tv

import "github.com/cwninja/media-controller/messages"
import "launchpad.net/xmlpath"
import "log"
import "bytes"

type PlaybackStatus string

const (
	STATUS_PAUSED  = PlaybackStatus("PAUSED_PLAYBACK")
	STATUS_PLAYING = PlaybackStatus("PLAYING")
	STATUS_STOPPED = PlaybackStatus("STOPPED")
	STATUS_UNKNOWN = PlaybackStatus("UNKNOWN")
)

var currentTransportStatePath = xmlpath.MustCompile("//CurrentTransportState")

func (tv *TV) GetTransportInfo() PlaybackStatus {
	response, err := tv.sendSoapMessage(messages.GetTransportInfo())
	if err != nil {
		log.Printf("SOAP Error: %s", err)
		return STATUS_UNKNOWN
	}

	root, err := xmlpath.Parse(bytes.NewBuffer(response))
	if err != nil {
		log.Printf("Bad XML from server: %s", err)
		return STATUS_UNKNOWN
	}

	if value, ok := currentTransportStatePath.String(root); ok {
		log.Printf("Transport state: %s", value)
		return PlaybackStatus(value)
	}

	return STATUS_UNKNOWN
}
