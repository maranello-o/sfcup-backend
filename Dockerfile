FROM alpine:3.17

RUN mkdir /app

# RUN mkdir /app/static

WORKDIR /app

COPY ./yyzgpt .

# COPY ./static/* ./static/

RUN apk --update add tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata && \
    rm -rf /var/cache/apk/*

RUN chmod +x /app/sfcup

EXPOSE 8099

CMD ["./sfcup"]
