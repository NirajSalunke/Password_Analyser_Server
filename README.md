# Password Analyser API

A Go-based HTTP API that provides real-time password analysis, scoring, and feedback to help users create strong and secure passwords.

---

## ✨ What It Does

### 🔐 Password Strength Analysis

* Evaluates passwords based on length, complexity, and entropy to assign a security score.

### 🔖 Feedback for Improvement

* Offers actionable suggestions to improve password strength, such as increasing length or adding special characters.

### 🔧 RESTful Endpoint

* **Endpoint:** `/api/v1/analyse`
* **Request Example:**

  ```json
  {
    "password": "P@ssw0rd123"
  }
  ```
* **Response Example:**

  ```json
  {
    "score": 4,
    "feedback": {
      "strength": "Strong",
      "suggestions": []
    }
  }
  ```

### ⚖️ Customizable Analysis

* Configurable thresholds for password scoring and customizable feedback messages.

---

## 🔧 What It Uses

### 💀 Go (1.19+)

* High-performance backend server.

### 🌐 Gin Web Framework

* Efficient routing and middleware support.

### ⚖️ Modular Codebase

* **`controllers/`**: Handles HTTP requests for password analysis.
* **`helpers/`**: Includes functions for password scoring and feedback generation.
* **`models/`**: Data models for request and response structures.
* **`routes/`**: API route definitions.

### 🔠 YAML Configuration

* Flexible settings for password scoring criteria, thresholds, and logging options.

---

## 🚀 Use Cases

### 🔐 User Account Security

* Integrate into registration flows to ensure users create strong passwords.

### 🔒 Authentication Platforms

* Provide real-time feedback during password creation to encourage better security practices.

### 🔨 Educational Tools

* Help users understand password security and improve awareness of strong password practices.

### 🔑 Enterprise Security Audits

* Assess existing password policies and enforce strength requirements across systems.

---

## 🔖 License

Distributed under the **MIT License**.
