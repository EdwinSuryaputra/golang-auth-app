package publicfacingutil

import (
	"golang-auth-app/app/interfaces/errorcode"

	"github.com/rotisserie/eris"
	"github.com/speps/go-hashids/v2"
)

var salt = "jKHrbgzTbPXwXiGmSox2fca3y6ZE6k0x"

func initHashId() *hashids.HashID {
	hd := hashids.NewData()
	hd.Salt = salt
	hd.MinLength = 8
	h, _ := hashids.NewWithData(hd)

	return h
}

func Encode(id int32) (string, error) {
	encoded, err := initHashId().Encode([]int{int(id)})
	if err != nil {
		return "", eris.Wrap(err, "error occurred during encode id public facing")
	}

	return encoded, nil
}

func Decode(id string) (int32, error) {
	if id == "" {
		return 0, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidId, "unable to decode empty id")
	}

	decoded, err := initHashId().DecodeWithError(id)
	if err != nil {
		return 0, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidId, "invalid public facing id")
	} else if len(decoded) < 1 {
		return 0, errorcode.WithCustomMessage(errorcode.ErrCodeInvalidId, "no decode result")
	}

	return int32(decoded[0]), nil
}
