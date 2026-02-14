package render

import (
	"github.com/AchrafSoltani/GlowQuest/config"
	"github.com/AchrafSoltani/GlowQuest/entity"
	"github.com/AchrafSoltani/glow"
)

var ColorProjectile = glow.RGB(255, 100, 50)

// DrawProjectile renders a small diamond-shaped projectile.
func DrawProjectile(sc *ScaledCanvas, proj *entity.Projectile) {
	DrawProjectileAt(sc, proj, 0, 0)
}

// DrawProjectileAt renders a projectile with offset.
func DrawProjectileAt(sc *ScaledCanvas, proj *entity.Projectile, offsetX, offsetY int) {
	if proj.Dead {
		return
	}
	px := int(proj.X) + offsetX
	py := int(proj.Y) + config.HUDHeight + offsetY

	// 4×4 diamond
	sc.SetPixel(px+1, py, ColorProjectile)
	sc.DrawRect(px, py+1, 4, 2, ColorProjectile)
	sc.SetPixel(px+1, py+3, ColorProjectile)
}
