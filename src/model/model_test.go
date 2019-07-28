package model

import (
	"config"
	"middleware"

	"github.com/stretchr/testify/assert"

	"testing"
)

func TestLibraryInsert(t *testing.T) {
	assert := assert.New(t)
	config.DoInit()
	middleware.DoInit()

	geo := Geolocation{
		IP:           "127.0.0.1",
		Ccode:        "IN",
		Country:      "India",
		City:         "Delhi",
		Lat:          "28.7041",
		Long:         "77.1025",
		MysteryValue: "123456789",
	}

	ok := geo.Validate()

	assert.Equal(ok, true, "Error generated for Valid payload")

	geo.Handler = middleware.DBHandler
	err := geo.Insert()

	assert.Nil(err)

	err = geo.Get()

	assert.Nil(err)

	err = geo.Delete()

	assert.Nil(err)
}
