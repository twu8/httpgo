FROM alpine:3.5

COPY ./httpgo /app/httpgo
RUN chmod +x /app/httpgo
ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/httpgo