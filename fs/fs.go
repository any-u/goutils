package fs

import (
	"io"
	"io/ioutil"
	"os"
	"path"
)

/**
 * 读取文件
 * @param {string} file 文件路径
 * @return {[]byte} 文件数据
 */
func ReadFile(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}

	return data
}

/**
 * 写入文件
 * @param {string} file 文件路径
 * @param {string} data 写入数据
 */
func WriteFile(file string, data string) {
	err := ioutil.WriteFile(file, []byte(data), 0616)
	if err != nil {
		panic(err)
	}
}

/**
 * 向文件追加内容
 * @param {string} path 文件路径
 * @param {string} data 追加的内容
 */
func AppendFile(path string, data string) {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		panic(err)
	}
}

/**
 * 创建文件夹
 * @param {string} path 文件夹路径
 */
func Mkdir(path string) {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}
}

/**
 * 复制文件
 * @param {string} src 要复制的源文件名
 * @param {string} dest 复制操作的模板文件名
 */
func CopyFile(src string, dest string) {
	originFile, err := os.Open(src)
	if err != nil {
		panic(err)
	}

	defer originFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		panic(err)
	}

	defer destFile.Close()

	_, err = io.Copy(destFile, originFile)
	if err != nil {
		panic(err)
	}

	err = destFile.Sync()
	if err != nil {
		panic(err)
	}
}

/**
 * 文件或文件夹重命名或移动
 * @param {string} oldPath 旧路径
 * @param {string} newPath 新路径
 * @return {*}
 */
func Rename(oldPath string, newPath string) {
	err := os.Rename(oldPath, newPath)
	if err != nil {
		panic(err)
	}
}

/**
 * 删除文件或文件夹
 * @param {string} path 路径
 */
func Remove(path string) {
	err := os.RemoveAll(path)
	if err != nil {
		panic(err)
	}
}

/**
 * 判断路径是否存在
 * @param {string} path 路径
 */
func PathExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/**
 * 确保目录为空。如果目录不为空，则删除目录内容。如果目录不存在，则创建它。目录本身不会被删除。
 * @param {string} dir
 * @return {*}
 */
func EmptyDir(dir string) {
	items, err := ioutil.ReadDir(dir)
	if err != nil {
		Mkdir(dir)
		return
	}
	for _, item := range items {
		os.RemoveAll(path.Join([]string{dir, item.Name()}...))
	}
}
