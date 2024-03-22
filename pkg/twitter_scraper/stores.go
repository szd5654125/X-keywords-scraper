package twitter_scraper

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
)

const STATE_PATH = "state.json"

type StateItems struct {
	GuestToken string         `json:"guest_token"`
	Cookie     []*http.Cookie `json:"cookie"`
}

func (s *Scraper) IsLoginState() bool {
	_, err := os.Stat(STATE_PATH)
	s.isLogged = !os.IsNotExist(err)
	return s.isLogged
}

func (s *Scraper) Load() (err error) {
	var f *os.File
	if f, err = os.OpenFile(STATE_PATH, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		return
	}
	var b []byte
	if b, err = io.ReadAll(f); err != nil {
		return
	}
	var data StateItems
	if err = json.Unmarshal(b, &data); err != nil {
		return
	}
	s.guestToken = data.GuestToken
	s.cookie = data.Cookie
	return
}

func (s *Scraper) SetGuestToken(value string) (err error) {
	var f *os.File
	if f, err = os.OpenFile(STATE_PATH, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		return
	}
	var b []byte
	if b, err = io.ReadAll(f); err != nil {
		return
	}
	var data StateItems
	if len(b) > 0 {
		if err = json.Unmarshal(b, &data); err != nil {
			return
		}
	}

	data.GuestToken = value
	encodeData, _ := json.Marshal(data)
	if err = os.WriteFile(STATE_PATH, encodeData, 0755); err != nil {
		return
	}
	return
}

func (s *Scraper) SetCookieStore(value []*http.Cookie) (err error) {
	var f *os.File
	if f, err = os.OpenFile(STATE_PATH, os.O_RDWR|os.O_CREATE, 0755); err != nil {
		return
	}
	var b []byte
	if b, err = io.ReadAll(f); err != nil {
		return
	}
	var data StateItems
	if len(b) > 0 {
		if err = json.Unmarshal(b, &data); err != nil {
			return
		}
	}
	data.Cookie = value
	encodeData, _ := json.Marshal(data)
	if err = os.WriteFile(STATE_PATH, encodeData, 0755); err != nil {
		return
	}
	return
}
