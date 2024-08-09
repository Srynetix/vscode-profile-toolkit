package archiver

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/Srynetix/vscode-profile-toolkit/pkg/models"
)

type ProfilePackArchiver struct{}

func (p *ProfilePackArchiver) ArchiveTo(pack *models.ProfilePack, destination string) {
	outputData := map[string]any{}
	outputData["name"] = pack.Name

	if pack.Icon != nil {
		outputData["icon"] = pack.Icon
	}

	// Settings
	if pack.Settings != nil {
		settingsData := map[string]any{}
		settingsData["settings"] = pack.Settings.Text
		settingsStr, err := json.Marshal(settingsData)
		if err != nil {
			panic(err)
		}
		outputData["settings"] = string(settingsStr)
	}

	// Keybindings
	if pack.Keybindings != nil {
		keybindingsData := map[string]any{}
		keybindingsData["keybindings"] = pack.Keybindings.Text
		keybindingsStr, err := json.Marshal(keybindingsData)
		if err != nil {
			panic(err)
		}
		outputData["keybindings"] = string(keybindingsStr)
	}

	// Snippets
	if pack.Snippets != nil {
		snippetsData := map[string]map[string]any{}
		snippetsData["snippets"] = map[string]any{}
		for key, value := range *pack.Snippets {
			snippetsData["snippets"][key] = value.Text
		}
		snippetsStr, err := json.Marshal(snippetsData)
		if err != nil {
			panic(err)
		}
		outputData["snippets"] = string(snippetsStr)
	}

	// Extensions
	if pack.Extensions != nil {
		extensionsStr, err := json.Marshal(pack.Extensions)
		if err != nil {
			panic(err)
		}
		outputData["extensions"] = string(extensionsStr)
	}

	// Global state
	if pack.GlobalState != nil {
		globalStateStr, err := json.Marshal(pack.GlobalState)
		if err != nil {
			panic(err)
		}
		outputData["globalState"] = string(globalStateStr)
	}

	// Export everything
	outputStr, err := json.Marshal(outputData)
	if err != nil {
		panic(err)
	}

	if destination == "-" {
		fmt.Print(string(outputStr))
	} else {
		err = os.WriteFile(destination, outputStr, 0644)
		if err != nil {
			panic(err)
		}
	}
}
