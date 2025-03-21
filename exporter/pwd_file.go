package exporter

import (
	"encoding/json"
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

// LoadPwdFile reads the redis password file and returns the password map
func LoadPwdFile(passwordFile string) (map[string]string, error) {
	res := make(map[string]string)

	log.Debugf("start load password file: %s", passwordFile)
	bytes, err := os.ReadFile(passwordFile)
	if err != nil {
		log.Warnf("load password file failed: %s", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Warnf("password file format error: %s", err)
		return nil, err
	}

	log.Infof("Loaded %d entries from %s", len(res), passwordFile)
	for k := range res {
		log.Debugf("%s", k)
	}

	return res, nil
}

// BY Hurricanezwf;
type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// LoadCustomizedPwdFile reads the redis password file and returns the password map
func LoadCustomizedPwdFile(passwordFile string) (map[string]Credential, error) {
	res := make(map[string]Credential)

	log.Debugf("start load password file: %s", passwordFile)
	bytes, err := os.ReadFile(passwordFile)
	if err != nil {
		log.Warnf("load password file failed: %s", err)
		return nil, err
	}
	err = json.Unmarshal(bytes, &res)
	if err != nil {
		log.Warnf("password file format error: %s", err)
		return nil, err
	}

	log.Infof("Loaded %d entries from %s", len(res), passwordFile)
	for k := range res {
		log.Debugf("%s", k)
		// username can be empty
		if res[k].Password == "" {
			return nil, fmt.Errorf("password of %s is empty", k)
		}
	}

	return res, nil
}
