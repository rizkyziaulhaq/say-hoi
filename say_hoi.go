package sayhoi

import "os"

func SayHai(name string) string {
	return "Hai Zia" + name
}

func CreateNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)

	return nil
}
