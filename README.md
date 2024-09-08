Amazon Price Tracker
====================

A simple Go-based Amazon Price Tracker where users can add Amazon product IDs. The system automatically tracks the product's price, and when the price drops, the user will receive an email notification.

Features
--------

-   Add Amazon product ID for tracking.
-   Automatically fetch and track product prices at regular intervals.
-   Sends an email notification to the user when the price drops below the initial price.

Requirements
------------

-   Go 1.18+
-   MySQL (or your preferred database)
-   An email service (SMTP)

Getting Started
---------------

### 1\. Clone the repository

```bash
git clone https://github.com/notgman/go-amazon-price-tracker.git
cd go-amazon-price-tracker
```

### 2\. Install dependencies

Make sure you have Go installed. Then, run the following command to install required Go packages:

```bash
go mod tidy
```

### 3\. Setup Environment Variables

Create a `.env` file in the root of your project with the following content:

```bash
DB_USER="your_db_username"
DB_PASS="your_db_password"
DB_NAME="AmazonPrice"
SENDER="your_email@example.com"
APP_PASSWORD="your_email_password_or_app_password"
```

These environment variables are required to:

-   Connect to the MySQL database (`DB_USER`, `DB_PASS`, `DB_NAME`).
-   Send email notifications (`SENDER`, `APP_PASSWORD`).

### 4\. Setup Database

Create a MySQL database for tracking products:

```bash
CREATE DATABASE AmazonPrice;
```

Make sure to match the database name in the `.env` file.

### 5\. Database Schema

Run the application once to create the necessary tables in the database (`products` and `prices`).

-   **products table**: Stores product information (ID, name, URL, and user email).
-   **prices table**: Stores price information and is linked to the `products` table.

### 6\. Running the Application

Once everything is set up, you can run the application with:

#### Add product to track

```bash 
go run main.go add
```

This command is used to add a new product to track.

#### Get all tracked products

```bash
go run main.go get
```

This command fetches and displays all tracked products.

#### Update product prices

```bash
go run main.go update
```

The system will track prices and send email notifications when price drops are detected.

### 7\. Cron Job Setup (Optional)

To continuously track product prices, you can set up a cron job or a scheduled task to run the price check at regular intervals.

Usage
-----

-   **Add a product**: Add the Amazon product ID that you want to track, and the system will start tracking the product's price.
-   **Price drop notification**: Once the system detects a price drop below the initial price, an email will be sent to the user.

Contributing
------------

Feel free to open issues or create pull requests to contribute to the project.

License
-------

This project is licensed under the GNU Public License - see the LICENSE file for details.

* * * * *