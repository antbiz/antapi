package api

import (
	"fmt"
	"strings"

	"github.com/BeanWei/apikit/errors"
	"github.com/BeanWei/apikit/resp"
	"github.com/antbiz/antapi/internal/app/dao"
	"github.com/antbiz/antapi/internal/app/service"
	"github.com/antbiz/antapi/internal/common/errmsg"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/text/gstr"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/guid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// File 文件上传下载接口
var File = &fileApi{}

type fileApi struct{}

// Upload 上传文件
func (fileApi) Upload(r *ghttp.Request) {
	file := r.GetUploadFile("file")
	if file == nil {

	}
	oriFileName := file.Filename
	file.Filename = fmt.Sprintf("%s.%s", guid.S([]byte(r.RemoteAddr), []byte(fmt.Sprintf("%v", r.Header))), gfile.ExtName(oriFileName))

	savePathNodes := g.SliceStr{
		gfile.Pwd(),
		"upload",
		fmt.Sprintf("%d", gtime.Now().Year()),
		fmt.Sprintf("%d", gtime.Now().Month()),
		fmt.Sprintf("%d", gtime.Now().Day()),
	}
	savePath := gfile.Join(savePathNodes...)
	if _, err := file.Save(savePath); err != nil {
		resp.Error(r, errors.InternalServer(errmsg.ErrFileUpload, g.I18n().T(errmsg.ErrFileUpload)).WithOrigErr(err))
	}

	opt := &dao.InsertOptions{
		IncludeHiddenField:  true,
		IncludePrivateField: true,
	}
	data := g.Map{
		"url":  "/" + strings.Join(savePathNodes[1:], "/"),
		"name": oriFileName,
		"size": file.Size,
	}
	if id, err := dao.Insert(r.Context(), service.File.CollectionName(), data, opt); err != nil {
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrFileUpload)).WithOrigErr(err))
	} else {
		var thumbURL string
		if gregex.IsMatchString(`\.(gif|png|jpg|jpeg|webp)`, oriFileName) {
			// 默认封面图按照自动填充100*100裁剪
			thumbURL = fmt.Sprintf("%s?ir=fill_100_100", data["url"])
		}
		data["id"] = id
		data["thumbURL"] = thumbURL
		resp.OK(r, data)
	}
}

// Preview 文件预览
func (fileApi) Preview(r *ghttp.Request) {
	opt := &dao.GetOptions{
		Filter:             bson.M{"url": r.RequestURI},
		IncludeHiddenField: true,
		IgnoreFieldsCheck:  true,
	}
	data, err := dao.Get(r.Context(), service.File.CollectionName(), opt)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			resp.Error(r, errors.NotFound(errmsg.ErrDBNotFound, g.I18n().T(errmsg.ErrDBNotFound)))
		}
		resp.Error(r, errors.DatabaseError(g.I18n().T(errmsg.ErrDBGet)))
	}

	filepath := gfile.Join(gstr.SplitAndTrimSpace(fmt.Sprintf("%s/%s", gfile.Pwd(), r.RequestURI), "/")...)
	fileExt := gfile.ExtName(r.RequestURI)
	// 图片裁剪参数:  ir=mode_width_height
	// mode: [fit,fill,fixed]指定图片缩放的策略，有三种策略，fit 表示固定一边，另一边按比例缩放；
	// 		fill表示先将图片延伸出指定W与H的矩形框外，然后进行居中裁剪；
	//		fixed表示直接按照指定的W和H缩放图片，这种方式可能导致图片变形
	// width: 1 ~ 4096 指定的宽度，0表示自动适应
	// height: 1 ~ 4096 指定的宽度，0表示自动适应
	ir := gstr.ToLower(r.GetQueryString("ir"))
	if ir != "" && service.File.CanCrop(fileExt) {
		cropImgPath := fmt.Sprintf("%s/%s_%s.%s", gfile.Dir(filepath), gfile.Name(filepath), ir, fileExt)
		if !gfile.Exists(cropImgPath) {
			irArgs := strings.Split(ir, "_")
			if len(irArgs) != 3 {
				resp.Error(r, errors.InvalidArgument("ir param error. the format is mode_w_h"))
			}
			imgResizeMode := irArgs[0]
			if imgResizeMode == "" {
				imgResizeMode = "fit"
			} else if imgResizeMode != "fit" && imgResizeMode != "fill" && imgResizeMode != "fixed" {
				resp.Error(r, errors.InvalidArgument("mode can only be fit/fill/fixed"))
			}
			imgResizeWidth := gconv.Int(irArgs[1])
			if irArgs[1] != "" && (imgResizeWidth < 0 || imgResizeWidth > 4096) {
				resp.Error(r, errors.InvalidArgument("zoom size cannot exceed 4096"))
			}
			imgResizeHeight := gconv.Int(irArgs[2])
			if irArgs[2] != "" && (imgResizeHeight < 0 || imgResizeHeight > 4096) {
				resp.Error(r, errors.InvalidArgument("zoom size cannot exceed 4096"))
			}
			if imgResizeMode == "fit" && imgResizeWidth == 0 && imgResizeHeight == 0 {
				resp.Error(r, errors.InvalidArgument("mode fit required width or height"))
			} else if (imgResizeMode == "fill" || imgResizeMode == "fixed") && (imgResizeWidth == 0 || imgResizeHeight == 0) {
				resp.Error(r, errors.InvalidArgument("mode fill/fixed required width and height"))
			}
			if err := service.File.GenCropImage(filepath, imgResizeMode, cropImgPath, imgResizeWidth, imgResizeHeight); err != nil {
				resp.Error(r, errors.InvalidArgument("file not support preview"))
			}
		}
		filepath = cropImgPath
	}
	if gstr.ToLower(r.GetQueryString("action")) == "download" {
		r.Response.Header().Set("Content-Type", "application/octet-stream")
		r.Response.Header().Set("Accept-Ranges", "bytes")
		r.Response.Header().Set("Content-type", "application/force-download")
		r.Response.Header().Set("Content-Disposition", fmt.Sprintf("attachment;filename=%s", data.GetString("name")))
	}
	r.Response.ServeFile(filepath)
}
