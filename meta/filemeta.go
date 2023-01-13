package meta

// FileMeta 文件元信息结构
type FileMeta struct {
	FileSha1 string //文件元Hash值
	FileName string
	FileSize int64
	Location string
	UploadAt string //上传时间
}

var fileMetas map[string]FileMeta

func init() {
	fileMetas = make(map[string]FileMeta)
}

// UpdateFileMeta 新增或者更新文件元信息
func UpdateFileMeta(fmeta FileMeta) {
	fileMetas[fmeta.FileSha1] = fmeta
}

func GetFileMeta(fileSha1 string) FileMeta {
	return fileMetas[fileSha1]
}
