package shortener

import (
	"crypto/md5"
	"fmt"
	"log"
	"net/url"
	"strings"
)

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type ShortenerStorage struct {
	Host    string
	storage map[string]string
}

func (ss *ShortenerStorage) getById(id string) string {

	if NormalLink, ok := ss.storage[id]; ok {
		return NormalLink
	}

	return ""
}

func (ss *ShortenerStorage) Shorten(Url string) string {

	u, err := url.Parse(Url)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("URL", Url)

	hash := fmt.Sprintf("%x", md5.Sum([]byte(u.String())))

	if Link := ss.getById(hash); Link != "" {
		return Link
	}

	newU := url.URL{Scheme: "https", Host: ss.Host, Path: hash}

	newShortLink := newU.String()
	ss.storage[hash] = Url

	log.Println("Short URL", newShortLink)
	return newShortLink
}

func (ss *ShortenerStorage) Resolve(Url string) string {

	u, err := url.Parse(Url)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Short URL", Url)
	hash := strings.TrimLeft(u.RequestURI(), "/")

	if Link := ss.getById(hash); Link != "" {
		log.Println("URL", Link)
		return Link
	}

	return ss.Host
}

func (ss *ShortenerStorage) Clean() {
	ss.storage = make(map[string]string)
}

func NewShortener(Host string) ShortenerStorage {

	return ShortenerStorage{Host: Host, storage: make(map[string]string)}
}
