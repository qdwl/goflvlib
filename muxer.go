package goflvlib

import (
	"io"
	"time"
)

// Muxer is a FLV muxer.
type Muxer struct {
	W              io.Writer
	VideoTrack     *Track
	AudioTrack     *Track
	b              []byte
	filehdrwritten bool
}

func NewMuxer(w io.Writer) *Muxer {
	m := &Muxer{
		W:              w,
		b:              make([]byte, 256),
		filehdrwritten: false,
	}

	return m
}

// WriteH26x writes an H264 or an H265 access unit.
func (m *Muxer) WriteH26x(ntp time.Time, pts time.Duration, au [][]byte) error {
	return nil
}

// WriteMPEG4Audio writes MPEG-4 Audio access units.
func (m *Muxer) WriteMPEG4Audio(ntp time.Time, pts time.Duration, aus [][]byte) error {
	return nil
}
