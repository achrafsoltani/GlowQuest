package audio

import "math"

const sampleRate = 44100

// GenerateSwordSwing creates a short descending sweep (woosh).
func GenerateSwordSwing() []byte {
	duration := 0.1
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		freq := 600.0 - 300.0*progress
		val := math.Sin(2 * math.Pi * freq * t)

		env := 1.0 - progress
		env *= env

		sample := int16(val * env * 6000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GenerateEnemyHit creates a short impact burst.
func GenerateEnemyHit() []byte {
	duration := 0.08
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	lfsr := uint16(0xBEEF)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		tone := math.Sin(2*math.Pi*500*t) * 0.6

		bit := ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 5)) & 1
		lfsr = (lfsr >> 1) | (bit << 15)
		noise := float64(int16(lfsr)) / 32768.0 * 0.4

		env := 1.0 - progress
		env *= env

		sample := int16((tone + noise) * env * 7000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GenerateEnemyDie creates a noise crunch with descending pitch.
func GenerateEnemyDie() []byte {
	duration := 0.3
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	lfsr := uint16(0xACE1)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		freq := 400.0 - 200.0*progress
		tone := math.Sin(2*math.Pi*freq*t) * 0.4

		bit := ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 5)) & 1
		lfsr = (lfsr >> 1) | (bit << 15)
		noise := float64(int16(lfsr)) / 32768.0 * 0.6

		env := 1.0 - progress
		env = env * env * env

		sample := int16((tone + noise) * env * 8000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GeneratePlayerHit creates a low thud with noise.
func GeneratePlayerHit() []byte {
	duration := 0.15
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	lfsr := uint16(0xDEAD)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		thud := math.Sin(2*math.Pi*200*t) * 0.7

		bit := ((lfsr >> 0) ^ (lfsr >> 2) ^ (lfsr >> 3) ^ (lfsr >> 5)) & 1
		lfsr = (lfsr >> 1) | (bit << 15)
		noise := float64(int16(lfsr)) / 32768.0 * 0.3

		env := 1.0 - progress
		env *= env

		sample := int16((thud + noise) * env * 8000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GenerateItemPickup creates an ascending arpeggio chime.
func GenerateItemPickup() []byte {
	duration := 0.25
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	notes := []float64{659.25, 783.99, 880.0} // E5, G5, A5
	noteLen := samples / len(notes)

	for i := 0; i < samples; i++ {
		noteIdx := i / noteLen
		if noteIdx >= len(notes) {
			noteIdx = len(notes) - 1
		}
		freq := notes[noteIdx]
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		val := math.Sin(2*math.Pi*freq*t)*0.7 + math.Sin(2*math.Pi*freq*2*t)*0.2

		env := 1.0 - progress*0.5
		sample := int16(val * env * 5000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GenerateDoorOpen creates a 3-note ascending sequence.
func GenerateDoorOpen() []byte {
	duration := 0.3
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	notes := []float64{261.63, 329.63, 392.0} // C4, E4, G4
	noteLen := samples / len(notes)

	for i := 0; i < samples; i++ {
		noteIdx := i / noteLen
		if noteIdx >= len(notes) {
			noteIdx = len(notes) - 1
		}
		freq := notes[noteIdx]
		t := float64(i) / float64(sampleRate)
		localT := float64(i%noteLen) / float64(noteLen)

		val := math.Sin(2 * math.Pi * freq * t)

		env := 1.0
		if localT > 0.8 {
			env = (1.0 - localT) * 5.0
		}

		sample := int16(val * env * 5000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GenerateMenuSelect creates a short blip.
func GenerateMenuSelect() []byte {
	duration := 0.05
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		val := math.Sin(2 * math.Pi * 1000 * t)
		env := 1.0 - progress
		sample := int16(val * env * 5000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}

// GenerateGameOver creates a descending sad tone.
func GenerateGameOver() []byte {
	duration := 1.0
	samples := int(float64(sampleRate) * duration)
	buf := make([]byte, samples*2)

	for i := 0; i < samples; i++ {
		t := float64(i) / float64(sampleRate)
		progress := float64(i) / float64(samples)

		freq := 400.0 - 200.0*progress
		val := math.Sin(2*math.Pi*freq*t)*0.6 + math.Sin(2*math.Pi*freq*0.5*t)*0.3

		env := 1.0 - progress
		sample := int16(val * env * 8000)
		buf[i*2] = byte(sample)
		buf[i*2+1] = byte(sample >> 8)
	}
	return buf
}
