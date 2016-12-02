FROM alpine:3.4

RUN apk add --no-cache poppler-utils
COPY pdf2text /pdf2text
RUN chmod +x /pdf2text

CMD ["/pdf2text"]
