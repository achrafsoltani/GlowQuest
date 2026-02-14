package main

import (
	"flag"
	"image"
	"image/png"
	"log"
	"os"
	"time"

	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/game"
	"github.com/AchrafSoltani/glow"
)

func main() {
	screenshot := flag.String("screenshot", "", "save screenshot to file after rendering and exit")
	flag.Parse()

	win, err := glow.NewWindow("GlowQuest", config.WindowWidth*4, config.WindowHeight*4)
	if err != nil {
		log.Fatal(err)
	}
	defer win.Close()

	g := game.NewGame()
	canvas := win.Canvas()
	running := true
	lastTime := time.Now()
	frameCount := 0

	for running {
		now := time.Now()
		dt := now.Sub(lastTime).Seconds()
		lastTime = now

		if dt > 0.05 {
			dt = 0.05
		}

		for {
			event := win.PollEvent()
			if event == nil {
				break
			}
			switch event.Type {
			case glow.EventQuit:
				running = false
			case glow.EventKeyDown:
				if event.Key == glow.KeyF11 {
					win.SetFullscreen(!win.IsFullscreen())
				}
				// ESC in menu quits; otherwise delegated to game
				if event.Key == glow.KeyEscape && g.State == game.StateMenu {
					g.ShouldQuit = true
				}
				// Audio controls
				if event.Key == glow.KeyM {
					g.ToggleMute()
				}
				if event.Key == glow.KeyEqual {
					g.VolumeUp()
				}
				if event.Key == glow.KeyMinus {
					g.VolumeDown()
				}
				g.KeyDown(event.Key)
			case glow.EventKeyUp:
				g.KeyUp(event.Key)
			case glow.EventWindowResize:
				g.OnResize(event.Width, event.Height)
			}
		}

		if g.ShouldQuit {
			running = false
		}

		g.Update(dt)

		canvas.Clear(glow.Black)
		g.Draw(canvas)
		win.Present()

		frameCount++
		if *screenshot != "" && frameCount >= 5 {
			saveScreenshot(canvas, *screenshot)
			running = false
		}

		elapsed := time.Since(now)
		target := time.Second / 60
		if elapsed < target {
			time.Sleep(target - elapsed)
		}
	}
}

func saveScreenshot(canvas *glow.Canvas, path string) {
	w, h := canvas.Width(), canvas.Height()
	img := image.NewRGBA(image.Rect(0, 0, w, h))
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			c := canvas.GetPixel(x, y)
			off := (y*w + x) * 4
			img.Pix[off+0] = c.R
			img.Pix[off+1] = c.G
			img.Pix[off+2] = c.B
			img.Pix[off+3] = 255
		}
	}
	f, err := os.Create(path)
	if err != nil {
		log.Printf("screenshot: %v", err)
		return
	}
	defer f.Close()
	png.Encode(f, img)
	log.Printf("screenshot saved to %s", path)
}
