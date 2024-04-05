package txt_file

import (
	"log"
	"os"
	"strings"

	conf "github.com/ViPDanger/Golang/Internal/config"
)

func field_splitter(r rune) bool {
	if r == '\n' {
		return true
	} else {
		return false
	}

}

func TXT_Read_Data(filename string) *[]string {
	data := make([]byte, 1024)
	file, err := os.Open(filename)
	if conf.Err_log(err) {
		os.Create(filename)
	}

	file_size, err := file.Read(data)
	m_data := strings.FieldsFunc(string(data[:file_size]), field_splitter)
	defer file.Close()
	conf.Err_log(err)
	return &m_data
}

func TXT_Add_Data(str string) {
	config := conf.Read_Config()
	// Считывание данных с файла, корректировка
	data := TXT_Read_Data(config.Data_File)
	*data = append((*data), str)
	// Запись обьекта
	file, err := os.Create(config.Data_File)
	conf.Err_log(err)
	for i := 0; i < len(*data); i++ {
		file.Write([]byte((*data)[i] + "\n"))
	}

	defer file.Close()
	log.Println("Done.")

}

func TXT_Delete_Data(del_int int) {
	config := conf.Read_Config()
	// Считывание данных с файла, корректировка
	data := TXT_Read_Data(config.Data_File)
	// Запись обьекта
	file, err := os.Create(config.Data_File)
	conf.Err_log(err)
	for i := 0; i < len(*data); i++ {
		if i != del_int {
			file.Write([]byte((*data)[i] + "\n"))
		}
	}

	defer file.Close()
	log.Println("Done.")

}
