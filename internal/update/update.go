package update

import (
	"encoding/json"
	"github.com/radding/retroPieManager/internal/update/settings"
	"os"
	"path"
)

// IfNeeded Checks if the settings updated and returns the Settings and a bool if they were updated
func IfNeeded() (settings.Settings, bool) {
	didUpdate := false
	home, ok := os.LookupEnv("HOME")
	if !ok {
		panic("Home is not defined")
	}
	fileLocation := path.Join(home, ".retropie-manager.json")
	if _, err := os.Stat(fileLocation); os.IsNotExist(err) {
		file, err := os.Create(fileLocation)
		if err != nil {
			panic(err)
		}
		file.Write([]byte(`{}`))
		file.Close()
		didUpdate = true
	}
	newSettings, err := settings.FromEnvironment()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(fileLocation)
	if err != nil {
		panic(err)
	}
	var data []byte
	file.Read(data)
	sets := settings.Settings{}
	if err := json.Unmarshal(); err != nil {
		panic(err)
	}

	return newSettings, didUpdate
}
