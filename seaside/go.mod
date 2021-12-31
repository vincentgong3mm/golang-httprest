module github.com/vincentgong3mm/golanghttprest/seaside

go 1.15

require (
	github.com/vincentgong3mm/golanghttprest/seaside/mongowrap v0.0.0
	go.mongodb.org/mongo-driver v1.8.1
)

replace github.com/vincentgong3mm/golanghttprest/seaside/mongowrap v0.0.0 => ./mongowrap
