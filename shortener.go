package main

type Shortener interface {
	Shorten(url string) string
	Resolve(url string) string
}

type ShortenerStorage struct {
	storage map[string]string
}

func (ss *ShortenerStorage) Shorten(url string) string {
	return ""
}

func (ss *ShortenerStorage) Resolve(url string) string {
	return ""
}

func main() {

}
