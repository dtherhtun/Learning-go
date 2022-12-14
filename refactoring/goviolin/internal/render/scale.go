package render

import "strings"

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
			{"Pitch", "Major", false, false, "Major"},
			{"Pitch", "Minor", false, true, "Minor"},
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
			{"Octave", "1", false, true, "1 Octave"},
			{"Octave", "2", false, false, "2 Octave"},
		}
	case "2":
		options = []Option{
			{"Octave", "1", false, false, "1 Octave"},
			{"Octave", "2", false, true, "2 Octave"},
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
			{Name: "Key", Value: "A", IsDisabled: false, IsChecked: true, Text: "A"},
			{Name: "Key", Value: "Bb", IsDisabled: false, IsChecked: false, Text: "Bb"},
			{Name: "Key", Value: "B", IsDisabled: false, IsChecked: false, Text: "B"},
			{Name: "Key", Value: "C", IsDisabled: false, IsChecked: false, Text: "C"},
			{Name: "Key", Value: "C#/Db", IsDisabled: false, IsChecked: false, Text: "C#/Db"},
			{Name: "Key", Value: "D", IsDisabled: false, IsChecked: false, Text: "D"},
			{Name: "Key", Value: "Eb", IsDisabled: false, IsChecked: false, Text: "Eb"},
			{Name: "Key", Value: "E", IsDisabled: false, IsChecked: false, Text: "E"},
			{Name: "Key", Value: "F", IsDisabled: false, IsChecked: false, Text: "F"},
			{Name: "Key", Value: "F#/Gb", IsDisabled: false, IsChecked: false, Text: "F#/Gb"},
			{Name: "Key", Value: "G", IsDisabled: false, IsChecked: false, Text: "G"},
			{Name: "Key", Value: "G#/Ab", IsDisabled: false, IsChecked: false, Text: "G#/Ab"},
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

// ChangeSharpToS WE DON'T KNOW WHY YET.
func ChangeSharpToS(path string) string {
	if strings.Contains(path, "#") {
		path = path[:len(path)-1]
		path += "s"
	}

	return path
}
