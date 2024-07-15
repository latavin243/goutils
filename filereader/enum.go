package filereader

type fileType uint32

const (
	fileTypeInvalid fileType = iota
	fileTypeTOML
	fileTypeYAML
	fileTypeJSON
	fileTypeCSV
)

func (t fileType) String() string {
	switch t {
	case fileTypeTOML:
		return "TOML"
	case fileTypeYAML:
		return "YAML"
	case fileTypeJSON:
		return "JSON"
	case fileTypeCSV:
		return "CSV"
	default:
		return "Invalid"
	}
}

var fileExtTypeMap = map[string]fileType{
	".toml": fileTypeTOML,
	".yaml": fileTypeYAML,
	".yml":  fileTypeYAML,
	".json": fileTypeJSON,
	".csv":  fileTypeCSV,
}
