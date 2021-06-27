#  \
#  \\,
#   \\\,^,.,,.                    “Zero to Hero”
#   ,;7~((\))`;;,,               <zerotohero.dev>
#   ,(@') ;)`))\;;',    stay up to date, be curious: learn
#    )  . ),((  ))\;,
#   /;`,,/7),)) )) )\,,
#  (& )`   (,((,((;( ))\,

FROM golang:1.16.4-alpine AS builder

# `git` is required to fetch go dependencies:
RUN apk add --no-cache ca-certificates git

# Create the user and group files that will be used in the running
# container to run the process as an unprivileged user:
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Copy the predefined `.netrc` file into the location that git depends on:
COPY ./.netrc /root/.netrc
RUN chmod 600 /root/.netrc

# Set the working directory outside `$GOPATH` to enable the support for modules:
WORKDIR /src

# Fetch dependencies first; they are less susceptible to change on every build
# and will therefore be cached for speeding up the next build:
COPY ./go.mod ./go.sum ./
RUN go mod download

# Import the code from the context:
COPY . .

# Build the executable to `/app`. Mark the build as statically linked:
RUN CGO_ENABLED=0 go build -installsuffix 'static' -o /app ./cmd/main.go

# This is the final “minimal” container image:
FROM scratch AS final

# Import the user and group files from the build container:
COPY --from=builder /user/group /user/passwd /etc/

# Import the Certificate-Authority certificates for enabling HTTPS:
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Import the compiled executable from the build container:
COPY --from=builder /app /app

# Perform any further action as an unprivileged user:
USER nobody:nobody

# Start listening (should match FIZZ_CRYPTO_SVC_PORT):
EXPOSE 9001/tcp

# Run the compiled binary:
ENTRYPOINT ["/app"]