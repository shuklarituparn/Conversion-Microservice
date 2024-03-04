# Установка ⚙️

![944960 512 (1)](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/e31ed3cb-cfa1-454a-b664-5a2e63c579e3)


## Первые шаги (docker-compose) 🚀


* Клонируйте проект, выполнив следующую команду:

    `git clone git@github.com:shuklarituparn/Conversion-Microservice.git`
    

* Теперь выполните следующую команду, чтобы убедиться, что вы находитесь в корневой директории проекта:

    `cd Conversion-Microservice`



* Заполните файл `.env.example` и переименуйте его в `.env`


* Заполните следующее в файле docker compose

*    > `POSTGRES_USER: <Юзернэм вашего постгреса>` 

*    > `POSTGRES_PASSWORD: <Пароль вашего постгреса>`

*    >  `POSTGRES_DB: <Название ваши базы данных>`
     
*  Используйте образ docker `rituparnshukla/ffmpegservice-withoutgraphics:latest`, если у вас на компьютере нет графики nvidia


* Находясь в в корневой директории проекта, выполните следующую команду, чтобы запустить: `docker compose up`

> Убедитесь, что у вас установлен Docker перед выполнением вышеуказанной команды


* Cервис доступен по адресу `localhost:8085`


* Вы можете получить доступ к метрикам prometheus по адресу `localhost:8085/metrics`


* Графана доступна по адресу `localhost:3030`


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

`EMAIL_URL='<REDIRECT_EMAIL>'`

`SENTRY_DSN='<SENTRY>'`

`RESEND_API_KEY='<апи ключ от resend.com для отправки электронного письма>'`

> Вы можете получить апи ключ от resend после подтверждения вашего домена





