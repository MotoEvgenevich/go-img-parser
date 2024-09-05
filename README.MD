# Image Downloader

[English](#english) | [Русский](#русский)

## English

### Overview

This project is a simple Go application that downloads images from a set of generated URLs and saves them into a local directory. The links are generated for every 3-hour interval, up to 120 hours. Each image is stored in the `img` directory.

### Features

- Automatically generates links for images based on a predefined URL structure.
- Downloads images over HTTP and saves them to the `img` directory.
- Creates the `img` directory if it does not exist.

### Requirements

- Go 1.19+
- Internet connection

### Installation

1. Clone the repository:

    ```bash
    git clone https://github.com/yourusername/image-downloader.git
    cd image-downloader
    ```

2. Run the application:

    ```bash
    go run main.go
    ```

### How It Works

- The `linkGenerate` function creates a list of URLs, each pointing to an image that corresponds to a 3-hour interval.
- The `parser` function downloads each image and saves it to the `img` directory.
- If the `img` directory does not exist, it is created automatically.

### Example Output

```bash
Image successfully saved to img/black_swh_000.png
Image successfully saved to img/black_swh_003.png
...
```

### License
This project is open-source and available under the MIT License.

## Русский

### Обзор
Этот проект — простое Go-приложение, которое загружает изображения с набора сгенерированных URL-адресов и сохраняет их в локальную директорию. Ссылки генерируются с интервалом в 3 часа на 120 часов вперед. Каждое изображение сохраняется в директорию img.

### Особенности
Автоматическая генерация ссылок на изображения на основе заранее заданной структуры URL.

Загрузка изображений через HTTP и их сохранение в директорию img.

Автоматическое создание директории img, если она не существует.

### Требования

Go 1.19+

Подключение к интернету

### Установка
Клонируйте репозиторий:


```bash
git clone https://github.com/yourusername/image-downloader.git
cd image-downloader
```

### Запустите приложение:

```bash
go run main.go
```
### Как это работает
Функция linkGenerate создает список URL-адресов, каждый из которых указывает на изображение, соответствующее 3-часовому интервалу.

Функция parser загружает каждое изображение и сохраняет его в директорию img.
Если директория img не существует, она создается автоматически.

### Пример вывода

``` bash
Изображение успешно сохранено в файл img/black_swh_000.png
Изображение успешно сохранено в файл img/black_swh_003.png
...
```

### Лицензия
Этот проект является открытым и доступен под лицензией MIT.

