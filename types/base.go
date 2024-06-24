package types

import (
	"crypto/sha1"
	"fmt"
	"io"
	"log"
	"os"
	"time"
)

type baseMediaPart struct {
	filename   string
	dateTime   time.Time
	cachedHash string
}

func (m *baseMediaPart) OriginalFilename() string {
	return m.filename
}

func (m *baseMediaPart) DateTime() string {
	return m.dateTime.Format("2006-01-02 15.04.05")
}

func (m *baseMediaPart) setTimeByString(timeString string) error {
	t, err := parseTimeString(timeString)
	if err != nil {
		return err
	}

	m.dateTime = t.In(time.Local)

	return nil
}

func (m *baseMediaPart) Hash() string {
	if len(m.cachedHash) > 0 {
		return m.cachedHash
	}

	f, err := os.Open(m.filename)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	h := sha1.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	m.cachedHash = fmt.Sprintf("%x", h.Sum(nil))

	return m.cachedHash
}
