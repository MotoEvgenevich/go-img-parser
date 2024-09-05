package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

func main() {

	links := linkGenerate()

	if _, err := os.Stat("img"); os.IsNotExist(err) {
		err := os.Mkdir("img", os.ModePerm)
		if err != nil {
			fmt.Println("Ошибка при создании директории img:", err)
			return
		}
	}

	for _, link := range links {
		parser(link)
	}
}

func parser(imgURL string) {

	filename := imgURL[45:]

	filepath := "img/" + filename

	res, err := http.Get(imgURL)
	if err != nil {
		fmt.Println("Ошибка при загрузке изображения:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Ошибка: Не удалось загрузить изображение, статус код:", res.StatusCode)
		return
	}

	file, err := os.Create(filepath)
	if err != nil {
		fmt.Println("Ошибка при создании файла:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, res.Body)
	if err != nil {
		fmt.Println("Ошибка при сохранении изображения:", filepath, err)
		return
	}

	fmt.Printf("Изображение успешно сохранено в файл %v\n", filepath)
}

func linkGenerate() []string {
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
		links = append(links, link)
	}

	return links
}
