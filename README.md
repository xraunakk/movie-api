🎬 Movie Review API — Built with Go & Gin

A lightweight, fast, and beginner-friendly REST API for managing movie reviews — built using the Go programming language and Gin web framework.

This project is part of my backend portfolio to demonstrate CRUD operations, routing, data handling, and API testing with Postman. It is ideal for learning backend fundamentals and could be the foundation for a larger project like a Netflix-style review platform.

🚀 Features
✅ Create, Read, Update, Delete (CRUD) movie reviews
✅ In-memory and JSON file-based persistent storage
✅ Clean RESTful API design using Gin
✅ Tested with Postman
✅ JSON request and response handling
✅ Ready to connect to PostgreSQL via GORM (coming soon!)

🛠 Tech Stack
• Language: Go
• Framework: Gin
• Tools: Postman, Git, GitHub
• Storage: In-memory slice + JSON file (for now)

📁 Project Structure
movie-api/
│
├── main.go                 # Main application
├── reviews.json            # Optional - persistent storage file
├── go.mod / go.sum         # Go module files
└── README.md               # This file

🧪 How to Run Locally
1. Clone the repo
git clone https://github.com/xraunakk/movie-api.git
cd movie-api
2. Run the server
go run main.go
Server runs on:
http://localhost:8080

📬 API Endpoints
Method	Endpoint	Description
GET	/reviews	Get all reviews
POST	/reviews	Add a new review
PUT	/reviews/:id	Update review by ID
DELETE	/reviews/:id	Delete review by ID

📦 Sample JSON for POST or PUT
{
  "id": "4",
  "movie": "Inception",
  "review": "Amazing concept and visuals!",
  "rating": 5
}

📌 Future Improvements
🔐 Add JWT-based authentication
💾 Connect to PostgreSQL using GORM
🌐 Deploy on Render/Railway
🎨 Build frontend in React or HTML
🔍 Add filtering/search by movie title

🙋‍♂️ Author
Raunak Raj
Backend Developer in Progress | Passionate about Go, APIs, and building real-world projects.

📫 LinkedIn/GitHub: https://github.com/xraunakk/movie-api
📝 License
This project is open source and free to use under the MIT License.