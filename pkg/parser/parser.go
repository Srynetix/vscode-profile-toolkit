package parser

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/Srynetix/vscode-profile-toolkit/pkg/models"
	"github.com/titanous/json5"
)

type ProfilePackParser struct{}

func (p *ProfilePackParser) ParsePath(path string) *models.ProfilePack {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return p.ParseBytes(data)
}

func (p *ProfilePackParser) ParseBytes(bytes []byte) *models.ProfilePack {
	var mapData map[string]any

	err := json.Unmarshal(bytes, &mapData)
	if err != nil {
		panic(err)
	}

	var icon *string
	if mapData["icon"] != nil {
		var stringIcon = mapData["icon"].(string)
		icon = &stringIcon
	}

	return &models.ProfilePack{
		Name:        mapData["name"].(string),
		Icon:        icon,
		Settings:    p.parseSettings(mapData["settings"].(string)),
		Keybindings: p.parseKeybindings(mapData["keybindings"].(string)),
		Snippets:    p.parseSnippets(mapData["snippets"].(string)),
		Extensions:  p.parseExtensions(mapData["extensions"].(string)),
		GlobalState: p.parseGlobalState(mapData["globalState"].(string)),
	}
}

func (p *ProfilePackParser) ParseFolder(path string) *models.ProfilePack {
	profileStr, err := os.ReadFile(filepath.Join(path, "profile.jsonc"))
	if err != nil {
		panic(err)
	}

	var profileData map[string]any
	err = json.Unmarshal([]byte(profileStr), &profileData)
	if err != nil {
		panic(err)
	}

	var profileIcon *string
	profileIconData := profileData["icon"]
	if profileIconData != nil {
		profileIconString := profileIconData.(string)
		profileIcon = &profileIconString
	}

	profileName := profileData["name"].(string)

	// Settings
	settingsStr, err := os.ReadFile(filepath.Join(path, "settings.jsonc"))
	if err != nil {
		panic(err)
	}
	var settingsData = map[string]any{}
	settingsData["settings"] = string(settingsStr)
	settingsStr, err = json.Marshal(settingsData)
	if err != nil {
		panic(err)
	}

	// Keybindings
	keybindingsStr, err := os.ReadFile(filepath.Join(path, "keybindings.jsonc"))
	if err != nil {
		panic(err)
	}
	var keybindingsData = map[string]any{}
	keybindingsData["keybindings"] = string(keybindingsStr)
	keybindingsStr, err = json.Marshal(keybindingsData)
	if err != nil {
		panic(err)
	}

	// Snippets
	snippetsData := map[string]any{}
	snippetFiles, err := os.ReadDir(filepath.Join(path, "snippets"))
	if err != nil {
		panic(err)
	}

	for _, snippetFile := range snippetFiles {
		snippetContent, err := os.ReadFile(filepath.Join(path, "snippets", snippetFile.Name()))
		if err != nil {
			panic(err)
		}

		snippetFileName := snippetFile.Name()
		snippetFileJson := snippetFileName[:len(snippetFileName)-1]
		snippetsData[snippetFileJson] = string(snippetContent)
	}

	var snippetsContainer = map[string]any{}
	snippetsContainer["snippets"] = snippetsData
	snippetsDataStr, err := json.Marshal(snippetsContainer)
	if err != nil {
		panic(err)
	}

	// Extensions
	extensionsStr, err := os.ReadFile(filepath.Join(path, "extensions.jsonc"))
	if err != nil {
		panic(err)
	}

	// Global state
	globalStateStr, err := os.ReadFile(filepath.Join(path, "globalState.jsonc"))
	if err != nil {
		panic(err)
	}

	return &models.ProfilePack{
		Name:        profileName,
		Icon:        profileIcon,
		Settings:    p.parseSettings(string(settingsStr)),
		Keybindings: p.parseKeybindings(string(keybindingsStr)),
		Snippets:    p.parseSnippets(string(snippetsDataStr)),
		Extensions:  p.parseExtensions(string(extensionsStr)),
		GlobalState: p.parseGlobalState(string(globalStateStr)),
	}
}

func (p *ProfilePackParser) parseSettings(data string) models.ProfilePackSettings {
	var settingsData map[string]any
	err := json.Unmarshal([]byte(data), &settingsData)
	if err != nil {
		panic(err)
	}

	settingsInnerString := settingsData["settings"].(string)
	var parsedData map[string]any
	err = json5.Unmarshal([]byte(settingsInnerString), &parsedData)
	if err != nil {
		panic(err)
	}

	return models.ProfilePackSettings{
		Text:   settingsInnerString,
		Parsed: parsedData,
	}
}

func (p *ProfilePackParser) parseKeybindings(data string) models.ProfilePackKeybindings {
	var keybindingsData map[string]any
	err := json.Unmarshal([]byte(data), &keybindingsData)
	if err != nil {
		panic(err)
	}

	keybindingsInnerString := keybindingsData["keybindings"].(string)
	var keybindingsInnerData []map[string]any
	err = json5.Unmarshal([]byte(keybindingsInnerString), &keybindingsInnerData)
	if err != nil {
		panic(err)
	}

	return models.ProfilePackKeybindings{
		Text:   keybindingsInnerString,
		Parsed: keybindingsInnerData,
	}
}

func (p *ProfilePackParser) parseSnippets(data string) map[string]models.ProfilePackSnippets {
	var snippetsData map[string]map[string]string
	err := json.Unmarshal([]byte(data), &snippetsData)
	if err != nil {
		panic(err)
	}

	var outputData map[string]models.ProfilePackSnippets = make(map[string]models.ProfilePackSnippets)
	for key, value := range snippetsData["snippets"] {
		var parsedData map[string]any
		err = json5.Unmarshal([]byte(value), &parsedData)
		if err != nil {
			panic(err)
		}

		outputData[key] = models.ProfilePackSnippets{
			Text:   value,
			Parsed: parsedData,
		}
	}

	return outputData
}

func (p *ProfilePackParser) parseExtensions(data string) []models.ProfilePackExtension {
	var extensionsData []map[string]any
	err := json.Unmarshal([]byte(data), &extensionsData)
	if err != nil {
		panic(err)
	}

	extensions := []models.ProfilePackExtension{}
	for _, extension := range extensionsData {
		identifierMap := extension["identifier"].(map[string]any)

		extensions = append(extensions, models.ProfilePackExtension{
			Identifier: models.ProfilePackExtensionIdentifier{
				Id:   identifierMap["id"].(string),
				Uuid: identifierMap["uuid"].(string),
			},
			DisplayName: extension["displayName"].(string),
		})
	}

	return extensions
}

func (p *ProfilePackParser) parseGlobalState(data string) models.ProfilePackGlobalState {
	var globalStateData map[string]map[string]any
	err := json.Unmarshal([]byte(data), &globalStateData)
	if err != nil {
		panic(err)
	}

	return models.ProfilePackGlobalState{
		Storage: globalStateData["storage"],
	}
}
