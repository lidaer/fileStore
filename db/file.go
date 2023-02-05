package db

import (
	"fmt"
)

// TableFile : 文件表结构体
type TableFile struct {
	ID       int64
	FileHash string `gorm:"column:file_sha1"`
	FileName string `gorm:"colum:file_name"`
	FileSize int64  `gorm:"colum:file_size"`
	FileAddr string `gorm:"colum:file_addr"`
	Status   int    `gorm:"colum:status"`
}

func (t TableFile) TableName() string {
	return "tbl_file"
}

// OnFileUploadFinished : 文件上传完成，保存meta
func OnFileUploadFinished(filehash string, filename string,
	filesize int64, fileaddr string) bool {

	file := TableFile{FileHash: filehash, FileName: filename, FileSize: filesize, FileAddr: fileaddr, Status: 1}
	err := DB.Create(&file).Error
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

// GetFileMeta : 从mysql获取文件元信息
func GetFileMeta(filehash string) (*TableFile, error) {
	tableFile := TableFile{FileHash: filehash}
	err := DB.Find(&tableFile).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	return &tableFile, err
}

// GetFileMetaList : 从mysql批量获取文件元信息
func GetFileMetaList(limit int) ([]TableFile, error) {
	var tfiles []TableFile
	err := DB.Find(&tfiles).Error
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	fmt.Println(len(tfiles))
	return tfiles, nil

}

// UpdateFileLocation : 更新文件的存储地址(如文件被转移了)
func UpdateFileLocation(filehash string, fileaddr string) bool {
	err := DB.Model(&TableFile{}).Where("FileHash = ?", filehash).Update("FileAddr", fileaddr).Error
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
