echo "build image"
docker build -t praktikum/hello-world .

echo "create container"
docker run -d --name app -p 3000:80 praktikum/hello-world

echo "push image"
docker push praktikum/hello-world