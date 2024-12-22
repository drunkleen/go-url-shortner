package shortener

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const UserId = "d3b07384-fc8c-4a1b-b11c-364a29dd0f45"

func TestShortLinkGenerator(t *testing.T) {
	initialLink_1 := "https://www.youtube.com/@drunkleen/"
	shortLink_1 := GenerateShortLink(initialLink_1, UserId)

	initialLink_2 := "https://www.youtube.com/@drunkleen/community"
	shortLink_2 := GenerateShortLink(initialLink_2, UserId)

	initialLink_3 := "https://techhub.social/@drunkleen"
	shortLink_3 := GenerateShortLink(initialLink_3, UserId)

	assert.Equal(t, shortLink_1, "cWeetHYM")
	assert.Equal(t, shortLink_2, "6uUfWi2b")
	assert.Equal(t, shortLink_3, "LGNFLMUN")
}
