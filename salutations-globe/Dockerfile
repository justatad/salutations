FROM golang:1.9
RUN git clone https://github.com/justatad/salutations src
RUN go install salutations-globe
ENV PORT 8080
CMD ["./salutations-globe"]
