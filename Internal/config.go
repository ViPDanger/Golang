package Internal

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Adress      string `json:"addres"`
	Port        string `json:"port"`
	Data_File   string `json:"data_file"`
	PG_host     string `json:"pg_host"`
	PG_port     string `json:"pg_port"`
	PG_user     string `json:"pg_user"`
	PG_password string `json:"pg_password"`
	PG_bdname   string `json:"pg_bdname"`
}

func err_log(err error) bool {
	if err != nil {
		log.Println("Error: ", err, " - ", err.Error())
		panic(err)
	}
	return false
}

func Read_Config() Conf {
	var config Conf
	data := make([]byte, 1024)
	file, err := os.Open("cmd/config.cfg")
	if err_log(err) {
		panic(err)
	}
	len, err := file.Read(data)
	err_log(err)
	defer file.Close()
	data = append([]byte{byte(123)}, data[0:len]...)
	data = append(data[0:len], byte(125))
	err = json.Unmarshal(data, &config)
	err_log(err)
	if config.Adress == "" || config.Data_File == "" || config.Port == "" {
		log.Fatalln("config.txt read incorrectly.")
	}

	return config
}
