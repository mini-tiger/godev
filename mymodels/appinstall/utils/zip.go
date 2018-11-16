package utils

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
)

//func main() {
//	sfile := "c:\\image\\1.zip"        // zip文件
//	files := []string{"c:\\1", "d:\\1.bat", "c:\\safemon"} //需要压缩的文件夹或文件 的切片
//
//	err := Compress(files, sfile)
//	fmt.Println(err)
//	//time.Sleep(time.Duration(5)*time.Second)
//	err = UnCompress(sfile, "c:\\sss") // 压缩文件，解压缩路径
//	fmt.Println(err)
//}

func ReOsFile(f string) (of *os.File, err error) {
	of, err = os.Open(f)
	if err != nil {
		return
	}
	return
}

func Compress(filesStr []string, dst string) error {
	files := make([]*os.File, 0)
	for _, f := range filesStr { // 文件生成为 *os.File 类型
		if tf, err := ReOsFile(f); err != nil {
			log.Println("跳过压缩文件: ", f, "err:", err)
			continue
		} else {
			log.Println("添加压缩文件: ", f)
			files = append(files, tf)
		}

	}

	d, err := os.Create(dst)
	if err != nil {
		return errors.New(fmt.Sprintf("zipfile:%s,err:%s", dst, err))
	}
	defer d.Close()
	w := zip.NewWriter(d)
	defer w.Close()
	for _, file := range files {
		err := compress(file, "", w)
		if err != nil {
			return err
		}
	}
	return nil
}

func compress(file *os.File, prefix string, zw *zip.Writer) error { // prefix 是否为压缩的文件 添加子目录

	info, err := file.Stat()
	if err != nil {
		return err
	}
	if info.IsDir() {

		prefix = filepath.Join(prefix, info.Name())

		//header, err := zip.FileInfoHeader(info)
		//header.Name = filepath.Join(prefix, header.Name)
		//header.SetMode(os.ModeDir)

		fileInfos, err := file.Readdir(-1) // n < 0 所有目录中文件
		if err != nil {
			return err
		}
		for _, fi := range fileInfos {
			f, err := os.Open(file.Name() + "/" + fi.Name())
			if err != nil {
				return err
			}
			err = compress(f, prefix, zw) // 回调自身， 目录下文件逐一
			if err != nil {
				return err
			}
		}
	} else {

		header, err := zip.FileInfoHeader(info)
		header.Name = filepath.Join(prefix, header.Name)
		//header.Name = strings.TrimPrefix(prefix, string(filepath.Separator))
		//header.SetMode(os.ModeDir)
		if err != nil {
			return err
		}

		writer, err := zw.CreateHeader(header)
		if err != nil {
			return err
		}
		_, err = io.Copy(writer, file)
		file.Close()
		if err != nil {
			return err
		}
	}
	return nil
}

func IsDir(dir string) (b bool, err error) {
	dirInfo, err := os.Stat(dir)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil //不存在目录
		}
	}
	if dirInfo.IsDir() { //是目录
		return true, nil
	} else { //是文件
		return false, errors.New(fmt.Sprintf("dir:%s not dir", dir))
	}

	return false, nil

}

func UnCompress(src, dst string) (err error) {
	var bd bool
	if bd, err = IsDir(dst); err != nil {
		return
	}
	if !bd {
		err = os.MkdirAll(dst, os.ModePerm)
		if err != nil {
			return
		}
	}

	srcFile, err := os.Open(src)
	if err != nil {
		//logs.Error("open srcFile Error：", err.Error())
		return err
	}
	zipFile, err := zip.OpenReader(srcFile.Name())
	if err != nil {
		//logs.Error("Unzip File Error：", err.Error())
		return err
	}
	defer zipFile.Close()

	for _, innerFile := range zipFile.File {
		info := innerFile.FileInfo()

		if info.IsDir() {
			err = os.MkdirAll(innerFile.Name, os.ModePerm)
			if err != nil {
				log.Println("Unzip File mkdir Error : " + err.Error())
			}
			continue
		}
		srcFile, err := innerFile.Open()
		if err != nil {
			log.Println("Unzip File Error : " + err.Error())
			continue
		}

		// 使用zip模块压缩的文件夹，解压缩时不能判断文件夹是否是目录，每个文件创建时，都要判断是否有父目录
		//if bd, err = IsDir(filepath.Dir(filepath.Join(dst, innerFile.Name))); err != nil {
		//	continue
		//}
		//if !bd {
		//	os.MkdirAll(filepath.Dir(filepath.Join(dst, innerFile.Name)), os.ModePerm)
		//}

		newFile, err := os.Create(filepath.Join(dst, innerFile.Name))
		if err != nil {
			if strings.Contains(err.Error(), "The system cannot find the path specified") {
				os.MkdirAll(filepath.Dir(filepath.Join(dst, innerFile.Name)), os.ModePerm)
			} else {
				log.Println("Unzip File Create Error : " + err.Error())
				continue
			}

		}
		io.Copy(newFile, srcFile)
		newFile.Close()
	}
	return nil
}
