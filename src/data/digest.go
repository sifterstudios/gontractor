package data

import (
	"encoding/json"
	"errors"
	"os"

	util "github.com/sifterstudios/gontractor/src/util"
)

func WriteDataToFile(weeks map[string]Week, goal Goal) error {
	jsonData, err := json.MarshalIndent(FileData{weeks, goal}, util.JsonIndent, util.JsonPrefix)
	if err != nil {
		return errors.New("could not parse data file")
	}

	err = os.WriteFile(util.DataFilePath, jsonData, 0644)

	if err != nil {
		return errors.New("could not write to data file")
	}

	return nil
}
