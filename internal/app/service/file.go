package service

import (
	"errors"
	"image"

	"github.com/disintegration/imaging"
	"github.com/gogf/gf/os/gfile"
)

var File = &fileSrv{
	collectionName: "file",
}

type fileSrv struct {
	collectionName string
}

// CollectionName .
func (srv *fileSrv) CollectionName() string {
	return srv.collectionName
}

// GenCropImage 生成裁剪图片
func (srv *fileSrv) GenCropImage(src, mode, dst string, width, height int) error {
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
	imgFormat, _ := srv.GetCropImageFormat(gfile.ExtName(src))
	return imaging.Encode(cropImgFile, img, imgFormat)
}

// GetCropImageFormat 获取图片类型
func (srv *fileSrv) GetCropImageFormat(fileExtName string) (imaging.Format, bool) {
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
func (srv *fileSrv) CanCrop(fileExtName string) bool {
	if _, ok := srv.GetCropImageFormat(fileExtName); ok {
		return true
	}
	return false
}
