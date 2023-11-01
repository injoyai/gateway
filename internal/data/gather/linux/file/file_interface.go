package file

type Interface interface {
	List(op Option) (Info, error)
	SearchUploadWithPage(req SearchUploadWithPage) (int64, interface{}, error)
	Tree(op Option) ([]Tree, error)
	Create(op Create) error
	Delete(op Delete) error
	BatchDelete(op BatchDelete) error
	ChangeMode(op Create) error
	// Compress 压缩文件
	Compress(c Compress) error
	// DeCompress 解压文件
	DeCompress(c DeCompress) error
	GetContent(op Option) (Info, error)
	SaveContent(edit Edit) error
	Download(d Download) (string, error)
	Size(req DirSizeReq) (DirSizeRes, error)
	// Rename 重命名
	Rename(req Rename) error
	// Wget wget
	Wget(w Wget) (string, error)
	// Move 移动文件
	Move(m Move) error
	ChangeOwner(req RoleUpdate) error
}
