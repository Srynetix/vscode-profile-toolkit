package models

type ProfilePack struct {
	Name        string                          `json:"name"`
	Icon        *string                         `json:"icon"`
	Settings    *ProfilePackSettings            `json:"settings"`
	Keybindings *ProfilePackKeybindings         `json:"keybindings"`
	Snippets    *map[string]ProfilePackSnippets `json:"snippets"`
	Extensions  *[]ProfilePackExtension         `json:"extensions"`
	GlobalState *ProfilePackGlobalState         `json:"globalState"`
}

type ProfilePackSettings struct {
	Text   string
	Parsed map[string]any
}

type ProfilePackKeybindings struct {
	Text   string
	Parsed []map[string]any
}

type ProfilePackSnippets struct {
	Text   string
	Parsed map[string]any
}

type ProfilePackExtension struct {
	Identifier  ProfilePackExtensionIdentifier `json:"identifier"`
	DisplayName string                         `json:"displayName"`
}

type ProfilePackExtensionIdentifier struct {
	Id   string `json:"id"`
	Uuid string `json:"uuid"`
}

type ProfilePackGlobalState struct {
	Storage map[string]any `json:"storage"`
}
