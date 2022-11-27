WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go get github.com/titoyudha/golang_todo_list/config
RUN go get github.com/titoyudha/golang_todo_list/controller
RUN go get github.com/titoyudha/golang_todo_list/model

EXPOSE 8080

CMD [ "/todo-list" ]
