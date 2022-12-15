package handlers

import (
	"log"
	"net/http"

	"github.com/dtherhtun/Learning-go/refactoring/goviolin/internal/render"
)

// Base represent the base handlers.
type Base struct {
	log *log.Logger
}

// Home handler renders the home.html page.
func (b *Base) Home(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)

	pv := render.PageVars{
		Title: "GoViolin",
	}

	if err := render.Render(w, "home.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}

// Scale handles GET calls for the scale page.
func (b *Base) Scale(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)

	scale, pitch, key, octave := render.SetDefaultOptions()

	pv := render.PageVars{
		Title:        "Practice Scales and Arpeggios",
		Scalearp:     "Scale",
		Pitch:        "Major",
		Key:          "A",
		ScaleImgPath: "img/scale/major/a1.png",
		GifPath:      "",
		AudioPath:    "mp3/scale/major/a1.mp3",
		AudioPath2:   "mp3/drone/a1.mp3",
		LeftLabel:    "Listen to Major scale",
		RightLabel:   "Listen to Drone",
		Scales:       scale,
		Pitches:      pitch,
		Keys:         key,
		Octaves:      octave,
	}

	if err := render.Render(w, "scale.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}

// ScaleShow handles POST calls for the scale page
func (b *Base) ScaleShow(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)

	// TODO: write a function to handle errors and missing data
	r.ParseForm()
	pitch := r.Form["Pitch"][0]
	octave := r.Form["Octave"][0]
	scale := r.Form["Scale"][0]
	key := r.Form["Key"][0]

	// sOptions, pOptions, kOptions, oOptions := render.SetDefaultScaleOptions()
	keys := render.SetKeyOptions(key)
	scales := render.SetScaleOptions(scale)
	pitches := render.SetPitchOptions(pitch)
	octaves := render.SetOctaveOptions(octave)

	key = render.SetActualKey(pitch, key)

	leftMusicLabel, rightMusicLabel := render.SetMusicLabels(pitch, scale)
	imgPath, audioPath, audioPath2 := render.SetAssetPaths(pitch, scale, key, octave)

	// if the img or audio path contain #, delete last character and replace it with s

	pv := render.PageVars{
		Title:        "Practice Scales and Arpeggios",
		Scalearp:     scale,
		Key:          key,
		Pitch:        pitch,
		ScaleImgPath: imgPath,
		GifPath:      "img/major/gif/a1.gif",
		AudioPath:    audioPath,
		AudioPath2:   audioPath2,
		LeftLabel:    leftMusicLabel,
		RightLabel:   rightMusicLabel,
		Scales:       scales,
		Pitches:      pitches,
		Keys:         keys,
		Octaves:      octaves,
	}

	if err := render.Render(w, "scale.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}

// Duet handles get calls for the duets page.
func (b *Base) Duet(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)

	// define default duet options
	dOptions := []render.Option{
		{"Duet", "G Major", false, true, "G Major"},
		{"Duet", "D Major", false, false, "D Major"},
		{"Duet", "A Major", false, false, "A Major"},
	}

	pv := render.PageVars{
		Title:         "Practice Duets",
		Key:           "G Major",
		DuetImgPath:   "img/duet/gmajor.png",
		DuetAudioBoth: "mp3/duet/gmajorduetboth.mp3",
		DuetAudio1:    "mp3/duet/gmajorduetpt1.mp3",
		DuetAudio2:    "mp3/duet/gmajorduetpt2.mp3",
		Duets:         dOptions,
	}

	if err := render.Render(w, "duets.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}

// DuetShow handles post calls for the duets page.
func (b *Base) DuetShow(w http.ResponseWriter, r *http.Request) {
	b.log.Printf("%s %s -> %s", r.Method, r.URL.Path, r.RemoteAddr)

	// define default duet options
	dOptions := []render.Option{
		{"Duet", "G Major", false, true, "G Major"},
		{"Duet", "D Major", false, false, "D Major"},
		{"Duet", "A Major", false, false, "A Major"},
	}

	// Set a placeholder image path, this will be changed later.
	DuetImgPath := "img/duet/gmajor.png"
	DuetAudioBoth := "mp3/duet/gmajorduetboth.mp3"
	DuetAudio1 := "mp3/duet/gmajorduetpt1"
	DuetAudio2 := "mp3/duet/gmajorduetpt2"

	r.ParseForm() //r is url.Values which is a map[string][]string
	var dvalues []string
	for _, values := range r.Form { // range over map
		for _, value := range values { // range over []string
			dvalues = append(dvalues, value) // stick each value in a slice I know the name of
		}
	}

	switch dvalues[0] {
	case "D Major":
		dOptions = []render.Option{
			{"Duet", "G Major", false, false, "G Major"},
			{"Duet", "D Major", false, true, "D Major"},
			{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/dmajor.png"
		DuetAudioBoth = "mp3/duet/dmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/dmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/dmajorduetpt2.mp3"
	case "G Major":
		dOptions = []render.Option{
			{"Duet", "G Major", false, true, "G Major"},
			{"Duet", "D Major", false, false, "D Major"},
			{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/gmajor.png"
		DuetAudioBoth = "mp3/duet/gmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/gmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/gmajorduetpt2.mp3"

	case "A Major":
		dOptions = []render.Option{
			{"Duet", "G Major", false, false, "G Major"},
			{"Duet", "D Major", false, false, "D Major"},
			{"Duet", "A Major", false, true, "A Major"},
		}
		DuetImgPath = "img/duet/amajor.png"
		DuetAudioBoth = "mp3/duet/amajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/amajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/amajorduetpt2.mp3"
	}

	//	imgPath := "img/"

	// set default page variables
	pv := render.PageVars{
		Title:         "Practice Duets",
		Key:           "G Major",
		DuetImgPath:   DuetImgPath,
		DuetAudioBoth: DuetAudioBoth,
		DuetAudio1:    DuetAudio1,
		DuetAudio2:    DuetAudio2,
		Duets:         dOptions,
	}
	if err := render.Render(w, "duets.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}
