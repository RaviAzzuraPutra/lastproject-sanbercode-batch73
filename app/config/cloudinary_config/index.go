package cloudinary_config

import "os"

type CloudinaryConfig struct {
	CLOUDINARY_CLOUD_NAME string
	CLOUDINARY_API_KEY    string
	CLOUDINARY_API_SECRET string
}

func Cloudinary_Config() *CloudinaryConfig {
	return &CloudinaryConfig{
		CLOUDINARY_CLOUD_NAME: os.Getenv("CLOUDINARY_CLOUD_NAME"),
		CLOUDINARY_API_KEY:    os.Getenv("CLOUDINARY_API_KEY"),
		CLOUDINARY_API_SECRET: os.Getenv("CLOUDINARY_API_SECRET"),
	}
}
