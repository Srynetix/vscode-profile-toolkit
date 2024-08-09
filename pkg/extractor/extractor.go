package extractor

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Srynetix/vscode-profile-toolkit/pkg/models"
)

type ProfilePackExtractor struct{}

func (p *ProfilePackExtractor) Extract(pack *models.ProfilePack, destination string) {
	if !pathExists(destination) {
		fmt.Fprintf(os.Stderr, "Destination folder \"%s\" does not exist.", destination)
		os.Exit(1)
	}

	profilePath := filepath.Join(destination, pack.Name)
	err := os.MkdirAll(profilePath, 0755)
	if err != nil {
		panic(err)
	}

	p.writeProfileFile(pack, destination)
	p.writeSettingsFile(pack, destination)
	p.writeKeybindingsFile(pack, destination)
	p.writeSnippetsDirectory(pack, destination)
	p.writeExtensionsFile(pack, destination)
	p.writeGlobalStateFile(pack, destination)
}

func (p *ProfilePackExtractor) writeProfileFile(pack *models.ProfilePack, destination string) {
	profilePath := filepath.Join(destination, pack.Name)

	// Create profile file
	profileData := map[string]any{
		"name": pack.Name,
		"icon": pack.Icon,
	}
	profileDataSerialized, err := json.MarshalIndent(profileData, "", "\t")
	if err != nil {
		panic(err)
	}
	err = os.WriteFile(filepath.Join(profilePath, "profile.jsonc"), profileDataSerialized, 0644)
	if err != nil {
		panic(err)
	}
}

func (p *ProfilePackExtractor) writeSettingsFile(pack *models.ProfilePack, destination string) {
	profilePath := filepath.Join(destination, pack.Name)

	err := os.WriteFile(filepath.Join(profilePath, "settings.jsonc"), []byte(pack.Settings.Text), 0644)
	if err != nil {
		panic(err)
	}
}

func (p *ProfilePackExtractor) writeKeybindingsFile(pack *models.ProfilePack, destination string) {
	profilePath := filepath.Join(destination, pack.Name)

	err := os.WriteFile(filepath.Join(profilePath, "keybindings.jsonc"), []byte(pack.Keybindings.Text), 0644)
	if err != nil {
		panic(err)
	}
}

func (p *ProfilePackExtractor) writeSnippetsDirectory(pack *models.ProfilePack, destination string) {
	profilePath := filepath.Join(destination, pack.Name)
	var err error

	snippetsPath := filepath.Join(profilePath, "snippets")
	err = os.MkdirAll(snippetsPath, 0755)
	if err != nil {
		panic(err)
	}

	for lang, snippet := range pack.Snippets {
		langJsonc := lang + "c"
		err = os.WriteFile(filepath.Join(snippetsPath, langJsonc), []byte(snippet.Text), 0644)
		if err != nil {
			panic(err)
		}
	}
}

func (p *ProfilePackExtractor) writeExtensionsFile(pack *models.ProfilePack, destination string) {
	profilePath := filepath.Join(destination, pack.Name)

	extensionsDataSerialized, err := json.MarshalIndent(pack.Extensions, "", "\t")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(profilePath, "extensions.jsonc"), extensionsDataSerialized, 0644)
	if err != nil {
		panic(err)
	}
}

func (p *ProfilePackExtractor) writeGlobalStateFile(pack *models.ProfilePack, destination string) {
	profilePath := filepath.Join(destination, pack.Name)

	serialized, err := json.MarshalIndent(pack.GlobalState, "", "\t")
	if err != nil {
		panic(err)
	}

	err = os.WriteFile(filepath.Join(profilePath, "globalState.jsonc"), serialized, 0644)
	if err != nil {
		panic(err)
	}
}

func pathExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return !os.IsNotExist(err)
	}

	return true
}
