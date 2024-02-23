package models

type ConversionMessage struct {
	UserId       string `json:"user_id"`
	UserName     string `json:"user_name"`
	FileName     string `json:"file_name"`
	FilePath     string `json:"file_path"`
	OutputFormat string `json:"output_format"`
} //This message the convert handler will make

type AfterConvertUpload struct {
	UserId   string `json:"user_id"`
	UserName string `json:"user_name"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
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
