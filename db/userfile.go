package db

import (
	"fmt"
	"time"
)

// UserFile : 用户文件表结构体
type UserFile struct {
	UserName   string
	FileHash   string `gorm:"colum:file_sha1"`
	FileName   string
	FileSize   int64
	UploadAt   time.Time
	LastUpdate string
}

func (u UserFile) TableName() string {
	return "tbl_user_file"
}

// OnUserFileUploadFinished : 更新用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	err := DB.Model(&UserFile{}).Where("UserName=?").
		Update("FileName", filename, "FileSize", filesize, username, "FileHash=?", filehash).Error
	if err != nil {
		return false
	}
	return true
}

// QueryUserFileMetas : 批量获取用户文件信息
func QueryUserFileMetas(username string, limit int) ([]UserFile, error) {
	var userFiles []UserFile
	err := DB.Where("UserName=?", username).Limit(limit).Find(userFiles).Error
	if err != nil {
		return nil, err
	}
	return userFiles, nil
}

// DeleteUserFile : 删除文件(标记删除)
func DeleteUserFile(username, filehash string) bool {
	err := DB.Where("UserName=?", username, "FileHash=?", filehash).Delete(UserFile{}).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// RenameFileName : 文件重命名
func RenameFileName(username, filehash, filename string) bool {
	err := DB.Model(&UserFile{}).Where("UserName=?", username, "FileHash=?", filehash).
		Update("FileName", filename).Error
	if err != nil {
		return false
	}
	return true
}

// QueryUserFileMeta : 获取用户单个文件信息
func QueryUserFileMeta(username string, filehash string) (*UserFile, error) {
	var userFile UserFile
	err := DB.Where("UserName=?", username, "FileHash=?", filehash).Find(userFile).Error
	if err != nil {
		return nil, err
	}
	return &userFile, nil
}
