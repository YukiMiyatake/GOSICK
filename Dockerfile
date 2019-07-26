FROM golang:1.12.7



WORKDIR /build
COPY src .
CMD ./build.sh
CMD ls


WORKDIR /app
COPY /build/GOSICK .

CMD rmdir /build
CMD [GOSICK]

