package logic

import (
	"errors"
	"image"

	"github.com/disintegration/imaging"
	"github.com/gogf/gf/os/gfile"
)

var File = fileLogic{}

type fileLogic struct{}

// GenCropImage 生成裁剪图片
func (file *fileLogic) GenCropImage(src, mode, dst string, width, height int) error {
	srcFile, err := gfile.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcImg, err := imaging.Decode(srcFile)
	if err != nil {
		return err
	}
	var img *image.NRGBA
	switch mode {
	case "fit":
		img = imaging.Fit(srcImg, width, height, imaging.Lanczos)
	case "fill":
		img = imaging.Fill(srcImg, width, height, imaging.Center, imaging.Lanczos)
	case "fixed":
		img = imaging.Resize(srcImg, width, height, imaging.Lanczos)
	default:
		return errors.New("mode can only be fit/fill/fixed")
	}
	cropImgFile, err := gfile.Create(dst)
	if err != nil {
		return err
	}
	defer cropImgFile.Close()
	imgFormat, _ := file.GetCropImageFormat(gfile.ExtName(src))
	return imaging.Encode(cropImgFile, img, imgFormat)
}

// GetCropImageFormat
func (file *fileLogic) GetCropImageFormat(fileExtName string) (imaging.Format, bool) {
	formats := map[string]imaging.Format{
		".jpg":  imaging.JPEG,
		"jpg":   imaging.JPEG,
		".jpeg": imaging.JPEG,
		"jpeg":  imaging.JPEG,
		".png":  imaging.PNG,
		"png":   imaging.PNG,
		".tif":  imaging.TIFF,
		"tif":   imaging.TIFF,
		".tiff": imaging.TIFF,
		"tiff":  imaging.TIFF,
		".bmp":  imaging.BMP,
		"bmp":   imaging.BMP,
		".gif":  imaging.GIF,
		"gif":   imaging.GIF,
	}
	imgFormat, ok := formats[fileExtName]
	return imgFormat, ok
}

// CanCrop 是否支持裁剪
func (file *fileLogic) CanCrop(fileExtName string) bool {
	if _, ok := file.GetCropImageFormat(fileExtName); ok {
		return true
	}
	return false
}
