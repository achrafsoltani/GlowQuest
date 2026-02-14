package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/glow"
)

var (
	ColorDialogueBG     = glow.RGB(15, 15, 40)
	ColorDialogueBorder = glow.RGB(200, 200, 200)
	ColorDialogueText   = glow.RGB(255, 255, 255)
	ColorDialogueName   = glow.RGB(200, 200, 50)
	ColorDialogueArrow  = glow.RGB(200, 200, 200)
)

// DrawDialogueBox renders the dialogue box at the bottom of the play area.
func DrawDialogueBox(sc *ScaledCanvas, name string, lines []string, currentLine int, hasMore bool) {
	boxH := config.DialogueBoxH
	boxY := config.HUDHeight + config.PlayAreaHeight - boxH
	boxW := config.PlayAreaWidth

	// Dark background
	sc.DrawRect(0, boxY, boxW, boxH, ColorDialogueBG)
	// Border
	sc.DrawRectOutline(0, boxY, boxW, boxH, ColorDialogueBorder)

	// Name
	DrawText(sc, name, 4, boxY+3, ColorDialogueName)

	// Show up to 3 lines of dialogue starting at currentLine
	textY := boxY + 12
	for i := 0; i < 3 && currentLine+i < len(lines); i++ {
		DrawText(sc, lines[currentLine+i], 4, textY+i*charSpaceY, ColorDialogueText)
	}

	// Arrow indicator if there are more lines
	if hasMore {
		arrowX := boxW - 10
		arrowY := boxY + boxH - 8
		sc.SetPixel(arrowX+1, arrowY, ColorDialogueArrow)
		sc.DrawRect(arrowX, arrowY+1, 3, 1, ColorDialogueArrow)
		sc.SetPixel(arrowX+1, arrowY+2, ColorDialogueArrow)
	}
}
