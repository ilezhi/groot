FROM alpine

COPY groot /app/groot

EXPOSE 9000

CMD ["/app/groot"]
