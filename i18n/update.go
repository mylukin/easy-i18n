package i18n

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

// Update messages
func Update(id, srcFile, destFile string) error {
	if len(srcFile) == 0 {
		return fmt.Errorf(Sprintf(id, "srcFile cannot be empty"))
	}

	if len(destFile) == 0 {
		return fmt.Errorf(Sprintf(id, "destFile cannot be empty"))
	}

	srcMessages, err := unmarshal(id, srcFile)
	if err != nil {
		return err
	}
	dstMessages, err := unmarshal(id, destFile)
	if err != nil {
		return err
	}

	result := *dstMessages
	for key, value := range *srcMessages {
		if _, ok := result[key]; !ok {
			result[key] = value
		}
	}

	var content []byte
	of := strings.ToLower(destFile)
	if strings.HasSuffix(of, ".json") {
		content, err = marshal(result, "json")
	}
	if strings.HasSuffix(of, ".toml") {
		content, err = marshal(result, "toml")
	}
	if strings.HasSuffix(of, ".yaml") {
		content, err = marshal(result, "yaml")
	}
	if err != nil {
		return err
	}

	err = os.MkdirAll(path.Dir(destFile), os.ModePerm)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destFile, content, os.ModePerm)
	if err != nil {
		return nil
	}

	fmt.Printf("Update %+v ...\n", destFile)

	return nil
}
