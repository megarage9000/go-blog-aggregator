module github.com/megarage9000/go-blog-aggregator

go 1.23.4

require internal/config v0.0.0

require (
	github.com/google/uuid v1.6.0 // indirect
	github.com/lib/pq v1.10.9 // indirect
)

replace internal/config => ./internal/config
