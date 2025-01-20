@echo off

echo Створюємо Docker образ...
docker build -t cert-generator .

echo Запускаємо контейнер...
for /f "delims=" %%i in ('docker run -d --name cert-container cert-generator') do set container_id=%%i

echo Чекаємо завершення генерації сертифікатів...
timeout /t 15

echo Копіюємо сертифікати з контейнера...
docker cp %container_id%:/certs/server.crt .\server.crt
docker cp %container_id%:/certs/server.key .\server.key

if exist ".\server.crt" if exist ".\server.key" (
    echo Сертифікати успішно скопійовані!
) else (
    echo Помилка під час копіювання сертифікатів.
)

docker rm -f %container_id%

echo Процес завершено!
pause
