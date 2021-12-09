package main

type Template struct {
	SystemsIds []string
	ChassisIds []string
}

func (template *Template) setDefaultTemplate() {
	// Computer System Collection
	template.SystemsIds = make([]string, 0)
	template.SystemsIds = append(template.SystemsIds, "Newell")

	// Chassis Collection
	template.ChassisIds = make([]string, 0)
	template.ChassisIds = append(template.ChassisIds, "1")
}
