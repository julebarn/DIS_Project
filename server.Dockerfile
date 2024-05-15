###############################
######### BUILD SQLC ##########
###############################



###############################
####### BUILD GO-SERVER #######
###############################

FROM golang:alpine AS server-builder
WORKDIR /src

COPY server .

# NOTE: If third-party deps are installed, then git has to be manually installed here.

RUN GOOS=linux go build -o server ./main.go

FROM scratch
EXPOSE 8080
COPY --from=server-builder /src/server /bin/server
CMD ["/bin/server"]