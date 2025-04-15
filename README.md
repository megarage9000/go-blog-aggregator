# Go Blog Agggregator (Gator)

This project is an application that enables users to aggregate RSS feeds from the internet. It is named gator to be as a play-on to the word **aggreGATOR**, as the applications aggregates RSS feeds. This application can do the following:

- Append RSS feeds from the internet into a Postgres SQL Database
- Enables the registering of multiple users, each with their own list of RSS feeds
- Store posts from the RSS feeds to the Postgres SQL Database
- Lists posts that are followed by a given user
- Follow and unfollow an RSS Feed

It is a part of [Boot.dev's](https://www.boot.dev/lessons/14b7179b-ced3-4141-9fa5-e67dbc3e5242) Backend courses, and this course provides students hands-on experiences with:

- Writing SQL queries
- Designing Database Schemas
- Writing Go for http requests and ORMs with [SQLC](https://sqlc.dev/)
- Executing up and down migrations with [Goose](https://github.com/pressly/goose)
- Developing long running services

### Required Software

This applications requires the following:

- [Postgres SQL](https://www.postgresql.org/)
- [Go toolchain (1.23+)](https://go.dev/doc/install)
- [Goose](https://github.com/pressly/goose)

The project was run on WSL (Windows Subsystem for Linux). 

### Setup

1. Create a `.gatorconfig.json`. Store it in your home directory, i.e. `~`.
    * For this project, I stored mine in `~/.gatorconfig.json`.
    
    The config should look something like this:
    ```
    {
        "db_url":"[insert postgres sql link here]",
    }
    ```
    The `db_url` should be set something like: `protocol://username:password@host:port/database?sslmode=disable`. For this project, an example db_url can look like: `postgres://postgres:postgres@localhost:5432/gator?sslmode=disable`. *(Note sslmode is only disabled for local runs)*.
    
    The application will fill in the `current_user_name` .

    
2. Start the PostgreSQL server, and configure your PostgreSQL server however you prefer.

3. Create a database with any name.

4. In the `sql/schema` folder, run `goose postgres <connection string> up`, where the `<connection string>` is the `db_url` from earlier: `protocol://username:password@host:port/database`

5. Return to the root folder and execute `go build` and `go install`. The application will be called `go-blog-aggregator`.

### Running the Go Blog Aggregator

There are several commands that the `go-blog-aggregator` performs:

- `go-blog-aggregator login [name]`

    Sets the user to the given `name`. The `name` must exists in the database.

- `go-blog-aggregator register [name]`

    Creates a user with `name`, and sets the current user to `name`. The `name` will be added to the database.

- `go-blog-aggregator reset`

    Clears data from the database, including all users, feeds and posts.

- `go-blog-aggregator users`   

    Lists all the users on the database.

- `go-blog-aggregator feeds`

    Lists all the feed on the database, along with which user added them.

- `go-blog-aggregator addfeed [feedname] [feedurl]`

    Creates a new feed and adds it to the database. The current user will also automatically follow this feed.

- `go-blog-aggregator follow [feedurl]`

    Sets the current user to follow the given `feedurl`.

- `go-blog-aggregator following`

    Lists all the feeds that the current user is following.

- `go-blog-aggregator unfollow [feedurl]`

    Sets the current user to unfollow the given `feedurl`.

- `go-blog-aggregator browse (optional)[numberOfPosts]`

    Browse posts that the current user follows. the `numberOfPosts` is optional, and will default to `2`.

- `go-blog-aggregator agg [frequency]`

    Scrapes the feeds that the user follows for posts to save on the database. `frequency` should be parseable to [Duration](https://pkg.go.dev/time#ParseDuration) in Go.



### Planned Features to Add

- [ ] Concurrent aggregations 
- [ ] Paginated browsing
- [ ] HTTP API
- [ ] Basic Web Frontend 
- [ ] Host project as continuing service on the cloud