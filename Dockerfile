FROM alpine:3.5

COPY ./httpgo /app/httpgo
RUN chmod +x /app/httpgo
ENV PORT 8080
EXPOSE 8080

RUN addgroup -S -g 10000 app \
    && adduser -S -D -u 10000 -s /sbin/nologin -h /app -G app app 
USER 10000

ENTRYPOINT /app/httpgo