package task

import (
	"encoding/json"
	"log"
	"os"
	"path/filepath"
)

// taskFilePath return path file tasks.json
// according on current directory
func taskFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		log.Printf("error getting current working directory : %s\n", err.Error())
		return ""
	}

	return filepath.Join(cwd, "tasks.json")
}

func ReadTaskFromFile() ([]Task, error) {
	filePath := taskFilePath()

	// check if file is exists
	// jika file tidak ada maka buat file dan isi dengan empty list
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		err := os.WriteFile(filePath, []byte("[]"), os.ModePerm)
		if err != nil {
			log.Printf("error creting file : %s\n", err.Error())
			return nil, err
		}

		return []Task{}, nil
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("error reading file : %s\n", err.Error())
		return nil, err
	}

	var Task []Task
	err = json.Unmarshal(file, &Task)
	if err != nil {
		log.Printf("error decoding data : %s\n", err.Error())
		return nil, err
	}

	return Task, nil
}

func WriteTaskToFile(data []Task) error {
	filepath := taskFilePath()
	file, err := os.Create(filepath)
	if err != nil {
		log.Printf("error creating file : %s\n", err.Error())
		return err
	}
	defer func() {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file : %s\n", err.Error())
			return
		}
	}()

	err = json.NewEncoder(file).Encode(data)
	if err != nil {
		log.Printf("error encode file : %s\n", err.Error())
		return err
	}

	return nil
}
