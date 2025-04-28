# Gator

Gator is a multi-user CLI application for aggregating and browsing RSS feeds. It's designed for single-device use but supports multiple user profiles.
**Note**: This application does not include user-based authorization. Anyone with database credentials can act as any user.

## Prerequisites

- Go 1.23 or later
- PostgreSQL database
- Unix-like environment

## Installation

Clone the repository:
```
go install https://github.com/ewielguszewski/gator@latest
```

After installation, the `gator` command will be available in your terminal, assuming your Go bin directory is in your PATH.

## Config

Create a `.gatorconfig.json` file in your home directory with the following structure:

```json
{
    "db_url": "postgres://username:@localhost:5432/database?sslmode=disable"
}
```

Replace the values with your PostgreSQL database connection details. Make sure the database exists before running the application.

## Usage

### User management

- `gator register <name>` - Create a new user
- `gator login <name>` - Switch to an existing user profile
- `gator users` - List all users

### Feed management

- `gator addfeed <url>` - Add a new RSS feed to the database
- `gator feeds` - List all feeds in the database
- `gator follow <url>` - Subscribe to a feed that already exists in the database
- `gator unfollow <url>` - Unsubscribe from a feed

### Content

- `gator agg <time_between_reps>` - Start the aggregator to collect new posts (e.g., gator agg 30m for every 30 minutes). This needs to run continuously to update feeds.
- `gator browse <limit>` - View the latest posts, where <limit> is the number of posts to display



