package config

import (
	"encoding/json"
	"log"
	"os"
)

type Conf struct {
	Adress          string `json:"adress"`
	Port            string `json:"port"`
	Data_File       string `json:"data_file"`
	PG_host         string `json:"pg_host"`
	PG_port         string `json:"pg_port"`
	PG_user         string `json:"pg_user"`
	PG_password     string `json:"pg_password"`
	PG_bdname       string `json:"pg_bdname"`
	PG_Con_Attempts string `json:"max_connection_attempts"`
}

func Err_log(err error) bool {
	if err != nil {
		panic(err)
	}
	return false
}

func Read_Config() Conf {
	var config Conf
	data := make([]byte, 1024)

	file, err := os.Open("cmd/config.cfg")
	if err != nil {
		// DEBUG conf
		file, err = os.Open("config.cfg")
		if Err_log(err) {
			panic(err)
		}
	}
	len, err := file.Read(data)
	Err_log(err)
	defer file.Close()
	data = append([]byte{byte(123)}, data[0:len]...)
	data = append(data[0:len], byte(125))
	err = json.Unmarshal(data, &config)
	Err_log(err)
	if config.Adress == "" || config.Data_File == "" || config.Port == "" {
		log.Fatalln("config.txt read incorrectly.")
	}

	return config
}
