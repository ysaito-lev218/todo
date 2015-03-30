package todo

type Todo struct {
	ID       int64  `meddler:"id,pk"`
	Title    string `meddler:"title"`
	Finished bool   `meddler:"finished"`
}