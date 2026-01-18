# Smart Inventory Management System Backend

## GENERAL DESCRIPTION
This project is a comprehensive backend system designed to help stores and warehouses manage their inventory intelligently based on data. The system provides robust management for users, stores, warehouses, product categories, and goods. It features intelligent stock tracking (Stock In/Out), automatic stock updates, and calculated metrics.

A key differentiator of this system is its integration with **Google Gemini AI** to provide actionable insights. The system analyzes transaction patterns to offer recommendations on restocking and usage, helping business owners make data-driven decisions. History of these analyses is stored for tracking and future reference.

## PROBLEM IDENTIFICATION AND GOAL

### Problem Identification
Manual inventory management is prone to human error, leading to inaccurate stock records, overstocking, or stockouts. Traditional systems often lack predictive capabilities, making it difficult for store owners to anticipate demand and reorder at the optimal time, resulting in lost sales or tied-up capital.

### Goal
The goal of this system is to automate inventory control and empower store owners with an AI-driven decision support system. By digitizing the entire flow from store management to stock transactions and leveraging AI for insights (such as restocking needs), the system aims to optimize stock levels, reduce waste, and ensure product availability.

## FUNCTIONAL REQUIREMENTS ANALYSIS

| Generic Function | Details |
| :--- | :--- |
| **User Management** | Register, Login, Update Profile, Delete Account. (One User per Store implementation). |
| **Store Management** | Create (on register), Update details. |
| **Warehouse Management** | Create multiple warehouses per store, Update, Delete, View details. |
| **Category Management** | Create, Read, Update, Delete product categories. |
| **Goods (Stock) Management** | Add items, Upload images (Cloudinary), Update details, Manage safety stock levels. |
| **Transaction Recording** | Record Stock In/Out. Trigger AI analysis on transaction. |
| **AI Insights** | Generate and Store insights (Smart Log) using Google Gemini based on transaction context. |

## DIRECTORY STRUCTURE

```text
.
├── app
│   ├── bootstrap
│   │   └── index.go                  # Entry point for app initialization
│   ├── config
│   │   ├── cloudinary_config
│   │   │   └── index.go              # Cloudinary configuration
│   │   ├── db_config
│   │   │   └── index.go              # Database connection configuration
│   │   ├── gemini_config
│   │   │   └── index.go              # Gemini AI configuration
│   │   ├── jwt_config
│   │   │   └── index.go              # JWT configuration
│   │   ├── port_config
│   │   │   └── index.go              # Port configuration
│   │   └── index.go                  # Main config loader
│   ├── controller
│   │   ├── auth_controller
│   │   │   └── index.go              # Auth handlers (Register, Login)
│   │   ├── barang_controller
│   │   │   └── index.go              # Barang handlers (CRUD)
│   │   ├── category_controller
│   │   │   └── index.go              # Category handlers
│   │   ├── gudang_controller
│   │   │   └── index.go              # Gudang handlers
│   │   ├── toko_controller
│   │   │   └── index.go              # Toko handlers
│   │   ├── trx_controller
│   │   │   └── index.go              # Transaction handlers
│   │   └── user_controller
│   │       └── index.go              # User profile handlers
│   ├── database
│   │   └── index.go                  # GORM Connection & Migration
│   ├── helper
│   │   ├── cloudinary.go             # Cloudinary upload helper
│   │   ├── error.go                  # Error handling helper
│   │   ├── gemini.go                 # Google Gemini AI helper
│   │   ├── hash.go                   # Password hashing helper
│   │   └── jwt.go                    # JWT token helper
│   ├── interface
│   │   ├── repository
│   │   │   ├── auth_repository_interface/index.go
│   │   │   ├── barang_repository_interface/index.go
│   │   │   ├── category_repository_interface/index.go
│   │   │   ├── gudang_repository_interface/index.go
│   │   │   ├── smartlog_repository_interface/index.go
│   │   │   ├── toko_repository_interface/index.go
│   │   │   ├── trx_repository_interface/index.go
│   │   │   └── user_repository_interface/index.go
│   │   └── service
│   │       ├── auth_service_interface/index.go
│   │       ├── barang_service_interface/index.go
│   │       ├── category_service_interface/index.go
│   │       ├── gudang_service_interface/index.go
│   │       ├── toko_service_interface/index.go
│   │       ├── trx_service_interface/index.go
│   │       └── user_service_interface/index.go
│   ├── middleware
│   │   └── jwt_middleware.go         # Authentication middleware
│   ├── models
│   │   ├── barang.go                 # Barang entity definition
│   │   ├── category.go               # Category entity definition
│   │   ├── gudang.go                 # Gudang entity definition
│   │   ├── smart_log.go              # SmartLog entity definition
│   │   ├── toko.go                   # Toko entity definition
│   │   ├── trx_log.go                # TrxLog entity definition
│   │   └── user.go                   # User entity definition
│   ├── registry
│   │   ├── auth_registry/index.go    # Dependency injection for Auth
│   │   ├── barang_registry/index.go  # Dependency injection for Barang
│   │   ├── category_registry/index.go
│   │   ├── gudang_registry/index.go
│   │   ├── toko_registry/index.go
│   │   ├── trx_registry/index.go
│   │   └── user_registry/index.go
│   ├── repository
│   │   ├── auth_repository/index.go      # DB logic for Auth
│   │   ├── barang_repository/index.go    # DB logic for Barang
│   │   ├── category_repository/index.go
│   │   ├── gudang_repository/index.go
│   │   ├── smartlog_repository/index.go
│   │   ├── toko_repository/index.go
│   │   ├── trx_repository/index.go
│   │   └── user_repository/index.go
│   ├── request
│   │   ├── auth_request/index.go     # Request structs for Auth
│   │   ├── barang_request/index.go   # Request structs for Barang
│   │   ├── category_request/index.go
│   │   ├── gudang_request/index.go
│   │   ├── toko_request/index.go
│   │   ├── trx_request/index.go
│   │   └── user_request/index.go
│   ├── router
│   │   ├── auth_router/index.go      # Route definitions for Auth
│   │   ├── barang_router/index.go    # Route definitions for Barang
│   │   ├── category_router/index.go
│   │   ├── gudang_router/index.go
│   │   ├── toko_router/index.go
│   │   ├── trx_router/index.go
│   │   └── user_router/index.go
│   └── service
│       ├── auth_service/index.go     # Business logic for Auth
│       ├── barang_service/index.go   # Business logic for Barang
│       ├── category_service/index.go
│       ├── gudang_service/index.go
│       ├── toko_service/index.go
│       ├── trx_service/index.go
│       └── user_service/index.go
├── postgres
│   └── INIT.sql                      # SQL initialization script
├── docker-compose.yml                # Docker Compose config
├── Dockerfile                        # Docker build config
├── go.mod                            # Go dependencies
├── main.go                           # Main application entry point
└── README.md                         # Project documentation
```

## ENTITY / MODEL

The system uses a Relational Database with the following entities:

### 1. User
Represents the shop owner.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the user. |
| 2 | `name` | Varchar(225) | Full name of the user. |
| 3 | `email` | Varchar(225) | Email address used for login. |
| 4 | `password` | Varchar(225) | Encrypted password string. |
| 5 | `no_telp` | Varchar(225) | Contact phone number. |
| 6 | `created_at` | Date | Timestamp when the user was created. |
| 7 | `updated_at` | Date | Timestamp when the user was last updated. |

**Relationship**: A **User** has exactly **one Toko** (One-to-One).

### 2. Toko (Store)
Represents the business entity managed by a user.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the store. |
| 2 | `user_id` | UUID (FK) | Reference to the owner (User). |
| 3 | `name` | Varchar(225) | Name of the store. |
| 4 | `address` | Varchar(225) | Physical address of the store. |
| 5 | `created_at` | Date | Timestamp when the store was created. |
| 6 | `updated_at` | Date | Timestamp when the store was last updated. |

**Relationship**: A **Toko** belongs to a **User**. A **Toko** has many **Gudang**, **Category**, and **Barang**.

### 3. Gudang (Warehouse)
Physical storage locations for inventory.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the warehouse. |
| 2 | `toko_id` | UUID (FK) | Reference to the parent Store. |
| 3 | `name` | Varchar(225) | Name of the warehouse (e.g., "Main Warehouse"). |
| 4 | `address` | Varchar(225) | Location of the warehouse. |
| 5 | `created_at` | Date | Timestamp when the warehouse was created. |
| 6 | `updated_at` | Date | Timestamp when the warehouse was last updated. |

**Relationship**: A **Gudang** belongs to a **Toko**. A **Gudang** contains many **Barang** and records many **Trx_Log** and **Smart_Log**.

### 4. Category
Classification for goods.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the category. |
| 2 | `toko_id` | UUID (FK) | Reference to the parent Store. |
| 3 | `name` | Varchar(225) | Name of the category (e.g., "Electronics"). |
| 4 | `description` | Varchar(225) | Brief description of the category. |
| 5 | `created_at` | Date | Timestamp when the category was created. |
| 6 | `updated_at` | Date | Timestamp when the category was last updated. |

**Relationship**: A **Category** belongs to a **Toko**. A **Category** can have many **Barang**.

### 5. Barang (Goods)
Inventory items stored in warehouses.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the item. |
| 2 | `gudang_id` | UUID (FK) | Reference to the Warehouse where it is stored. |
| 3 | `category_id` | UUID (FK) | Reference to the Category it belongs to. |
| 4 | `name` | Varchar(225) | Name of the product. |
| 5 | `sku` | Varchar(225) | Stock Keeping Unit (Unique code). |
| 6 | `image_url` | Varchar(225) | URL to the product image (Cloudinary). |
| 7 | `stock` | Int | Current quantity available. |
| 8 | `safety_stock` | Int | Minimum stock level before restock is needed. |
| 9 | `need_restock` | Boolean | Flag indicating if stock is low (True/False). |
| 10 | `lead_time_days` | Int | Days required to restock the item. |
| 11 | `created_at` | Date | Timestamp when the item was added. |
| 12 | `updated_at` | Date | Timestamp when the item was last updated. |

**Relationship**: A **Barang** belongs to a **Gudang** and a **Category**. It relates to many **Trx_Log** and **Smart_Log** entries.

### 6. Trx_Log (Transaction Log)
Record of stock movements.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the transaction. |
| 2 | `barang_id` | UUID (FK) | Reference to the item moved. |
| 3 | `gudang_id` | UUID (FK) | Reference to the warehouse involved. |
| 4 | `qty` | Int | Quantity moved. |
| 5 | `type` | Enum("in"/"out") | Type of movement: 'in' (Stock In) or 'out' (Stock Out). |
| 6 | `created_at` | Date | Timestamp when transaction occurred. |
| 7 | `updated_at` | Date | Timestamp when transaction was last modified. |

**Relationship**: A **Trx_Log** links a specific **Barang** within a **Gudang**.

### 7. Smart_Log (AI Insight)
Stored AI analysis results for inventory intelligence.

| NO | COLUMN | DATA TYPES | EXPLANATION |
|:---|:---|:---|:---|
| 1 | `id` | UUID (PK) | Unique identifier for the log. |
| 2 | `barang_id` | UUID (FK) | Reference to the analyzed item. |
| 3 | `gudang_id` | UUID (FK) | Reference to the warehouse of the item. |
| 4 | `period_month` | Int | The month of the analysis period. |
| 5 | `period_year` | Int | The year of the analysis period. |
| 6 | `ai_insight` | Varchar(225) | Text content of the AI's recommendation/analysis. |
| 7 | `created_at` | Date | Timestamp when the log was generated. |
| 8 | `updated_at` | Date | Timestamp when the log was last updated. |

**Relationship**: A **Smart_Log** provides insights for a specific **Barang** in a **Gudang**.

## FEATURES, HELPERS, AND MIDDLEWARE USED IN EACH API

The system uses a Clean Architecture consisting of Controllers, Services, Repositories, and Registries.

**Global Helpers & Middleware:**
- **JWT Middleware**: Protections for private routes.
- **Error Helper**: Standardized error response handling.
- **Hash Helper**: Secure password hashing.
- **Validator**: Request payload validation.

### 1. Auth API
Manages user registration and limits.

**Features Used**: `Hash Helper`, `JWT Helper`.

#### Register User & Store
- **Endpoint**: `POST /api/auth/register`
- **Description**: Registers a new user and automatically creates a Store (Toko) for them.
- **Request Body**:
```json
{
  "name": "John Doe",
  "email": "john@example.com",
  "password": "securepassword123",
  "no_telp": "08123456789",
  "toko_name": "John's Electronics"
}
```
- **Response** (201 Created):
```json
{
  "Message": "Success Register User",
  "Data": {
    "ID": "user-uuid...",
    "name": "John Doe",
    "email": "john@example.com",
    "no_telp": "08123456789",
    "Toko": {
        "ID": "toko-uuid...",
        "name": "John's Electronics",
        ...
    }
  }
}
```

#### Login
- **Endpoint**: `POST /api/auth/login`
- **Description**: Authenticates a user and returns a Bearer Token.
- **Request Body**:
```json
{
  "email": "john@example.com",
  "password": "securepassword123"
}
```
- **Response** (200 OK):
```json
{
  "Message": "Success Login",
  "Token": "eyJhGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

---

### 2. User Profile API
**Features Used**: `JWT Middleware`.

#### Get Profile
- **Endpoint**: `GET /api/profile/`
- **Headers**: `Authorization: Bearer <token>`
- **Response** (200 OK):
```json
{
  "Data": {
    "ID": "user-uuid...",
    "name": "John Doe",
    "email": "john@example.com",
    ...
  }
}
```

#### Update Profile
- **Endpoint**: `PUT /api/profile/update`
- **Request Body**:
```json
{
  "name": "John Doe Updated",
  "email": "john.new@example.com",
  "password": "newpassword123",
  "no_telp": "08199999999"
}
```
- **Response** (200 OK):
```json
{
  "Message": "Success Update User",
  "Data": { ...updated_user_data }
}
```

---

### 3. Store (Toko) API
**Features Used**: `JWT Middleware`.

#### Get Store Details
- **Endpoint**: `GET /api/toko/`
- **Response** (200 OK):
```json
{
  "Data": {
    "ID": "toko-uuid...",
    "user_id": "user-uuid...",
    "name": "John's Electronics",
    "address": "Downtown St."
  }
}
```

#### Update Store
- **Endpoint**: `PUT /api/toko/update/:id`
- **Request Body**:
```json
{
  "name": "John's Mega Store",
  "address": "New Location Ave."
}
```

---

### 4. Warehouse (Gudang) API
**Features Used**: `JWT Middleware`.

#### Create Warehouse
- **Endpoint**: `POST /api/gudang/create`
- **Request Body**:
```json
{
  "name": "Main Warehouse",
  "address": "Industrial Park Zone A"
}
```
- **Response** (201 Created):
```json
{
  "Message": "Success Create Gudang",
  "Data": {
    "ID": "gudang-uuid...",
    "name": "Main Warehouse",
    ...
  }
}
```

#### Get All Warehouses
- **Endpoint**: `GET /api/gudang/`
- **Response** (200 OK):
```json
{
  "Data": [
    { "ID": "1", "name": "Warehouse A" },
    { "ID": "2", "name": "Warehouse B" }
  ]
}
```

---

### 5. Category API
**Features Used**: `JWT Middleware`.

#### Create Category
- **Endpoint**: `POST /api/category/create`
- **Request Body**:
```json
{
  "name": "Smartphones",
  "description": "Mobile phones and accessories"
}
```

#### Get Categories
- **Endpoint**: `GET /api/category/`
- **Response** (200 OK):
```json
{
  "Data": [
    { "ID": "cat-1", "name": "Smartphones", "description": "..." }
  ]
}
```

---

### 6. Goods (Barang) API
**Features Used**: `JWT Middleware`, `Cloudinary Helper`.

#### Create Goods (Barang)
- **Endpoint**: `POST /api/barang/:id_gudang/create`
- **Content-Type**: `multipart/form-data`
- **Form Data**:
    - `name`: "iPhone 15 Pro"
    - `sku`: "IP15-PRO-BLK"
    - `stock`: 100
    - `safety_stock`: 10
    - `lead_time_days`: 3
    - `category_id`: "category-uuid..."
    - `image_url`: [File Upload]
- **Response** (201 Created):
```json
{
  "Message": "Success Create Barang",
  "Data": {
    "ID": "barang-uuid...",
    "name": "iPhone 15 Pro",
    "image_url": "https://res.cloudinary.com/...",
    "stock": 100,
    ...
  }
}
```

#### Get Goods List
- **Endpoint**: `GET /api/barang/:id_gudang`
- **Response** (200 OK):
```json
{
  "Data": [ ...list of items ]
}
```

---

### 7. Transaction (Trx) API
**Features Used**: `JWT Middleware`, `Gemini AI Helper`.

#### Create Transaction (Stock In/Out)
- **Endpoint**: `POST /api/trx/:id_gudang/barang/:id_barang/create`
- **Description**: Records stock movement. **Triggers Gemini AI Analysis** synchronously.
- **Request Body**:
```json
{
  "qty": 5,
  "type": "out" 
}
```
*Note: `type` can be "in" (restock) or "out" (sale/usage).*

- **Response** (201 Created):
```json
{
  "Message": "Success Create Trx Log",
  "Data": {
    "ID": "trx-uuid...",
    "qty": 5,
    "type": "out",
    "Barang": { ... },
    "Gudang": { ... }
  }
}
```
*Internal Effect*: A `SmartLog` entry is created in the database containing the AI's insight regarding this transaction.

#### Get Transaction History
- **Endpoint**: `GET /api/trx/:id_gudang/barang/:id_barang/detail`
- **Response** (200 OK):
```json
{
  "Data": [
    {
       "ID": "trx-1",
       "qty": 5,
       "type": "out",
       "created_at": "2024-01-01T10:00:00Z"
    },
    ...
  ]
}
```

---

## Beyond the Code — Trade-off Analysis

### Trade-off Analysis

| Decision | Trade-off | Rationale |
| :--- | :--- | :--- |
| **Monolithic Architecture** | **Pros**: Simple deployment, easier to develop initially. <br> **Cons**: Scaling specific components (like AI processing) requires scaling the whole app. | Suitable for the current scale and bootcamp scope. easy to manage state and database connections. |
| **PostgreSQL Database** | **Pros**: Reliable, ACID compliance for transactions, strong relational data integrity for inventory. <br> **Cons**: Slightly more setup than SQLite/NoSQL. | Essential for financial/stock data where consistency is critical. |

## CONCLUSION

The Smart Inventory Management System Backend is conceived as a comprehensive, data-driven solution that automates and optimizes retail inventory control. It provides robust modules for managing users, stores, warehouses, product categories, and goods, enabling seamless administration across the entire supply chain with real-time stock-in/out tracking. Its innovation lies in the integration of Google’s Gemini AI: each transaction triggers on-the-fly analysis of stock movement patterns, generating actionable recommendations for restocking and usage that help business owners make data-driven decisions. These AI-generated insights are recorded in “SmartLog” entries so that historical analysis is retained for future reference. By digitizing the entire workflow from user registration to granular inventory transactions and leveraging Gemini for predictive intelligence, the platform directly addresses the well-known drawbacks of manual inventory management (error-proneness, overstocking, stockouts). The net result is a decision-support framework that optimizes stock levels, reduces waste, and ensures product availability – aligning technological innovation with clear business value.

Architecturally, the backend is built on a clean, modular design with well-defined layers of controllers, services, and repositories. This layered, monolithic deployment simplifies development and state management, although it also means that scaling a specific component (such as the AI analysis service) would require scaling the entire application. Every stock transaction is durably logged in a PostgreSQL database, and a corresponding SmartLog record is created to capture the AI insight. This commitment to a relational (ACID-compliant) datastore reflects a systems-thinking emphasis on data integrity and reliability for business-critical inventory data. These thoughtful trade-off analyses – from choosing monolithic simplicity for the bootcamp scope to favoring PostgreSQL over lightweight alternatives – underline the project’s rigor in balancing technical complexity with deployability and business value. In sum, the design demonstrates a critical, systems-level approach: uniting clean architecture and AI-powered analytics to deliver innovative functionality today while carefully planning for long-term scalability and robustness.

