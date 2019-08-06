FROM golang:1.12.7



WORKDIR /build
COPY src .
CMD ./build.sh
CMD ls


#WORKDIR /app
#COPY ./src /build/GOSICK

CMD ./GOSICK

