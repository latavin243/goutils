package fileutil

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"io"
	"io/fs"
	"os"
	"path"
	"path/filepath"
)

// ExePath returns the absolute path of the executable file.
func ExePath() string {
	fp, _ := filepath.Abs(os.Args[0])
	return fp
}

// ExeDir returns the directory of the executable file.
func ExeDir() string {
	return path.Dir(ExePath())
}

// RealPath returns the absolute path of the file.
func RealPath(fp string) (string, error) {
	if path.IsAbs(fp) {
		return fp, nil
	}
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(wd, fp), err
}

// Basename returns the last element of path. /path/to/file.txt -> file.txt
func Basename(fp string) string {
	return filepath.Base(fp)
}

// Dir returns the directory of path. /path/to/file.txt -> /path/to
func Dir(fp string) string {
	return filepath.Dir(fp)
}

// Ext returns the file name extension used by path. /path/to/file.txt -> .txt
func Ext(fp string) string {
	return filepath.Ext(fp)
}

// EnsureDir creates the necessary parent directory.
func EnsureDir(fp string, mode fs.FileMode) error {
	return os.MkdirAll(fp, mode)
}

func CreateFile(fp string) (*os.File, error) {
	return os.Create(fp)
}

func RemoveFile(fp string) error {
	return os.Remove(fp)
}

// IsExist checks whether a file or directory exists.
func IsExist(fp string) bool {
	_, err := os.Stat(fp)
	return err == nil || os.IsExist(err)
}

// IsFile checks whether the path is a file,
func IsFile(fp string) bool {
	f, e := os.Stat(fp)
	if e != nil {
		return false
	}
	return !f.IsDir()
}

// FileMTime returns the last modified time of the file.
// 10-digit timestamp
func FileMTime(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.ModTime().Unix(), nil
}

func FileSize(fp string) (int64, error) {
	f, e := os.Stat(fp)
	if e != nil {
		return 0, e
	}
	return f.Size(), nil
}

// FileMD5 returns the md5 value of the file.
func FileMD5(fp string) (string, error) {
	f, err := os.Open(fp)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r := bufio.NewReader(f)
	h := md5.New()

	_, err = io.Copy(h, r)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(h.Sum(nil)), nil
}
