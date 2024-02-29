# Setup 🔧

![944960 512 (1)](https://github.com/shuklarituparn/Conversion-Microservice/assets/66947051/b0db7aef-2982-416c-96bc-7f877a6e9ce7)

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

* If you dont' have NVIDIA graphics, then remove the parameter  `h264_nvenc` from the code at `internal/ffmpeg/conversion.go`

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

