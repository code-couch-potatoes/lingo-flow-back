package skyeng

type Meaning struct {
	// Meaning id
	ID string `json:"id"`
	// Word is a group of meanings. Skyeng combines meanings by word entity
	WordID int `json:"wordId"`
	// There are 6 difficulty levels: 1, 2, 3, 4, 5, 6
	DifficultyLevel int `json:"difficultyLevel"`
	// String representation of a part of speech.
	//
	// Available codes:
	// n - noun,
	// v - verb,
	// j - adjective,
	// r - adverb,
	// prp - preposition,
	// prn - pronoun,
	// crd - cardinal number,
	// cjc - conjunction,
	// exc - interjection,
	// det - article,
	// abb - abbreviation,
	// x - particle,
	// ord - ordinal number,
	// md - modal verb,
	// ph - phrase,
	// phi - idiom.
	PartOfSpeechCode string `json:"partOfSpeechCode"`
	// Infinitive particle (to) or articles (a, the)
	Prefix string `json:"prefix"`
	// Meaning text
	Text string `json:"text"`
	// URL to meaning sound
	SoundURL string `json:"soundUrl"`
	// IPA phonetic transcription
	Transcription string      `json:"transcription"`
	UpdatedAt     string      `json:"updatedAt"`
	Mnemonics     string      `json:"mnemonics"`
	Translation   Translation `json:"translation"`
	// A collection of an images
	Images     []Image    `json:"images"`
	Definition Definition `json:"definition"`
	// Usage examples
	Examples []Example `json:"examples"`
	// Collection of meanings with similar translations
	MeaningsWithSimilarTranslation []MeaningWithSimilarTranslation `json:"meaningsWithSimilarTranslation"`
	// Collection of alternative translations
	AlternativeTranslations []AlternativeTranslation `json:"alternativeTranslations"`
}

type Image struct {
	URL string `json:"url"`
}

type Definition struct {
	Text     string `json:"text"`
	SoundURL string `json:"soundUrl"`
}

type Example struct {
	Text     string `json:"text"`
	SoundURL string `json:"soundUrl"`
}

type MeaningWithSimilarTranslation struct {
	MeaningID int `json:"meaningId"`
	// Percentage of frequency among all other usages
	FrequencyPercent         string      `json:"frequencyPercent"`
	PartOfSpeechAbbreviation string      `json:"partOfSpeechAbbreviation"`
	Translation              Translation `json:"translation"`
}

type AlternativeTranslation struct {
	// A text of a meaning
	Text        string      `json:"text"`
	Translation Translation `json:"translation"`
}

type Word struct {
	ID       int            `json:"id"`
	Text     string         `json:"text"`
	Meanings []ShortMeaning `json:"meanings"`
}

type ShortMeaning struct {
	ID int `json:"id"`
	// String representation of a part of speech. Same as Meaning.PartOfSpeechCode
	PartOfSpeechCode string      `json:"partOfSpeechCode"`
	Transcription    string      `json:"transcription"`
	Translation      Translation `json:"translation"`

	PreviewURL string `json:"previewUrl"`
	ImageURL   string `json:"imageUrl"`
	SoundURL   string `json:"soundUrl"`
}

type Translation struct {
	// A text of a translation
	Text string `json:"text"`
	// A note about translation
	Note string `json:"note"`
}
