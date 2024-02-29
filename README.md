# 🎬 Микросервис конвертации



![Screenshot from 2024-02-27 21-23-46](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/ce4cfde8-0c11-496b-be62-4c3f73e8206c)

![Screenshot from 2024-02-27 21-24-35](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/6085f17c-e6c9-41a7-86bf-e1a5efada1f5)

Логотип ВК является собственностью «ООО "ВК"»

## Обзор

Микросервис конвертации - это надежное решение, разработанное для обработки различных медиа-операций. От точной обрезки видео по определенным временным рамкам до конвертации в различные форматы, создания скриншотов и добавления водяных знаков - этот микросервис предлагает полный набор функций.

Сервис доступен по сайту: http://videoconversion.heyaadi.ru/

## Особенности

- **Обрезка видео**: Точная обрезка на основе заданных временных рамок.
- **Конвертация формата**: Беспрепятственная конвертация видео в различные форматы, в настоящее время поддерживается от MP4 -> MOV, MKV, MP3.
- **Создание скриншотов**: Легкое создание скриншотов из видео.
- **Добавление водяных знаков**: Добавление настраиваемых водяных знаков к видео для брендинга или идентификации.

## Технологический стек

- **Фронтенд**: HTML + Tailwind CSS
- **Бэкенд**: Go (GIN для маршрутизации)
- **Базы данных**: GORM с PostgreSQL (локальная), MongoDB (облачная)
- **Сообщения**: Kafka
- **Электронная почта**: Hermes для генерации шаблонов, Gomail для отправки
- **Обработка медиа**: FFMPEG
- **Мониторинг**: Prometheus для метрик, Grafana для визуализации
- **Аутентификация пользователей**: Использование OAuth для входа пользователей

## Использование и установка

- [Использование](docs/usage.md)
- [Установка](docs/setup.md)
- [Что еще][/docs/what's_coming.md]


## Вклад

Ваши вклады приветствуются!

## Лицензия

Этот проект лицензирован в соответствии с [лицензией MIT](LICENSE).




# 🎬 Conversion Microservice


## Overview

The Conversion Microservice is a robust solution designed to handle diverse media operations. From cutting videos based on specific time frames to converting them into various formats, taking screenshots, and even adding watermarks, this microservice offers a comprehensive suite of functionalities.

The site is available to try at : http://videoconversion.heyaadi.ru/

## Features

- **Video Cutting**: Precision cutting based on defined start and end times.
- **Format Conversion**: Seamlessly convert videos to different formats, with MP4 to MP3, MKV, MOV currently supported.
- **Screenshot Capture**: Effortlessly capture screenshots from videos.
- **Watermark Addition**: Add customizable watermarks to videos for branding or identification purposes.

## Tech Stack

- **Frontend**: HTML + Tailwind CSS
- **Backend**: Go (GIN for routing)
- **Databases**: GORM with PostgreSQL (local), MongoDB (cloud)
- **Messaging**: Kafka
- **Email**: Hermes for template generation, Gomail for sending
- **Media Processing**: FFMPEG
- **Monitoring**: Prometheus for metrics, Grafana for visualization
- **User Authentication**: Using OAuth to log in users

## Usage and Installation

- [Usage](docs/usage_eng.md)
- [Installation](docs/setup_eng.md)
- [What's next][/docs/what's_coming_en.md]

## Contributing

Contributions are welcome!

## License

This project is licensed under the [MIT License](LICENSE).
