package event

type {{.Name}} struct {
	// define event fields here
}

func (e {{.Name}}) Name() string {
	return "{{.Name | lower}}"
}

func (e {{.Name}}) Description() string {
	return "{{.Description}}"
}
