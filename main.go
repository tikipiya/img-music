import (
	"fmt"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func main() {
	filePath := "input.png"
	imgFile, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("画像が開けない : %v\n", err)
		return
	}
	defer imgFile.Close()

	img, _, err := image.Decode(imgFile)
	if err != nil {
		fmt.Printf("画像のでコードに失敗 : %v\n", err)
		return
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y
	clock := smf.MetricTicks(96)
	s := smf.New()
	s.TimeFormat = clock
	var tr smf.Track
	tr.Add(0, smf.MetaInstrument("Sonification Piano"))
	stepX := width / 50
	if stepX == 0 {
		stepX = 1
	}
	for x := 0; x < width; x += stepX {
		var totalLuminance float64
		var totalSaturation float64
		count := 0
		for y := 0; y < height; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			rf, gf, bf := float64(r>>8), float64(g>>8), float64(b>>8)
			lum := 0.299*rf + 0.587*gf + 0.114*bf
			totalLuminance += lum
			max := math.Max(rf, math.Max(gf, bf))
			min := math.Min(rf, math.Min(gf, bf))
			if max > 0 {
				totalSaturation += (max - min) / max
			}
			count++
		}
		avgLum := totalLuminance / float64(count)
		avgSat := totalSaturation / float64(count)
		pitch := uint8(36 + (avgLum/255.0)*60)
		velocity := uint8(40 + (avgSat * 80))
		duration := uint32(96)
		tr.Add(0, midi.NoteOn(0, pitch, velocity))
		tr.Add(duration, midi.NoteOff(0, pitch))
	}
	tr.Add(0, smf.MetaEndOfTrack())
	s.Add(tr)
	outPath := "output.mid"
	err = s.WriteFile(outPath)
	if err != nil {
		fmt.Printf("MIDIの保存に失敗 : %v\n", err)
		return
	}
	fmt.Printf("成功: %s を生成しました（画像サイズ : %dx%d）\n", outPath, width, height)
}