package Internal

import (
	"log"
	"os"
	"strings"
)

func field_splitter(r rune) bool {
	if r == '\n' {
		return true
	} else {
		return false
	}

}

func txt_Read_Data(filename string) *[]string {
	data := make([]byte, 1024)
	file, err := os.Open(filename)
	if err_log(err) {
		os.Create(filename)
	}

	file_size, err := file.Read(data)
	m_data := strings.FieldsFunc(string(data[:file_size]), field_splitter)
	defer file.Close()
	err_log(err)
	return &m_data
}

func txt_Add_Data(str string) {
	config := Read_Config()
	// Считывание данных с файла, корректировка
	data := txt_Read_Data(config.Data_File)
	*data = append((*data), str)
	// Запись обьекта
	file, err := os.Create(config.Data_File)
	err_log(err)
	for i := 0; i < len(*data); i++ {
		file.Write([]byte((*data)[i] + "\n"))
	}

	defer file.Close()
	log.Println("Done.")

}

func txt_Delete_Data(del_int int) {
	config := Read_Config()
	// Считывание данных с файла, корректировка
	data := txt_Read_Data(config.Data_File)
	// Запись обьекта
	file, err := os.Create(config.Data_File)
	err_log(err)
	for i := 0; i < len(*data); i++ {
		if i != del_int {
			file.Write([]byte((*data)[i] + "\n"))
		}
	}

	defer file.Close()
	log.Println("Done.")

}
