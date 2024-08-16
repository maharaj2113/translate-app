#Translate App
This Go application provides a REST API endpoint to translate text using the Google Translator API. The application is containerized using Docker and can be deployed to a Kubernetes cluster.

Features
Translate Text: Accepts POST requests with a source language, target language, and text to translate.
Google Translator API Integration: The application interfaces with the Google Translator API for translation.
Dockerized: The application can be built and run within a Docker container.
Kubernetes Deployment: Easily deploy the application to a Kubernetes cluster.


#Prerequisites
Go 1.20+
Docker
Kubernetes 

# Build and Run the Application Locally
If you want to run the application locally without Docker:
go run main.go


#Build the Docker Image:
docker build -t translate-app .

#Run the Docker Container:
docker run -d -p 8080:8080 translate-app

#Deploy to Kubernetes:
kubectl apply -f deployment.yaml

Once the application is running (either locally, in Docker, or in Kubernetes), you can send a POST request to the /translate endpoint to translate text.

Example Request
curl -X POST http://localhost:8080/translate \
-H "Content-Type: application/json" \
-d '{
    "q": "Hello, world!",
    "source": "en",
    "target": "es",
    "format": "text"
}'


Example Response
{
  "data": {
    "translations": [
      {
        "translatedText": "¡Hola, mundo!"
      }
    ]
  }
}
