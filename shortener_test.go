package shortener

import (
	"fmt"
	"testing"
)

var testStorage ShortenerStorage

func init() {
	testStorage = NewShortener("test.local")
}

const (
	testUrl       = "http://example.com/1/2/3"
	rightShortUrl = "https://test.local/69c4fae842b00fdd0ac58faf135a754a"
)

func TestShorten(t *testing.T) {

	shortUrl := testStorage.Shorten(testUrl)
	if shortUrl != rightShortUrl {
		t.Fatal("shortUrl != rightShortUrl")
	}

	// request existing link
	shortUrl = testStorage.Shorten(testUrl)
	if shortUrl != rightShortUrl {
		t.Fatal("shortUrl != rightShortUrl")
	}
}

func TestResolve(t *testing.T) {

	link := testStorage.Resolve(rightShortUrl)
	if link != testUrl {
		t.Fatal("link != testUrl")
	}
}

func TestVoid(t *testing.T) {

	shortenerStorage := NewShortener("bbb.com")

	a := shortenerStorage.Shorten("https://example.org/path?foo=bar")
	fmt.Println(a)
	a = shortenerStorage.Resolve(a)
	fmt.Println(a)

	a = shortenerStorage.Shorten("otus.ru/123456")
	fmt.Println(a)
	a = shortenerStorage.Resolve(a)
	fmt.Println(a)

	a = shortenerStorage.Shorten("otus.ru/abcdef")
	fmt.Println(a)
	a = shortenerStorage.Resolve(a)
	fmt.Println(a)

}
