package util

import "os"

// ReadIfExists 阅读文件是否存在，如果存在则读取
func ReadIfExists(path string) (string, error) {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return "", nil
		}
		return "", err
	}
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
