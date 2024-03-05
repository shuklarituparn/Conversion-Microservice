# Setup 🔧

![944960 512 (1)](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/b0db7aef-2982-416c-96bc-7f877a6e9ce7)

## First steps (DOCKER-COMPOSE) 🚀


* Clone the project by running the following command:

    `git clone git@github.com:shuklarituparn/Conversion-Microservice.git`
    

* Now run the following command to ensure you're in the root directory of the project:

    `cd Conversion-Microservice`


* Fill the file `.env.example` and rename it to `.env`


* Fill the following in the `docker-compose` file

*    > `POSTGRES_USER: <Your postgres username>`

*    > `POSTGRES_PASSWORD: <Password of your postgres>`

*    >  `POSTGRES_DB: <Name of your database>`


   
* While located in the root folder of the project run `docker-compose up`

> Check that you have docker up and running


* Now you can access the service at `localhost:8085`


* You can access the metrics at `localhost:8085/metrics`


* You can access grafana at  `localhost:3030`
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

![Screenshot from 2024-02-27 22-29-28](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/ffa984be-92e5-4dec-9245-b6cf62dea457)



`SESSION_SECRET='<Secret for Creating/Storing Cookies>'`


`EMAIL='<Your Service Email>'`

`EMAIL_KEY='<Your Email Key to Send Emails Using SMTP>'`

`POSTGRES_USERNAME='<Your PostgreSQL Username>'`

`POSTGRES_PASSWORD='<Your PostgreSQL Password>'`

`DB_NAME='<Your Database Name>'`


`EMAIL_URL='<REDIRECT_EMAIL>'`

`SENTRY_DSN='<SENTRY>'`

`RESEND_API_KEY='<API Key from resend.com to send emails>'`

> You can get api key from resend.com after verifying your domain
