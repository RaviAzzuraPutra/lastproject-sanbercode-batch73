package helper

import (
	"context"
	"fmt"
	"last-project/app/config/cloudinary_config"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadFotoToCloudinary(filePath string, folderName string, fileName string) (string, error) {
	cld, err := cloudinary.NewFromParams(
		cloudinary_config.Cloudinary_Config().CLOUDINARY_CLOUD_NAME,
		cloudinary_config.Cloudinary_Config().CLOUDINARY_API_KEY,
		cloudinary_config.Cloudinary_Config().CLOUDINARY_API_SECRET,
	)

	fullFolderPath := fmt.Sprintf("sanbercode/lastproject/%s", folderName)

	if err != nil {
		NewInternalServerError("An error occurred while cloudinary. " + err.Error())
	}

	uploadResult, errUpload := cld.Upload.Upload(
		context.Background(),
		filePath,
		uploader.UploadParams{
			Folder:   fullFolderPath,
			PublicID: fileName,
		},
	)

	if errUpload != nil {
		NewInternalServerError("An error occurred while uploading the image. " + errUpload.Error())
	}

	return uploadResult.SecureURL, nil
}
