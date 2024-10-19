package main

import (
	"os"
	"os/exec"

	"github.com/gonutz/w32/v2"
	"github.com/gonutz/wui/v2"
)

func throwErrorMessageWindow(message string) {
	w32.MessageBox(0, message, "Error", w32.MB_ICONERROR)
	os.Exit(1)
}

func configWindow(window *wui.Window) {
	windowFont, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	window.SetFont(windowFont)
	window.SetInnerSize(425, 245)
	window.SetTitle("Window")
	window.SetHasMinButton(false)
	window.SetHasMaxButton(false)
	window.SetResizable(false)

	panel1Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel1 := wui.NewPanel()
	panel1.SetFont(panel1Font)
	panel1.SetBounds(10, 10, 200, 175)
	panel1.SetBorderStyle(wui.PanelBorderSingleLine)
	window.Add(panel1)

	networkName, err := getNetworkName()
	if err != nil {
		throwErrorMessageWindow("Error while getting network name" + err.Error())
	}
	editLine1 := wui.NewEditLine()
	editLine1.SetBounds(37, 116, 107, 21)
	editLine1.SetText(networkName)
	panel1.Add(editLine1)

	label1 := wui.NewLabel()
	label1.SetBounds(25, 96, 142, 16)
	label1.SetText("Online Name(max 8 letters)")
	panel1.Add(label1)

	charOutlineBool, err := getBoolConfig(outline)
	if err != nil {
		throwErrorMessageWindow("Error while getting character outline" + err.Error())
	}
	checkBox1 := wui.NewCheckBox()
	checkBox1.SetBounds(25, 25, 109, 17)
	checkBox1.SetText("Character Outline")
	checkBox1.SetChecked(charOutlineBool)
	panel1.Add(checkBox1)

	shadowBool, err := getBoolConfig(shadows)
	if err != nil {
		throwErrorMessageWindow("Error while getting shadows" + err.Error())
	}
	checkBox2 := wui.NewCheckBox()
	checkBox2.SetBounds(25, 50, 100, 17)
	checkBox2.SetText("Enable Shadows")
	checkBox2.SetChecked(shadowBool)
	panel1.Add(checkBox2)

	highTexturesBool, err := getBoolConfig(higerResTex)
	if err != nil {
		throwErrorMessageWindow("Error while getting high textures" + err.Error())
	}
	checkBox3 := wui.NewCheckBox()
	checkBox3.SetBounds(25, 75, 100, 17)
	checkBox3.SetText("High Textures")
	checkBox3.SetChecked(highTexturesBool)
	panel1.Add(checkBox3)

	label4 := wui.NewLabel()
	label4.SetBounds(67, -3, 59, 23)
	label4.SetText("config.dat")
	panel1.Add(label4)

	button1 := wui.NewButton()
	button1.SetBounds(14, 142, 162, 25)
	button1.SetText("Apply Config")
	button1.SetOnClick(func() {
		setBoolConfig(outline, checkBox1.Checked())
		setBoolConfig(shadows, checkBox2.Checked())
		setBoolConfig(higerResTex, checkBox3.Checked())
		setNetworkName(editLine1.Text())

	})
	panel1.Add(button1)

	panel2Font, _ := wui.NewFont(wui.FontDesc{
		Name:   "Tahoma",
		Height: -11,
	})

	panel2 := wui.NewPanel()
	panel2.SetFont(panel2Font)
	panel2.SetBounds(215, 10, 200, 175)
	panel2.SetBorderStyle(wui.PanelBorderSingleLine)
	window.Add(panel2)

	label5 := wui.NewLabel()
	label5.SetBounds(43, -2, 117, 17)
	label5.SetText("Patch DaybreakDX.exe")
	panel2.Add(label5)

	comboBox1 := wui.NewComboBox()
	comboBox1.SetBounds(30, 30, 150, 21)
	comboBox1.SetItems(resolutions)
	comboBox1.SetSelectedIndex(0)
	panel2.Add(comboBox1)

	label6 := wui.NewLabel()
	label6.SetBounds(30, 14, 150, 19)
	label6.SetText("Resolution")
	panel2.Add(label6)

	checkBox5 := wui.NewCheckBox()
	checkBox5.SetBounds(30, 65, 100, 17)
	checkBox5.SetText("Fullscreen")
	panel2.Add(checkBox5)

	button4 := wui.NewButton()
	button4.SetBounds(45, 100, 97, 25)
	button4.SetText("Patch")
	button4.SetOnClick(func() {
		patchAndSave(comboBox1.Text(), checkBox5.Checked())
	})
	panel2.Add(button4)

	button8 := wui.NewButton()
	button8.SetBounds(45, 140, 98, 25)
	button8.SetText("Revert to Original")
	button8.SetOnClick(func() {
		revertToOriginalEXE()
	})
	panel2.Add(button8)

	button5 := wui.NewButton()
	button5.SetBounds(214, 190, 200, 50)
	button5.SetText("Run Game")
	button5.SetOnClick(func() {
		fixxer()
		os.Exit(0)
	})
	window.Add(button5)

	button6 := wui.NewButton()
	button6.SetBounds(10, 190, 200, 50)
	button6.SetText("Edit IP Address")
	button6.SetOnClick(func() {
		cmd := exec.Command("notepad.exe", "addresslist.txt") // test this
		cmd.Run()
	})
	window.Add(button6)
}
