package types

type MediaFile interface {
	Extension() string
	OriginalFilename() string
	ParseTime() error
	DateTime() string
}

type baseMediaPart struct {
	filename string
	dateTime string
}

func (m *baseMediaPart) OriginalFilename() string {
	return m.filename
}

func (m *baseMediaPart) DateTime() string {
	return m.dateTime
}
