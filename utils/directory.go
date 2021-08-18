// @Title  
// @Description  
// @Author  Wangwengang  2021/8/18 下午10:50
// @Update  Wangwengang  2021/8/18 下午10:50
package utils

import (
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
