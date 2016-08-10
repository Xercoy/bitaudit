package bitaudit

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"os"
)

func CreateHash(auditFile, hash string) (string, error) {
	var hashResult [16]byte

	file, err := os.Open(auditFile)
	if err != nil {
		return "", err
	}

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}

	switch hash {
	case "":
		hashResult = (md5.Sum(content))
	}

	return fmt.Sprintf("%s", hashResult), nil
}
