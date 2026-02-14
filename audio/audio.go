package audio

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/AchrafSoltani/glow"
)

type Engine struct {
	ctx *glow.AudioContext

	swordSwingBuf []byte
	enemyHitBuf   []byte
	enemyDieBuf   []byte
	playerHitBuf  []byte
	itemPickupBuf []byte
	doorOpenBuf   []byte
	menuSelectBuf []byte
	gameOverBuf   []byte

	Muted  bool
	Volume float64
}

func NewEngine() *Engine {
	ctx, err := glow.NewAudioContext(sampleRate, 1, 2)
	if err != nil {
		log.Printf("audio: failed to init: %v", err)
		return &Engine{Volume: 1.0}
	}

	return &Engine{
		ctx:           ctx,
		swordSwingBuf: GenerateSwordSwing(),
		enemyHitBuf:   GenerateEnemyHit(),
		enemyDieBuf:   GenerateEnemyDie(),
		playerHitBuf:  GeneratePlayerHit(),
		itemPickupBuf: GenerateItemPickup(),
		doorOpenBuf:   GenerateDoorOpen(),
		menuSelectBuf: GenerateMenuSelect(),
		gameOverBuf:   GenerateGameOver(),
		Volume:        1.0,
	}
}

func (e *Engine) play(buf []byte) {
	if e.ctx == nil || len(buf) == 0 || e.Muted {
		return
	}

	scaled := buf
	if e.Volume < 1.0 {
		scaled = make([]byte, len(buf))
		for i := 0; i+1 < len(buf); i += 2 {
			sample := int16(binary.LittleEndian.Uint16(buf[i:]))
			sample = int16(float64(sample) * e.Volume)
			binary.LittleEndian.PutUint16(scaled[i:], uint16(sample))
		}
	}

	p := e.ctx.NewPlayer(bytes.NewReader(scaled))
	p.Play()
}

func (e *Engine) ToggleMute()  { e.Muted = !e.Muted }
func (e *Engine) VolumeUp()    { e.Volume += 0.1; if e.Volume > 1.0 { e.Volume = 1.0 } }
func (e *Engine) VolumeDown()  { e.Volume -= 0.1; if e.Volume < 0.0 { e.Volume = 0.0 } }

func (e *Engine) PlaySwordSwing() { e.play(e.swordSwingBuf) }
func (e *Engine) PlayEnemyHit()   { e.play(e.enemyHitBuf) }
func (e *Engine) PlayEnemyDie()   { e.play(e.enemyDieBuf) }
func (e *Engine) PlayPlayerHit()  { e.play(e.playerHitBuf) }
func (e *Engine) PlayItemPickup() { e.play(e.itemPickupBuf) }
func (e *Engine) PlayDoorOpen()   { e.play(e.doorOpenBuf) }
func (e *Engine) PlayMenuSelect() { e.play(e.menuSelectBuf) }
func (e *Engine) PlayGameOver()   { e.play(e.gameOverBuf) }
