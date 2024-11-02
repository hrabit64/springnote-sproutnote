package utils

import (
	"io"
	"os"
	"path/filepath"
)

func RemoveFile(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}

	err = os.RemoveAll(filePath)
	if err != nil {
		return false, err
	}

	return true, nil
}

func CheckFileExist(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}

	return true, nil
}

func CheckIsDir(filePath string) (bool, error) {
	fileInfo, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false, nil
	}

	return fileInfo.IsDir(), nil
}

func CopyFile(filePath string, destPath string) error {
	sourceFile, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destinationFile, err := os.Create(destPath)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// 파일의 내용을 복사합니다.
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}
	return nil
}

func CopyDir(filePath string, destPath string) error {
	return filepath.Walk(filePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 소스 경로에서 상대 경로를 가져옵니다.
		relPath, err := filepath.Rel(filePath, path)
		if err != nil {
			return err
		}

		destPath := filepath.Join(destPath, relPath)

		// 디렉토리인지 파일인지 확인합니다.
		if info.IsDir() {
			// 디렉토리를 만듭니다.
			if err := os.MkdirAll(destPath, info.Mode()); err != nil {
				return err
			}
		} else {
			// 파일을 복사합니다.
			if err := CopyFile(path, destPath); err != nil {
				return err
			}
		}
		return nil
	})
}
