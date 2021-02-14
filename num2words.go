package num2words

// N2w hold some function represent its feature
type N2w struct {
	NumberToWord       func(int64) string
	NumberToOrdinal    func(int64) string
	NumberToOrdinalNum func(int64) string
}

var languages = make(map[string]*N2w)

var (
	customNegativeWord = ""
)

// N2wOption is optional option to setting some data, like customNegativeWord, etc.
type N2wOption func()

// CustomNegativeWord is function to set the custom word of negative sign
func CustomNegativeWord(value string) N2wOption {
	return func() {
		customNegativeWord = value
	}
}

// New initiate new converter
func New(lang string, options ...N2wOption) *N2w {
	selectedLang := languages[lang]
	if selectedLang == nil {
		//default is english if not found
		selectedLang = languages["en"]
	}
	for _, opt := range options {
		opt()
	}
	return selectedLang
}
