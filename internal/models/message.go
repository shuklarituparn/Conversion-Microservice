package models

type ConversionMessage struct {
	UserId      string `json:"user_id"`
	UserName    string `json:"user_name"`
	UserPicture string `json:"user_picture"`
	FileName    string `json:"file_name"`
	FilePath    string `json:"file_path"`
}
