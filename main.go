package main

import (
	"flag"
	"os"

	"slices"
)

var resolutions = []string{"1152x648", "1280x720", "1366x768", "1600x900", "1920x1080", "2560x1440", "3840x2160"}

var IsCLIMode bool

func main() {
	cliRevert := flag.Bool("revert", false, "revert to original DaybreakDX.exe from ./backup/")
	cliShadow := flag.Bool("shadow", true, "by using this, turns off shadows")
	cliOutline := flag.Bool("outline", true, "by using this, turns off character outline")
	cliHiTexture := flag.Bool("htexture", true, "by using this, turns off high texture")
	cliName := flag.String("name", "", "Set network name (max 8 characters)")
	cliResolution := flag.String("resolution", "", "Set resolution (ex: 1920x1080)")

	flag.Parse()
	if flag.NFlag() > 0 || flag.NArg() > 0 {
		IsCLIMode = true
	}

	verifyRequiredFiles()
	if !IsCLIMode {
		startWindow()
	} else {
		if *cliRevert {
			revertToOriginalEXE()
			return
		}
		if *cliResolution != "" {
			resOk := slices.Contains(resolutions, *cliResolution)
			if !resOk {
				throwErrorMessageWindow("Invalid resolution")
				listAvaliableResolutions()
				os.Exit(1)
			}
			patchAndSave(*cliResolution)
		}
		if *cliName != "" {
			if len(*cliName) > 8 || !isAlphanumeric(*cliName) {
				throwErrorMessageWindow("Invalid name: must be max 8 characters and alphanumeric")
			} else {
				setNetworkName(*cliName)
			}
		}
		if *cliShadow {
			err := setBoolConfig(shadows, true)
			if err != nil {
				throwErrorMessageWindow("Error while setting shadows" + err.Error())
			}
		}
		if *cliOutline {
			err := setBoolConfig(outline, true)
			if err != nil {
				throwErrorMessageWindow("Error while setting outline" + err.Error())
			}
		}
		if *cliHiTexture {
			err := setBoolConfig(higerResTex, true)
			if err != nil {
				throwErrorMessageWindow("Error while setting high texture" + err.Error())
			}
		}
	}

}
