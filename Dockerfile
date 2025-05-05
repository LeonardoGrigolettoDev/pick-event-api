FROM golang:1.23


WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -o main .

CMD ["sh", "-c", "if [ \"$ENV\" = \"dev\" ]; then watchmedo auto-restart --directory=./ --pattern=*.py --recursive -- python main.py; else python main.py; fi"]