package main

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("failed to get current directory")
		return
	}

	fmt.Printf("the current dir:%s", dir)

	testDir := filepath.Join(dir, "test")
	// 调用示例
	err = CreateZip("output.zip", testDir, "config", "model", "lib", "manager.sh", "release.txt", "svapp")
	if err != nil {
		fmt.Println("创建 zip 包时发生错误:", err)
	} else {
		fmt.Println("已成功创建 zip 包")
	}

	//err = Zip(testDir, "output.zip")
	//if err != nil {
	//	fmt.Println("创建 zip 包时发生错误:", err)
	//} else {
	//	fmt.Println("已成功创建 zip 包")
	//}
}

func Zip(srcDir, zipFileName string) error {

	// 预防：旧文件无法覆盖
	os.RemoveAll(zipFileName)

	// 创建：zip文件
	zipfile, err := os.Create(zipFileName)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	// 打开：zip文件
	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	// 遍历路径信息
	filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {

		// 如果是源路径，提前进行下一个遍历
		if path == srcDir {
			return nil
		}

		// 获取：文件头信息
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		header.Name = strings.TrimPrefix(path, srcDir+`/`)

		// 判断：文件是不是文件夹
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置：zip的文件压缩算法
			header.Method = zip.Store
		}

		// 创建：压缩包头部信息
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}

		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			defer file.Close()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func CreateZip(zipName, dir string, filterPaths ...string) error {
	// 创建一个新的zip文件
	zipFile, err := os.Create(zipName)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	// 创建一个zip.Writer来写入zip文件
	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	// 遍历给定的过滤器路径
	for _, filterPath := range filterPaths {
		// 获取过滤器的绝对路径
		absFilterPath := filepath.Join(dir, filterPath)

		// 检查过滤器路径是否存在
		_, err := os.Stat(absFilterPath)
		if err != nil {
			fmt.Printf("过滤器路径不存在：%s\n", absFilterPath)
			continue
		}

		// 打开文件或文件夹
		file, err := os.Open(absFilterPath)
		if err != nil {
			return err
		}
		defer file.Close()

		// 获取文件信息
		fileInfo, err := file.Stat()
		if err != nil {
			return err
		}

		// 如果是文件夹，则递归地将文件夹中的文件添加到 zip.Writer 中
		if fileInfo.IsDir() {

			selfDirName := fileInfo.Name()
			// 根据过滤器路径递归添加文件或文件夹到zip包中
			filepath.Walk(absFilterPath, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}

				// 获取相对于过滤器路径的相对路径
				relPath, err := filepath.Rel(absFilterPath, path)
				if err != nil {
					return err
				}

				// 自身文件夹
				if relPath == "." {
					relPath = selfDirName
				} else {
					relPath = filepath.Join(selfDirName, relPath)
				}

				if info.IsDir() {
					// 在文件夹路径末尾添加斜杠，以确保它以文件夹形式显示
					relPath += "/"
				}

				// 创建一个新的zip文件条目
				zipEntry, err := zipWriter.Create(relPath)
				if err != nil {
					return err
				}

				// 如果当前路径是一个文件夹，则不需要写入内容，只需创建一个空文件夹
				if info.IsDir() {
					return nil
				}

				// 打开当前文件
				file, err := os.Open(path)
				if err != nil {
					return err
				}
				defer file.Close()

				// 将文件内容复制到zip文件条目中
				_, err = io.Copy(zipEntry, file)
				if err != nil {
					return err
				}

				return nil
			})
		} else {
			// 如果是文件，则直接将文件内容拷贝到 zip 中的文件中
			fileInZip, err := zipWriter.Create(fileInfo.Name())
			if err != nil {
				return err
			}

			_, err = io.Copy(fileInZip, file)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

//// CreateZip 函数接收一个 zip 文件名和可变参数的文件路径，将这些文件路径打包成一个 zip 包
//func CreateZip(zipName, dir string, paths ...string) error {
//	// 创建一个新的 zip 文件
//	zipFile, err := os.Create(zipName)
//	if err != nil {
//		return err
//	}
//	defer zipFile.Close()
//
//	// 创建一个 zip.Writer，用于写入 zip 文件
//	zipWriter := zip.NewWriter(zipFile)
//	defer zipWriter.Close()
//
//	// 遍历每个文件路径
//	for _, path := range paths {
//		tmpPath := filepath.Join(dir, path)
//		err := addFileToZip(zipWriter, tmpPath)
//		if err != nil {
//			return err
//		}
//	}
//
//	return nil
//}

// addFileToZip 函数将指定的文件或文件夹添加到 zip.Writer 中
func addFileToZip(zipWriter *zip.Writer, fp string) error {
	// 打开文件或文件夹
	file, err := os.Open(fp)
	if err != nil {
		return err
	}
	defer file.Close()

	// 获取文件信息
	fileInfo, err := file.Stat()
	if err != nil {
		return err
	}

	// 如果是文件夹，则递归地将文件夹中的文件添加到 zip.Writer 中
	if fileInfo.IsDir() {
		err = filepath.Walk(fp, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			// 创建 zip 中的文件或文件夹
			zipPath, err := filepath.Rel(fp, path)
			if err != nil {
				return err
			}

			if info.IsDir() {
				// 创建文件夹
				_, err = zipWriter.Create(zipPath + "/")
				if err != nil {
					return err
				}
			} else {
				// 创建文件
				file, err := zipWriter.Create(zipPath)
				if err != nil {
					return err
				}

				// 将文件内容拷贝到 zip 中的文件中
				sourceFile, err := os.Open(path)
				if err != nil {
					return err
				}
				defer sourceFile.Close()

				_, err = io.Copy(file, sourceFile)
				if err != nil {
					return err
				}
			}

			return nil
		})

		if err != nil {
			return err
		}
	} else {
		// 如果是文件，则直接将文件内容拷贝到 zip 中的文件中
		fileInZip, err := zipWriter.Create(fileInfo.Name())
		if err != nil {
			return err
		}

		_, err = io.Copy(fileInZip, file)
		if err != nil {
			return err
		}
	}

	return nil
}
