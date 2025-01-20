#!/bin/bash

# 1. Будуємо Docker образ
docker build -t cert-generator .

# 2. Запускаємо контейнер
container_id=$(docker run -d --name cert-container cert-generator)

# 3. Чекаємо деякий час для завершення генерації сертифікатів
echo "Чекаємо завершення генерації сертифікатів..."
sleep 10  # Затримка, щоб дати контейнеру час для виконання генерації сертифікатів

# 4. Копіюємо сертифікати з контейнера на хост
echo "Копіюємо сертифікати з контейнера..."
docker cp "$container_id:/certs/server.crt" ./server.crt
docker cp "$container_id:/certs/server.key" ./server.key

# 5. Перевіряємо, чи файли успішно скопійовані
if [ -f "./server.crt" ] && [ -f "./server.key" ]; then
    echo "Сертифікати успішно скопійовані!"
else
    echo "Помилка під час копіювання сертифікатів."
fi

# 6. Видаляємо контейнер після завершення
docker rm -f "$container_id"

# 7. Завершуємо скрипт
echo "Процес завершено!"
