package media

import (
	"context"
	"course-api/pkg/response"
	"course-api/pkg/utils"
	"errors"
	"mime/multipart"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
	"github.com/google/uuid"
)

type Media interface {
	Upload(file multipart.FileHeader) (*string, *response.Error)
	Delete(file string) (*string, *response.Error)
}

type mediaUsecase struct {
}

// Delete implements Media.
func (usecase *mediaUsecase) Delete(file string) (*string, *response.Error) {
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_API_KEY") + ":" + os.Getenv("CLOUDINARY_API_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUD_NAME"))
	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	filename := utils.GetFileName(file)

	res, err := cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: filename,
	})
	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	return &res.Result, nil
}

// Upload implements Media.
func (usecase *mediaUsecase) Upload(file multipart.FileHeader) (*string, *response.Error) {
	cld, err := cloudinary.NewFromURL("cloudinary://" + os.Getenv("CLOUDINARY_API_KEY") + ":" + os.Getenv("CLOUDINARY_API_SECRET") + "@" + os.Getenv("CLOUDINARY_CLOUD_NAME"))
	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}

	var ctx = context.Background()

	binary, err := file.Open()
	if err != nil {
		return nil, &response.Error{
			Code: 500,
			Err:  err,
		}
	}
	defer binary.Close()

	if binary != nil {
		uploadResult, err := cld.Upload.Upload(
			ctx,
			binary,
			uploader.UploadParams{
				PublicID: uuid.New().String(),
			})
		if err != nil {
			return nil, &response.Error{
				Code: 500,
				Err:  err,
			}
		}

		return &uploadResult.SecureURL, nil
	}
	return nil, &response.Error{
		Code: 500,
		Err:  errors.New("cannot read file"),
	}
}

func NewMediaUsecase() Media {
	return &mediaUsecase{}
}
