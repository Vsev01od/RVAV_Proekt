Распределенный вычислитель арифметических выражений

Архитектура проекта:
Оркестратор(он же сервер), принимает на вход выражение
Агент - решает выражение, посылает его на сервер, и сервер выводит ответ

Инструкция по запуску РВАВ:
1. Должен быть установлен Golang
2. Должен быть устанолен Git
3. Скачать zip архив с проектом на свай компьютер

4. Далее создайте 2 терменала
5. Передите спомощью cd в папку Server и введите команду:
go run Server.go
6. Зайдите во второй терменал и войдите в папку Vvod и тамже введите команду:
go run Vvod.go
7. В этом же терменале введите выражение, например "2+2*2"
8. В терменале Server вас будет ждать ответ



Примеры выражений:
2+2*2(ответ: 6)
4+5-3*3(ответ: 0)
4+4*3(ответ: 16)
