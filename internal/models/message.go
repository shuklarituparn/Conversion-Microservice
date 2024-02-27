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

type AfterCutUpload struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	VideoKey string `json:"video_key"`
}

type CutMessage struct {
	UserId    int    `json:"user_id"`
	UserName  string `json:"user_name"`
	FileName  string `json:"file_name"`
	FilePath  string `json:"file_path"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
	VideoKey  string `json:"video_key"`
}

type ScreenshotMessage struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	Time     string `json:"start_time"`
	VideoKey string `json:"video_key"`
}

type WatermarkMessage struct {
	UserId        int    `json:"user_id"`
	UserName      string `json:"user_name"`
	FileName      string `json:"file_name"`
	FilePath      string `json:"file_path"`
	WaterMarkFile string `json:"water_mark_file"`
	VideoKey      string `json:"video_key"`
}

type AfterScreenshotUpload struct {
	UserId   int    `json:"user_id"`
	UserName string `json:"user_name"`
	FileName string `json:"file_name"`
	FilePath string `json:"file_path"`
	VideoKey string `json:"video_key"`
}

type FiledownloadMailMessage struct {
	UserName string `json:"user_name"`
	Mode     string `json:"mode"`
	UserID   int    `json:"user_id"`
	FileId   string `json:"file_id"`
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
