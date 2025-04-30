package main

import "os"

type HudElement struct {
	// The name of the element
	name             string
	oldX, oldY       float32
	newX, newY       float32
	OffsetX, OffsetY int64
}

func newHudE(name string, x, y int64) *HudElement {
	return &HudElement{
		name:    name,
		OffsetX: x,
		OffsetY: y,
	}

}

func (h *HudElement) apply(file *os.File) {
	if h.OffsetX == 0 && h.OffsetY == 0 {
		panic("at least one Offset must be set")
	}
	if h.newX != 0 {
		if h.OffsetX != 0 {
			writeFloat32BytesToFile(file, h.OffsetX, h.newX)
		}
	}
	if h.newY != 0 {
		if h.OffsetY != 0 {
			writeFloat32BytesToFile(file, h.OffsetY, h.newY)
		}
	}

}

// trying to right align the elements
func (h *HudElement) setDefaultScaledValues(targetWidth, targetHeight float32) {
	if h.newX > 0 || h.newY > 0 {
		return
	}
	originalWidth, originalHeight := float32(800), float32(600)

	marginRight := originalWidth - h.oldX
	h.newX = targetWidth - marginRight

	h.newY = (h.oldY * targetHeight) / originalHeight
}

func getHudValues(width float32) []*HudElement {
	hud := []*HudElement{}

	hp := newHudE("HP", 0x2A4CD4, 0x2A4CD8)
	hp.oldX = 505
	hp.oldY = 565

	ChargeBarNum := newHudE("ChargeBarNum", 0x2A4C9C, 0x2A4CA0)
	ChargeBarNum.oldX = 676
	ChargeBarNum.oldY = 489

	ChargeBar := newHudE("ChargeBar", 0x2A4CAC, 0x2A4CC0)
	ChargeBar.oldX = 597
	ChargeBar.oldY = 517

	ChargeBarCharge := newHudE("ChargeBarCharge", 0, 0x2A4C98)
	ChargeBarCharge.oldY = 525

	MainBarNumber := newHudE("MainBarNumber", 0x2A4CA4, 0x2A4CA8)
	MainBarNumber.oldX = 702
	MainBarNumber.oldY = 443

	MainBar := newHudE("MainBar", 0x2A4CB0, 0x2A4CB4)
	MainBar.oldX = 623
	MainBar.oldY = 471

	CPUMode := newHudE("CPUMode", 0, 0x2A4CBC)
	CPUMode.oldY = 520

	MaxModeIcon := newHudE("MaxModeIcon", 0, 0x2A4CCC)
	MaxModeIcon.oldY = 534

	OyashiroText := newHudE("OyashiroText", 0, 0x2A4CC8)
	OyashiroText.oldY = 548

	OyashiroBar := newHudE("OyashiroBar", 0, 0x2A4CD0)
	OyashiroBar.oldY = 488

	BalloonRight := newHudE("BalloonRight", 0x2A4C8C, 0)
	BalloonRight.oldX = 715 // -50 works well in all resolutions
	//BalloonRightx := newHud(0x2A4C8C, 715.0-50, true) // -50 works well in all resolutions
	BalloonRightPoints := newHudE("BalloonRightPoints", 0x2A4C88, 0)
	BalloonRightPoints.oldX = 740 // -65 works well in all resolutions
	//BalloonRightPointsx := newHud(0x2A4C88, 740.0-50-15, true) // -65 works well in all resolutions?

	// Values that need to be set manually

	limit := newHudE("Limit", 0x2A4C84, 0)
	limit.newX = width - 35

	limitLabel := newHudE("LimitLabel", 0x2A1C80, 0)
	limitLabel.newX = width - 100

	charSelectExit := newHudE("CharSelectExit", 0xCAD6, 0xCAD1)
	charSelectExit.newX = width/2 + 70
	charSelectExit.newY = width / 2

	charSelectGo := newHudE("CharSelectGo", 0xCABA, 0xCAB5)
	charSelectGo.newX = width/2 - 60
	charSelectGo.newY = width / 2

	charSelectCursor := newHudE("CharSelectCursor", 0xB670, 0)
	charSelectCursor.newX = width / 2

	charSelectWide := newHudE("CharSelectWide", 0x8B4F, 0)
	charSelectWide.newX = width / 8

	charSelectPortrait := newHudE("CharSelectPortrait", 0x2A16D8, 0)
	charSelectPortrait.newX = width / 3.6

	hud = append(hud,
		hp,
		ChargeBarNum,
		ChargeBar,
		ChargeBarCharge,
		MainBarNumber,
		MainBar,
		CPUMode,
		MaxModeIcon,
		OyashiroText,
		OyashiroBar,
		BalloonRight,
		BalloonRightPoints,
		limit, limitLabel,
		charSelectExit,
		charSelectGo,
		charSelectCursor,
		charSelectWide,
		charSelectPortrait,
	)

	return hud
}
