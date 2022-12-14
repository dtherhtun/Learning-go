package handlers

import (
	"log"
	"net/http"
	"strings"

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

	// TODO: Validate we even need to run this.
	var sOptions []render.Option
	var pOptions []render.Option
	var oOptions []render.Option

	// sOptions, pOptions, kOptions, oOptions := render.SetDefaultScaleOptions()
	keyOptions := render.SetKeyOptions(key)
	scaleOptions := render.SetScaleOptions(scale)
	pitchOptions := render.SetPitchOptions(pitch)
	octaveOptions := render.SetOctaveOptions(octave)

	// work out what the actual key is and set its value
	if pitch == "Major" {
		// for major scales if the key is longer than 2 characters, we only care about the last 2 characters
		if len(key) > 2 { // only select last two characters for keys which contain two possible names e.g. C#/Db
			key = key[3:]
		}
	} else { // pitch is minor
		// for minor scales if the key is longer than 2 characters, we only care about the first 2 characters
		if len(key) > 2 { // only select first two characters for keys which contain two possible names e.g. C#/Db
			key = key[:2]
		}
	}

	var leftlabel string
	var rightlabel string

	// Set the labels, Major have a scale and a drone, while minor have melodic and harmonic minor scales
	if pitch == "Major" {
		leftlabel = "Listen to Major "
		rightlabel = "Listen to Drone"
		if scalearp == "Scale" {
			leftlabel += "Scale"
		} else {
			leftlabel += "Arpeggio"
		}
	} else {
		if scalearp == "Arpeggio" {
			leftlabel += "Listen to Minor Arpeggio"
			rightlabel = "Listen to Drone"
		} else {
			leftlabel += "Listen to Harmonic Minor Scale"
			rightlabel += "Listen to Melodic Minor Scale"
		}
	}

	// Intialise paths to the associated images and mp3s
	imgPath, audioPath, audioPath2 := "img/", "mp3/", "mp3/"

	// Build paths to img and mp3 files that correspond to user selection
	if scalearp == "Scale" {
		imgPath += "scale/"
		audioPath += "scale/"

	} else {
		// if arpeggio is selected, add "arps/" to the img and mp3 paths
		imgPath += "arps/"
		audioPath += "arps/"
	}

	if pitch == "Major" {
		imgPath += "major/"
		audioPath += "major/"
	} else {
		imgPath += "minor/"
		audioPath += "minor/"
	}

	audioPath += strings.ToLower(key)
	imgPath += strings.ToLower(key)
	// if the img or audio path contain #, delete last character and replace it with s
	imgPath = render.ChangeSharpToS(imgPath)
	audioPath = render.ChangeSharpToS(audioPath)

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

	//generate audioPath2
	// audio path2 can either be a melodic minor scale or a drone note.
	// Set to melodic minor scale - if the first 16 characters of audio path are:
	if audioPath[:16] == "mp3/scale/minor/" {
		audioPath2 = audioPath                      // set audioPath2 to the original audioPath
		audioPath2 = audioPath2[:len(audioPath2)-4] // chop off the last 4 characters, this removes .mp3
		audioPath2 += "m.mp3"                       // then add m for melodic and the .mp3 suffix
	} else { // audioPath2 needs to be a drone note.
		audioPath2 += "drone/"
		audioPath2 += strings.ToLower(key)
		// may have just added a # to the path, so use the function to change # to s
		audioPath2 = render.ChangeSharpToS(audioPath2)
		switch octave {
		case "1":
			audioPath2 += "1.mp3"
		case "2":
			audioPath2 += "2.mp3"
		}
	}

	pv := render.PageVars{
		Title:         "Practice Scales and Arpeggios",
		Scalearp:      scalearp,
		Key:           key,
		Pitch:         pitch,
		ScaleImgPath:  imgPath,
		GifPath:       "img/major/gif/a1.gif",
		AudioPath:     audioPath,
		AudioPath2:    audioPath2,
		LeftLabel:     leftlabel,
		RightLabel:    rightlabel,
		ScaleOptions:  sOptions,
		PitchOptions:  pOptions,
		KeyOptions:    kOptions,
		OctaveOptions: oOptions,
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
		render.ScaleOptions{"Duet", "G Major", false, true, "G Major"},
		render.ScaleOptions{"Duet", "D Major", false, false, "D Major"},
		render.ScaleOptions{"Duet", "A Major", false, false, "A Major"},
	}

	pv := render.PageVars{
		Title:         "Practice Duets",
		Key:           "G Major",
		DuetImgPath:   "img/duet/gmajor.png",
		DuetAudioBoth: "mp3/duet/gmajorduetboth.mp3",
		DuetAudio1:    "mp3/duet/gmajorduetpt1.mp3",
		DuetAudio2:    "mp3/duet/gmajorduetpt2.mp3",
		DuetOptions:   dOptions,
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
	dOptions := []render.ScaleOptions{
		render.ScaleOptions{"Duet", "G Major", false, true, "G Major"},
		render.ScaleOptions{"Duet", "D Major", false, false, "D Major"},
		render.ScaleOptions{"Duet", "A Major", false, false, "A Major"},
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
		dOptions = []render.ScaleOptions{
			render.ScaleOptions{"Duet", "G Major", false, false, "G Major"},
			render.ScaleOptions{"Duet", "D Major", false, true, "D Major"},
			render.ScaleOptions{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/dmajor.png"
		DuetAudioBoth = "mp3/duet/dmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/dmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/dmajorduetpt2.mp3"
	case "G Major":
		dOptions = []render.ScaleOptions{
			render.ScaleOptions{"Duet", "G Major", false, true, "G Major"},
			render.ScaleOptions{"Duet", "D Major", false, false, "D Major"},
			render.ScaleOptions{"Duet", "A Major", false, false, "A Major"},
		}
		DuetImgPath = "img/duet/gmajor.png"
		DuetAudioBoth = "mp3/duet/gmajorduetboth.mp3"
		DuetAudio1 = "mp3/duet/gmajorduetpt1.mp3"
		DuetAudio2 = "mp3/duet/gmajorduetpt2.mp3"

	case "A Major":
		dOptions = []render.ScaleOptions{
			render.ScaleOptions{"Duet", "G Major", false, false, "G Major"},
			render.ScaleOptions{"Duet", "D Major", false, false, "D Major"},
			render.ScaleOptions{"Duet", "A Major", false, true, "A Major"},
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
		DuetOptions:   dOptions,
	}
	if err := render.Render(w, "duets.html", pv); err != nil {
		b.log.Printf("%s %s -> %s : ERROR %v", r.Method, r.URL.Path, r.RemoteAddr, err)
		return
	}
}
