# South Park Messaging System

A distributed microservice system where South Park characters send and receive messages asynchronously.  
Built using **Go (Hexagonal Architecture)**, **RabbitMQ**, and **Python**.

---

## 🧱 Project Structure

southpark/

├── go-api/

    │ ├── app/

    │ ├── domain/

    │ ├── ports/

    │ ├── adapters/

    │ └── main.go

├── consumer/

    │ ├── main.py

    │ └── Dockerfile

├── index.html

├── docker-compose.yml

└── README.md


---

## 🚀 Features

✅ Go HTTP API (`/messages` endpoint) to send messages  
✅ RabbitMQ message broker for async communication  
✅ Python consumer that listens and prints messages  
✅ Clean **Hexagonal Architecture (Ports & Adapters)**  
✅ Fully Dockerized and runnable via one command  
✅ Interactive web interface (`index.html`) with character selection and auto-send  
✅ CORS-enabled API for frontend access  

---

## ⚙️ Setup and Run

### Clone and build
```bash
git clone https://github.com/Anthony-Shoshi/distributed-microservice-application

cd southpark

docker compose up --build
```

### Verify services

- Go API → http://localhost:8080

- RabbitMQ Dashboard → http://localhost:15672 (user: guest, password: guest)

### Use the web interface

1. Open `index.html` in your browser (or serve it via a local server)
2. Select a South Park character (Cartman, Kyle, Stan, Kenny, Butters, Randy)
3. Click **"Send Now"** to send a single message, or **"Start Auto-Send"** to send messages at regular intervals
4. Watch the live log for confirmation and check the consumer output

### Send a message via curl
```bash 
curl -X POST http://localhost:8080/messages \
  -H "Content-Type: application/json" \
  -d '{"author": "Cartman", "body": "Respect my authoritah!"}'
  ```

  ### Check consumer logs
  
  You’ll see:

  ``` Cartman says: Respect my authoritah! ```
