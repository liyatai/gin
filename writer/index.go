package writer

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
)

func WriteStringToFile(content, filePath string) error {
	// 1. 校验入参合法性
	if len(filePath) == 0 {
		return errors.New("文件路径不能为空")
	}
	if len(content) == 0 {
		return errors.New("写入内容不能为空（若需写入空文件，可传入空字符串并自行确认）")
	}

	// 2. 解析文件路径，分离目录和文件名
	dir := filepath.Dir(filePath)       // 获取文件所在目录（如 ./service/aa.go → ./service）
	fileName := filepath.Base(filePath) // 获取文件名（如 ./service/aa.go → aa.go）
	if len(fileName) == 0 {
		return errors.New("文件路径解析失败，未提取到合法文件名")
	}

	// 3. 创建目录（递归创建，如 ./a/b/c 不存在则逐层创建）
	// 权限 0755：当前用户可读可写可执行，其他用户可读可执行
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("创建目录失败：%w", err) // 包装错误，保留原始错误信息
	}

	// 4. 拼接完整文件路径（防止解析异常，重新拼接确保正确性）
	fullPath := filepath.Join(dir, fileName)

	// 5. 写入文件（覆盖原有内容，权限 0644：当前用户可读可写，其他用户只读）
	// 若需追加内容，可改用 os.OpenFile + os.O_APPEND 模式
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("写入文件失败：%w", err)
	}

	fmt.Printf("文件写入成功！路径：%s\n", fullPath)
	return nil
}
