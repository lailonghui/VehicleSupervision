/*
@Time : 2020/12/3 15:57
@Author : lai
@Description : 对文件和目录的操作
@File : file
*/
package file

import "os"

const (
	_        = iota             // iota = 0
	KB int64 = 1 << (10 * iota) // iota = 1，2的10次方
	MB                          // iota = 2, 2的20次方
	GB                          // iota = 3, 2的30次方
	TB                          // iota = 4, 2的40次方
)

// 判断目录是否存在
func IsDirExist(path string) (err error) {
	_, err = os.Stat(path)
	return
}

// 新建目录
func CreateDir(path string) error {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return err
	} else {
		return nil
	}
}

// 删除目录
func RemoveDir(path string) error {
	return os.RemoveAll(path)
}

// 判断文件是否存在
func IsFileExist(path string) error {
	_, err := os.Stat(path)
	return err
}

// 新建文件
func createFile(path string) error {
	file, err := os.Create(path)
	file.Close()
	return err
}

// 删除文件
func RemoveFile(path string) error {
	return os.Remove(path)
}
