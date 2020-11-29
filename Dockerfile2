FROM golang:1.15.5-alpine

ENV SOURCES /go/src/github.com/PacktPublishing/httpgo

COPY . ${SOURCES}
# COPY ./httpgo /app/httpgo

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

# RUN chmod +x /app/httpgo

ENV PORT 8086
EXPOSE 8086

ENTRYPOINT httpgo