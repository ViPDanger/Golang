package internal

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Conf struct {
	Port   string `json:"port"`
	Adress string `json:"addres"`
}

func err_log(err error) bool {
	if err != nil {

		log.Println("Error: ", err, " - ", err.Error())
		return true
	}
	return false
}

func Read_Config() (string, string) {
	var config Conf
	data := make([]byte, 1024)
	file, err := os.Open("cmd/config.cfg")
	if err_log(err) {
		panic(err)
	}
	fmt.Println(string(data))
	len, err := file.Read(data)
	err_log(err)
	defer file.Close()
	data = append([]byte{byte(123)}, data[0:len]...)
	data = append(data[0:len], byte(125))
	err = json.Unmarshal(data, &config)
	err_log(err)
	return config.Adress, config.Port
}
