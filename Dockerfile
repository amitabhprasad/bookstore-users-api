## We specify the base image we need for our
## go application
FROM golang:1.17-alpine

ENV REPO_URL=github.com/amitabhprasad/bookstore-app/bookstore-users-api/src

ENV GOPATH=/app

ENV APP_PATH=${GOPATH}/src/${REPO_URL}

ENV WORKPATH=$APP_PATH/src

COPY src ${WORKPATH}

WORKDIR $WORKPATH
RUN go build -o user-api .

# expose port 8081
EXPOSE 8081

CMD ["./user-api"]