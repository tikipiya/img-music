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

}