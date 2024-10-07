package config

// InputFieldConfig defines the configuration for a form field
type InputFieldConfig struct {
	Label       string
	Placeholder string
	ID          string
	Type        string
}

// FieldConfigManager defines a struct that holds both the slice and map of field configs
type FieldConfigManager struct {
	Configs   []InputFieldConfig
	ConfigMap map[string]InputFieldConfig
}

// NewFieldConfigManager initializes the FieldConfigManager with the provided configs
func NewFieldConfigManager(configs []InputFieldConfig) *FieldConfigManager {
	configMap := make(map[string]InputFieldConfig)
	for _, config := range configs {
		configMap[config.ID] = config
	}

	return &FieldConfigManager{
		Configs:   configs,
		ConfigMap: configMap,
	}
}

// ProjectFieldConfigManager holds the configurations for the project form
var ProjectFieldConfigManager = NewFieldConfigManager([]InputFieldConfig{
	{Label: "Project Name", Placeholder: "Enter project name", ID: "name", Type: "text"},
	{Label: "Description", Placeholder: "Enter project description", ID: "description", Type: "text"},
})
