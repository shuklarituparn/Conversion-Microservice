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


`SESSION_SECRET='<Секрет для создания/хранения куки>'`

`EMAIL='<Электронная почта вашего сервиса>'`

`EMAIL_KEY='<Ваш ключ электронной почты для отправки писем с использованием SMTP>'`

`POSTGRES_USERNAME='<Имя пользователя вашей PostgreSQL>'`

`POSTGRES_PASSWORD='<Пароль вашей PostgreSQL>'`

`DB_NAME='<Имя вашей базы данных>'`



---


# Setup 🔧

![944960 512 (1)](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/e31ed3cb-cfa1-454a-b664-5a2e63c579e3)


## First steps 🚀

* Clone the project by running the following command:

    `git clone git@github.com:shuklarituparn/Conversion-Microservice.git`
    
* Now run the following command to ensure you're in the root directory of the project:

    `cd Conversion-Microservice`
    
* Make sure you have PostgreSQL installed and running.

* While in the root directory of the project, run the following command to bring Kafka up:

   `docker compose up`
   
   > Make sure you have Docker installed before running the above command.

* Now fill the `.env.example` file and rename it to `.env`.

* Run `go mod tidy` to install all dependencies.


* Navigate to the `cmd` directory and then to the `video-converter` directory.
 
* You can run either of the commands below to run the service:
    `go run main.go` or `./main`.

* You can interact with the service at `localhost:8085` or at the `callback URL` described in `.env`.

---


## Environment File 📁

`
MONGO_URL='<MongoDB URL>'
`

> Your Atlas URL in the format
 
 >`mongodb+srv://user:password@yourmongocluster/?retryWrites=true&w=majority`
 
`VK_CLIENT_ID='<Your VK Application ID>'`

> You can create your application at `dev.vk.com`

`VK_CLIENT_SECRET='<Your VK Application Secret>'`

 >You can get it in the settings of your application at the site `dev.vk.com`


`REDIRECT_URL='<OAuth Redirect URL in the format "https://example.com/callback">'`

> You need to set it up in the settings of your application, after turning on `openAPI`


`SESSION_SECRET='<Secret for Creating/Storing Cookies>'`


`EMAIL='<Your Service Email>'`

`EMAIL_KEY='<Your Email Key to Send Emails Using SMTP>'`

`POSTGRES_USERNAME='<Your PostgreSQL Username>'`

`POSTGRES_PASSWORD='<Your PostgreSQL Password>'`

`DB_NAME='<Your Database Name>'`



