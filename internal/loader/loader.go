package loader

import (
	"encoding/json"
	"errors"
	"fmt"
	"image/color"
	"os"
	"os/user"
	"path"

	"github.com/hultan/planet-simulator/internal/data"
)

// au : Constant for the astronomical unit. Needed to correct the JSON-data
const au = 149.6e9

// Loader : Loads a JSON file containing the solar system
type Loader struct {
}

const defaultDataPath = "code/planet simulator/data/solar.json"

// NewLoader : Create a new loader
func NewLoader() *Loader {
	return new(Loader)
}

// Load : Loads the solar system from JSON
func (l *Loader) Load() (solar *data.SolarSystem, err error) {
	// Get the path to the Loader file
	loaderPath := l.getLoaderPath()

	// Open Loader file
	LoaderFile, err := os.Open(loaderPath)

	// Handle errors
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	defer func() {
		err = LoaderFile.Close()
	}()

	solar = &data.SolarSystem{}

	// Parse the JSON document
	jsonParser := json.NewDecoder(LoaderFile)
	err = jsonParser.Decode(solar)
	if err != nil {
		return nil, err
	}

	l.fixColors(solar)
	l.fixDistances(solar)

	return solar, nil
}

//
// Private functions
//

// Get current users home directory
func (l *Loader) getHomeDirectory() string {
	u, err := user.Current()
	if err != nil {
		errorMessage := fmt.Sprintf("Failed to get user home directory : %s", err)
		panic(errorMessage)
	}
	return u.HomeDir
}

// getLoaderPath : Get path to the Loader file
func (l *Loader) getLoaderPath() string {
	home := l.getHomeDirectory()

	return path.Join(home, defaultDataPath)
}

// fixColors : Converts HEX colors to RGB colors
func (l *Loader) fixColors(solar *data.SolarSystem) {
	for i := range solar.Bodies {
		c, err := parseHexColor(solar.Bodies[i].Color)
		if err != nil {
			panic(err)
		}
		solar.Bodies[i].ColorObj = c
	}
}

// fixDistances : The JSON contains data in AU, application needs it in meters
func (l *Loader) fixDistances(solar *data.SolarSystem) {
	for i := range solar.Bodies {
		body := solar.Bodies[i]
		body.Position = body.Position.Mul(au)
	}
}

var errInvalidFormat = errors.New("invalid format")

// parseHexColor : Parses colors in HEX format and returns an RGB color
// https://stackoverflow.com/questions/54197913/parse-hex-string-to-image-color
func parseHexColor(s string) (c color.RGBA, err error) {
	c.A = 0xff

	if s[0] != '#' {
		return c, errInvalidFormat
	}

	hexToByte := func(b byte) byte {
		switch {
		case b >= '0' && b <= '9':
			return b - '0'
		case b >= 'a' && b <= 'f':
			return b - 'a' + 10
		case b >= 'A' && b <= 'F':
			return b - 'A' + 10
		}
		err = errInvalidFormat
		return 0
	}

	switch len(s) {
	case 7:
		c.R = hexToByte(s[1])<<4 + hexToByte(s[2])
		c.G = hexToByte(s[3])<<4 + hexToByte(s[4])
		c.B = hexToByte(s[5])<<4 + hexToByte(s[6])
	case 4:
		c.R = hexToByte(s[1]) * 17
		c.G = hexToByte(s[2]) * 17
		c.B = hexToByte(s[3]) * 17
	default:
		err = errInvalidFormat
	}
	return
}
