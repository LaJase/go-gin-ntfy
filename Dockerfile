# Utilisez une image Alpine comme base
FROM golang:alpine

# Installez les dépendances nécessaires
RUN apk add --no-cache git

# Copiez le code de votre projet dans le conteneur
WORKDIR /app
COPY . .

# Construisez l'application Go
RUN go build -o myrestapp ./main.go

# Exposez le port sur lequel l'application écoute
EXPOSE 8080

# Démarrez l'application au lancement du conteneur
CMD ["./myrestapp"]

