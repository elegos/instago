package processor

import (
	"fmt"
	"image"
	_ "image/gif" // Enable gif decoding
	"image/jpeg"
	_ "image/png" // Enable png decoding
	"math"
	"os"

	"github.com/elegos/instago/source"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

type instaSize struct {
	maxHeight   uint
	maxWidth    uint
	aspectRatio float64
}

type instaSizeScore struct {
	size  instaSize
	score float64
}

var sizes = []instaSize{
	{maxHeight: 1080, maxWidth: 1080, aspectRatio: 1},            // square 1:1
	{maxHeight: 608, maxWidth: 1080, aspectRatio: 1.91},          // landscape 1.91:1
	{maxHeight: 1350, maxWidth: 608, aspectRatio: 0.8},           // portrait 4:5
	{maxHeight: 1920, maxWidth: 1080, aspectRatio: 0.5625},       // stories 9:16
	{maxHeight: 654, maxWidth: 420, aspectRatio: 0.645161290322}, // IGTV cover 1:1.55
}

// ProcessImage convert an image to .insta.jpg
func ProcessImage(src string) {
	if _, err := os.Stat(src); err != nil {
		logrus.Errorf("File '%s' does not exist", src)
		os.Exit(2)
	}

	srcReader, err := os.Open(src)
	if err != nil {
		logrus.WithError(err).Error("Impossible to open the source image")
		os.Exit(source.ExitCodeCantOpenFile)
	}
	defer srcReader.Close()

	srcImg, _, err := image.Decode(srcReader)
	if err != nil {
		logrus.WithError(err).Error("Impossible to read the image's data")
		os.Exit(source.ExitCodeCantParseImage)
	}

	bounds := srcImg.Bounds()
	srcWidth := uint(bounds.Max.X - bounds.Min.X)
	srcHeight := uint(bounds.Max.Y - bounds.Min.Y)
	srcRatio := float64(srcWidth / srcHeight)

	// Get the image's aspect ratio
	ratio := float64(srcWidth / srcHeight)
	// Get the most similar ratio from the instagram image sizes list
	scores := []instaSizeScore{}
	for _, size := range sizes {
		scores = append(scores, instaSizeScore{size: size, score: math.Abs(ratio - size.aspectRatio)})
	}

	// Best guess on the aspect ratio to get maxWidth and maxHeight from
	targetScore := scores[0]
	for _, score := range scores {
		if targetScore.score > score.score {
			targetScore = score
		}
	}

	targetSize := targetScore.size

	destImg := srcImg
	destWidth := srcWidth
	destHeight := srcHeight
	destRatio := srcRatio

	if targetSize.maxWidth < destWidth {
		destWidth = targetSize.maxWidth
		destHeight = uint(math.Round(float64(destWidth) / destRatio))
		destImg = resize.Resize(destWidth, destHeight, destImg, resize.Lanczos3)
	}

	if targetSize.maxHeight < destHeight {
		destHeight = targetSize.maxHeight
		destWidth = uint(math.Round(float64(destHeight) * destRatio))
		destImg = resize.Resize(destWidth, destHeight, destImg, resize.Lanczos3)
	}

	outFilePath := fmt.Sprintf("%s.insta.jpg", src)
	outFile, err := os.Create(outFilePath)
	if err != nil {
		logrus.WithError(err).Error("Impossible to create the output image")
		os.Exit(source.ExitCodeCantCreateImage)
	}
	defer outFile.Close()

	jpeg.Encode(outFile, destImg, nil)

	logrus.WithFields(logrus.Fields{
		"input file":  src,
		"output file": outFilePath,
	}).Info("The image was succesfully converted")
}
