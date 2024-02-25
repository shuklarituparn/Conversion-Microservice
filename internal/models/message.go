package models

type ConversionMessage struct {
	UserId       int    `json:"user_id"`
	UserName     string `json:"user_name"`
	FileName     string `json:"file_name"`
	FilePath     string `json:"file_path"`
	OutputFormat string `json:"output_format"`
	VideoKey     string `json:"video_key"`
} //This message the convert handler will make

type AfterConvertUpload struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	VideoKey string `json:"video_key"`
}

type EmailVerificationMessage struct {
	UserID           int    `json:"user_id"`
	UserName         string `json:"user_name"`
	UserEmail        string `json:"user_email"`
	VerificationCode string `json:"verification_code"`
}
type RestoreAccountMessage struct {
	UserId   int    `json:"userId"`
	UserName string `json:"userName"`
}
type MailSendMessage struct {
	Filepath string `json:"filepath"`
	TO       string `json:"TO"`
	Subject  string `json:"subject"`
}
