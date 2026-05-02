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

	}
}