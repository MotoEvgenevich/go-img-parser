package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	var logger = log.New(os.Stdout, "LOG MESSAGE: ", log.Ldate|log.Ltime)

	links, err := linkGenerate(logger)
	if err != nil {
		logger.Fatalf("Ошибка при генерации ссылок: %v\n", err)
	}

	if _, err := os.Stat("img"); os.IsNotExist(err) {
		err := os.Mkdir("img", os.ModePerm)
		if err != nil {
			logger.Fatalf("Ошибка при создании директории img: %v\n", err)
		}
	}

	for _, link := range links {
		err := parser(link, logger)
		if err != nil {
			logger.Printf("Ошибка при обработке ссылки %s: %v\n", link, err)
		}
	}
}

func parser(imgURL string, logger *log.Logger) error {
	// Извлечение имени файла из URL
	filename := imgURL[45:]
	filepath := "img/" + filename

	res, err := http.Get(imgURL)
	if err != nil {
		return fmt.Errorf("ошибка при загрузке изображения %s: %w", imgURL, err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("не удалось загрузить изображение, статус код: %d", res.StatusCode)
	}

	file, err := os.Create(filepath)
	if err != nil {
		return fmt.Errorf("ошибка при создании файла %s: %w", filepath, err)
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		return fmt.Errorf("ошибка при сохранении изображения %s: %w", filepath, err)
	}

	logger.Printf("Изображение успешно сохранено в файл %v\n", filepath)
	return nil
}

func linkGenerate(logger *log.Logger) ([]string, error) {
	links := []string{}
	baseURL := "http://193.7.160.230/web/esimo/black/wwf/img/black_swh_"

	for i := 0; i <= 120; i += 3 {
		num := strconv.Itoa(i)

		if i < 10 {
			num = "00" + num
		} else if i < 100 {
			num = "0" + num
		}

		link := baseURL + num + ".png"
		logger.Printf("Генерируем ссылку: %s", link)
		links = append(links, link)
	}

	if len(links) == 0 {
		err := fmt.Errorf("пустой список ссылок")
		logger.Printf("Ошибка: %v", err)
		return nil, err
	}

	return links, nil
}
