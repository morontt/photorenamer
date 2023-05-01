package types

import "time"

type MediaFile interface {
	Extension() string
	OriginalFilename() string
	ParseTime() error
	DateTime() string
}

type baseMediaPart struct {
	filename string
	dateTime time.Time
}

func (m *baseMediaPart) OriginalFilename() string {
	return m.filename
}

func (m *baseMediaPart) DateTime() string {
	return m.dateTime.Format("2006-01-02 15.04.05")
}
