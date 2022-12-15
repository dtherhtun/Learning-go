package render

import "strings"

// Option represents the options for generating content.
type Option struct {
	Name       string
	Value      string
	IsDisabled bool
	IsChecked  bool
	Text       string
}

// PageVars represents the input for generating a web page.
type PageVars struct {
	Title         string
	Scalearp      string
	Key           string
	Pitch         string
	DuetImgPath   string
	ScaleImgPath  string
	GifPath       string
	AudioPath     string
	AudioPath2    string
	DuetAudioBoth string
	DuetAudio1    string
	DuetAudio2    string
	LeftLabel     string
	RightLabel    string
	Scales        []Option
	Duets         []Option
	Pitches       []Option
	Keys          []Option
	Octaves       []Option
}

// SetDefaultOptions provide the defaults for rendering scales.
func SetDefaultOptions() ([]Option, []Option, []Option, []Option) {

	scale := []Option{
		{"Scale", "Scale", false, true, "Scales"},
		{"Scale", "Arpeggio", false, false, "Arpeggios"},
	}

	pitch := []Option{
		{"Pitch", "Major", false, true, "Major"},
		{"Pitch", "Minor", false, false, "Minor"},
	}

	key := []Option{
		{"Key", "A", false, true, "A"},
		{"Key", "Bb", false, false, "Bb"},
		{"Key", "B", false, false, "B"},
		{"Key", "C", false, false, "C"},
		{"Key", "C#/Db", false, false, "C#/Db"},
		{"Key", "D", false, false, "D"},
		{"Key", "Eb", false, false, "Eb"},
		{"Key", "E", false, false, "E"},
		{"Key", "F", false, false, "F"},
		{"Key", "F#/Gb", false, false, "F#/Gb"},
		{"Key", "G", false, false, "G"},
		{"Key", "G#/Ab", false, false, "G#/Ab"},
	}

	octave := []Option{
		{"Octave", "1", false, true, "1 Octave"},
		{"Octave", "2", false, false, "2 Octave"},
	}

	return scale, pitch, key, octave
}

// SetScaleOptions sets the scale options base on the specified scale.
func SetScaleOptions(scale string) []Option {
	var options []Option

	switch scale {
	case "Scale":
		options = []Option{
			{
				Name:       "Scale",
				Value:      "Scale",
				IsDisabled: false,
				IsChecked:  true,
				Text:       "Scales",
			},
			{
				Name:       "Scale",
				Value:      "Arpeggio",
				IsDisabled: false,
				IsChecked:  false,
				Text:       "Arpeggios",
			},
		}
	case "Arpeggio":
		options = []Option{
			{
				Name:       "Scale",
				Value:      "Scale",
				IsDisabled: false,
				IsChecked:  false,
				Text:       "Scales",
			},
			{
				Name:       "Scale",
				Value:      "Arpeggio",
				IsDisabled: false,
				IsChecked:  true,
				Text:       "Arpeggios",
			},
		}
	}

	return options
}

// SetPitchOptions sets the pitch options base on the specified pitch.
func SetPitchOptions(pitch string) []Option {
	var options []Option

	switch pitch {
	case "Major":
		options = []Option{
			{Name: "Pitch", Value: "Major", IsDisabled: false, IsChecked: true, Text: "Major"},
			{Name: "Pitch", Value: "Minor", IsDisabled: false, IsChecked: false, Text: "Minor"},
		}
	case "Minor":
		options = []Option{
			{Name: "Pitch", Value: "Major", IsDisabled: false, IsChecked: false, Text: "Major"},
			{Name: "Pitch", Value: "Minor", IsDisabled: false, IsChecked: true, Text: "Minor"},
		}
	}

	return options
}

// SetOctaveOptions sets the pitch options base on the specified pitch.
func SetOctaveOptions(octave string) []Option {
	var options []Option

	switch octave {
	case "1":
		options = []Option{
			{Name: "Octave", Value: "1", IsDisabled: false, IsChecked: true, Text: "1 Octave"},
			{Name: "Octave", Value: "2", IsDisabled: false, IsChecked: false, Text: "2 Octave"},
		}
	case "2":
		options = []Option{
			{Name: "Octave", Value: "1", IsDisabled: false, IsChecked: false, Text: "1 Octave"},
			{Name: "Octave", Value: "2", IsDisabled: false, IsChecked: true, Text: "2 Octave"},
		}
	}

	return options
}

// SetKeyOptions sets the key options based on specified key.
func SetKeyOptions(key string) []Option {
	var options []Option
	switch key {
	case "A":
		options = []Option{
			{"Key", "A", false, true, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Bb":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, true, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "B":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, true, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, true, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "C#/Db":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, true, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "D":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, true, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "Eb":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, true, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "E":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, true, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, true, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "F#/Gb":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, true, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}
	case "G":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, true, "G"},
			{"Key", "G#/Ab", false, false, "G#/Ab"},
		}

	case "G#/Ab":
		options = []Option{
			{"Key", "A", false, false, "A"},
			{"Key", "Bb", false, false, "Bb"},
			{"Key", "B", false, false, "B"},
			{"Key", "C", false, false, "C"},
			{"Key", "C#/Db", false, false, "C#/Db"},
			{"Key", "D", false, false, "D"},
			{"Key", "Eb", false, false, "Eb"},
			{"Key", "E", false, false, "E"},
			{"Key", "F", false, false, "F"},
			{"Key", "F#/Gb", false, false, "F#/Gb"},
			{"Key", "G", false, false, "G"},
			{"Key", "G#/Ab", false, true, "G#/Ab"},
		}
	}

	return options
}

// SetActualKey transforms the provided key based on the selected pitch.
// For major scales if the key is longer than 2 characters, we only care about
// the last 2 characters
// For minor scales if the key is longer than 2 characters, we only care about
// the first 2 characters
func SetActualKey(pitch, key string) string {
	switch pitch {
	case "Major":

		// only select last two characters for keys which contain two possible
		// names e.g. C#/Db.
		if len(key) > 2 {
			key = key[3:]
		}
	case "Minor":

		// only select first two characters for keys which contain two possible
		// names e.g. C#/Db.
		if len(key) > 2 {
			key = key[:2]
		}
	}

	return key
}

// SetMusicLabels sets the text for the music players. Set the labels, Major have a scale and
// a drone, while minor have melodic and harmonic minor scales
func SetMusicLabels(pitch, scale string) (string, string) {
	var left, right string

	switch pitch {
	case "Major":

		left = "Listen to Major "
		right = "Listen to Drone"
		if scale == "Scale" {
			left += "Scale"
			break
		}
		left += "Arpeggio"

	case "Minor":

		switch scale {
		case "Scale":
			left = "Listen to Harmonic Minor Scale"
			right = "Listen to Melodic Minor Scale"
		case "Arpeggio":
			left = "Listen to Minor Arpeggio"
			right = "Listen to Drone"
		}
	}

	return left, right
}

// SetAssetPaths build paths to img and mp3 files that correspond to user selection.
func SetAssetPaths(pitch, scale, key, octave string) (string, string, string) {

	imgPath, audioPath, audioPath2 := "img/", "mp3/", "mp3/"
	switch scale {
	case "Scale":
		imgPath += "scale/"
		audioPath += "scale/"
	case "Arpeggio":
		imgPath += "arps/"
		audioPath += "arps/"
	}

	switch pitch {
	case "Major":
		imgPath += "major/"
		audioPath += "major/"
	case "Minor":
		imgPath += "minor/"
		audioPath += "minor/"
	}
	audioPath += strings.ToLower(key)
	imgPath += strings.ToLower(key)

	imgPath = changeSharpToS(imgPath)
	audioPath = changeSharpToS(audioPath)

	switch octave {
	case "1":
		imgPath += "1"
		audioPath += "1"
	case "2":
		imgPath += "2"
		audioPath += "2"
	}

	audioPath += ".mp3"
	imgPath += ".png"

	// Audio path2 can either be a melodic minor scale or a drone note.
	// Set to melodic minor scale - if the first 16 characters of audio path.
	switch {
	case audioPath[:16] == "mp3/scale/minor/":
		// Set audioPath2 to the original audioPath.
		audioPath2 = audioPath

		// Chop off the last 4 characters, this removes .mp3.
		audioPath2 = audioPath2[:len(audioPath2)-4]

		// Then add m for melodic and the .mp3 suffix
		audioPath2 += "m.mp3"
	default:
		audioPath2 += "drone/"
		audioPath2 += strings.ToLower(key)

		// May have just added a # to the path, so use the function to
		// change # to s.
		audioPath2 = changeSharpToS(audioPath2)

		switch octave {
		case "1":
			audioPath2 += "1.mp3"
		case "2":
			audioPath2 += "2.mp3"
		}
		if audioPath[:16] == "mp3/scale/minor/" {

		} else { // audioPath2 needs to be a drone note.

		}
	}

	return imgPath, audioPath, audioPath2
}

// changeSharpToS WE DON'T KNOW WHY YET.
func changeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}

	return path
}
