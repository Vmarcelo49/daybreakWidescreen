package main

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/guigui/basicwidget"

	"github.com/hajimehoshi/guigui"
)

type Layout struct {
	guigui.DefaultWidget

	//configDatPanel    basicwidget.Panel
	configText        basicwidget.Text
	charOutlineToggle basicwidget.Toggle
	shadowToggle      basicwidget.Toggle
	HiTextureToggle   basicwidget.Toggle

	//patchLabel         basicwidget.Panel
	resolutionDropdown basicwidget.DropdownList[string]
	applyButton        basicwidget.Button
	revertButton       basicwidget.Button
}

func (l *Layout) Build(context *guigui.Context, appender *guigui.ChildWidgetAppender) error {
	l.configText.SetScale(2)
	l.configText.SetBold(true)
	l.configText.SetValue("config.dat")

	l.charOutlineToggle.SetValue(true)
	l.shadowToggle.SetValue(true)
	l.HiTextureToggle.SetValue(true)

	placeHolderResolutionList := []string{
		"640x480",
		"800x600",
		"1024x768",
	}
	l.resolutionDropdown.SetItemsByStrings(placeHolderResolutionList)
	l.resolutionDropdown.SelectItemByIndex(0)

	l.applyButton.SetText("Apply")
	l.applyButton.SetOnUp(func() {
		fmt.Println("Apply button clicked")
		// Here you would typically apply the settings
	})
	l.revertButton.SetText("Revert")
	l.revertButton.SetOnUp(func() {
		fmt.Println("Revert button clicked")
		// Here you would typically revert the settings
	})

	return nil
}

func runGuiGui() {
	op := &guigui.RunOptions{
		Title:          "Daybreak Widescreen",
		WindowMinSize:  image.Pt(320, 240),
		RunGameOptions: &ebiten.RunGameOptions{},
	}
	if err := guigui.Run(&Layout{}, op); err != nil {
		fmt.Fprintln(os.Stderr, "Error running GUI:", err)
		os.Exit(1)
	}
}
