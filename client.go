package main

import (
	"context"
	"log"
	"time"

	pb "ls/" // ⚠️ Замени на свой путь (смотрите объяснение ниже)

	"google.golang.org/grpc"
)

func main() {
	// 1. Устанавливаем соединение с сервером
	conn, err := grpc.Dial(
		":50051",                        // Адрес сервера
		grpc.WithInsecure(),             // Отключаем TLS (для примера)
		grpc.WithBlock(),                // Ждём подключения перед продолжением
		grpc.WithTimeout(time.Second*3), // Таймаут подключения
	)
	if err != nil {
		log.Fatalf("Не удалось подключиться к серверу: %v", err)
	}
	defer conn.Close() // Закроем соединение при выходе из main()

	// 2. Создаем клиентский стаб
	client := pb.NewGreeterClient(conn)

	// 3. Создаем контекст с таймаутом
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// 4. Формируем запрос
	name := "Alice"
	request := &pb.HelloRequest{
		Example: name,
	}

	// 5. Вызываем удалённый метод SayHello
	response, err := client.SayHello(ctx, request)
	if err != nil {
		log.Fatalf("Ошибка при вызове SayHello: %v", err)
	}

	// 6. Выводим ответ от сервера
	log.Printf("Ответ от сервера: %s", response.GetMessage())
}
