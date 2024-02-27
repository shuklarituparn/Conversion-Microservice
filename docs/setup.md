# Установка ⚙️

![944960 512 (1)](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/e31ed3cb-cfa1-454a-b664-5a2e63c579e3)


## Первые шаги 🚀

* Клонируйте проект, выполнив следующую команду:

    `git clone git@github.com:shuklarituparn/Conversion-Microservice.git`
    
* Теперь выполните следующую команду, чтобы убедиться, что вы находитесь в корневой директории проекта:

    `cd Conversion-Microservice`
    
* Убедитесь, что у вас установлен и работает PostgreSQL

* Находясь в в корневой директории проекта, выполните следующую команду, чтобы запустить Kafka:

   `docker compose up`
   
   > Убедитесь, что у вас установлен Docker перед выполнением вышеуказанной команды

* Теперь заполните файл `.env.example` и переименуйте его в `.env`

* Запустите `go mod tidy`, чтобы установить все зависимости


* Перейдите в директории `cmd`, а затем в директории `video-converter`
 
* Вы можете запустить любую из следующих команд для запуска сервиса:
    `go run main.go` или `./main`

* Вы можете взаимодействовать с сервисом по адресу `localhost:8085` или по `URL обратного вызова`, описанному в `.env`

---


## ENV Файл 📝


`MONGO_URL='<URL MongoDB>'`

> Ваш URL Atlas в формате
 
 >`mongodb+srv://user:password@yourmongocluster/?retryWrites=true&w=majority`
 
`VK_CLIENT_ID='<Идентификатор вашего приложения VK>'`

> Вы можете создать свое приложение на `dev.vk.com`

`VK_CLIENT_SECRET='<Секрет вашего приложения VK>'`

 >Вы можете получить его в настройках вашего приложения на сайте `dev.vk.com`


`REDIRECT_URL='<URL перенаправления OAuth в формате "https://example.com/callback">'`

> Вам нужно настроить это в настройках вашего приложения после включения `openAPI`

![Screenshot from 2024-02-27 22-29-28](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/bc0e5187-e533-4b11-babe-fed4775d10f4)



`SESSION_SECRET='<Секрет для создания/хранения куки>'`

`EMAIL='<Электронная почта вашего сервиса>'`

`EMAIL_KEY='<Ваш ключ электронной почты для отправки писем с использованием SMTP>'`

`POSTGRES_USERNAME='<Имя пользователя вашей PostgreSQL>'`

`POSTGRES_PASSWORD='<Пароль вашей PostgreSQL>'`

`DB_NAME='<Имя вашей базы данных>'`





