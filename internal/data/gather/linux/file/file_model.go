package file

import (
	"github.com/spf13/afero"
	"os"
	"time"
)

type Info struct {
	Fs         afero.Fs    `json:"-"`
	Path       string      `json:"path"`
	Name       string      `json:"name"`
	User       string      `json:"user"`
	Group      string      `json:"group"`
	Uid        string      `json:"uid"`
	Gid        string      `json:"gid"`
	Extension  string      `json:"extension"`
	Content    string      `json:"content"`
	Size       int64       `json:"size"`
	IsDir      bool        `json:"isDir"`
	IsSymlink  bool        `json:"isSymlink"`
	IsHidden   bool        `json:"isHidden"`
	LinkPath   string      `json:"linkPath"`
	Type       string      `json:"type"`
	Mode       string      `json:"mode"`
	MimeType   string      `json:"mimeType"`
	UpdateTime time.Time   `json:"updateTime"`
	ModTime    time.Time   `json:"modTime"`
	FileMode   os.FileMode `json:"-"`
	Items      []*Info     `json:"items"`
	ItemTotal  int         `json:"itemTotal"`
}

type Option struct {
	Path       string `json:"path"`
	Search     string `json:"search"`
	ContainSub bool   `json:"containSub"`
	Expand     bool   `json:"expand"`
	Dir        bool   `json:"dir"`
	ShowHidden bool   `json:"showHidden"`
	Page       int    `json:"page"`
	PageSize   int    `json:"pageSize"`
}

type Create struct {
	Path      string `json:"path" validate:"required"`
	Content   string `json:"content"`
	IsDir     bool   `json:"isDir"`
	Mode      int64  `json:"mode" validate:"required"`
	IsLink    bool   `json:"isLink"`
	IsSymlink bool   `json:"isSymlink"`
	LinkPath  string `json:"linkPath"`
	Sub       bool   `json:"sub"`
}

type Delete struct {
	Path  string `json:"path" validate:"required"`
	IsDir bool   `json:"isDir"`
}

type BatchDelete struct {
	Paths []string `json:"paths" validate:"required"`
	IsDir bool     `json:"isDir"`
}

type Compress struct {
	Files   []string `json:"files" validate:"required"`
	Dst     string   `json:"dst" validate:"required"`
	Type    string   `json:"type" validate:"required"`
	Name    string   `json:"name" validate:"required"`
	Replace bool     `json:"replace"`
}

type DeCompress struct {
	Dst  string `json:"dst"  validate:"required"`
	Type string `json:"type"  validate:"required"`
	Path string `json:"path" validate:"required"`
}

type Edit struct {
	Path    string `json:"path"  validate:"required"`
	Content string `json:"content"  validate:"required"`
}

type Wget struct {
	Url  string `json:"url" validate:"required"`
	Path string `json:"path" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type Move struct {
	Type     string   `json:"type" validate:"required"`
	OldPaths []string `json:"oldPaths" validate:"required"`
	NewPath  string   `json:"newPath" validate:"required"`
}

type Download struct {
	Paths    []string `json:"paths" validate:"required"`
	Type     string   `json:"type" validate:"required"`
	Name     string   `json:"name" validate:"required"`
	Compress bool     `json:"compress" validate:"required"`
}

type DirSizeReq struct {
	Path string `json:"path" validate:"required"`
}

type DirSizeRes struct {
	Size float64 `json:"size" validate:"required"`
}

type Rename struct {
	OldName string `json:"oldName" validate:"required"`
	NewName string `json:"newName" validate:"required"`
}

type RoleUpdate struct {
	Path  string `json:"path" validate:"required"`
	User  string `json:"user" validate:"required"`
	Group string `json:"group" validate:"required"`
	Sub   bool   `json:"sub" validate:"required"`
}

type Tree struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Path     string `json:"path"`
	Children []Tree `json:"children"`
}

type SearchUploadWithPage struct {
	Page     int    `json:"page" validate:"required,number"`
	PageSize int    `json:"pageSize" validate:"required,number"`
	Path     string `json:"path" validate:"required"`
}
