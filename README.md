# Практическая работа № 9 (Программирование корпоративных систем)

выполнила: **Полошкова Анастасия Юрьевна**

группа: **ЭФБО-01-22**


## Описание и этапы выполнения работы

В данной работе необходимо было реализовать взаимодействие клиента (приложения) с сервером на Go через REST API. Осуществить хранение данных на локальном сервере, а также управление ими используя методы GET, POST, PUT, DELETE.

1) Написала обработчики запросов в файле **main.go** ветка ***backend*** (https://github.com/jenkinsyoung/pks_pr_9/tree/backend)

В данной работе я использовала хранение двух структур: **Product** (информация о товаре) и **Item** (элемент корзины).

```
// Product представляет продукт
type Product struct {
	ID          int
	ImageURL    string
	Title       string
	Description string
	Rules       string
	Price       int
	Age         int
	Gamers      string
	Time        string
	Indicator   int
	IsFavorite  bool
}

// Item представляет элемент корзины
type Item struct {
	ID      int
	Counter int
}
```

Изначально корзина пуста. Содержимое каталога игр хранится в коде программы, без подключения БД

Запросы:
- "http://127.0.0.1:8080/products"  // Получить все продукты
-	"http://127.0.0.1:8080/products/create" // Создать продукт
-	"http://127.0.0.1:8080/products/" // Получить продукт по ID
-	"http://127.0.0.1:8080/products/delete/" // Удалить продукт
-	"http://127.0.0.1:8080/products/update/status/" // Обновление статуса избранного по ID
-	"http://127.0.0.1:8080/products/update/info/" // Обновление информации о товаре при редактировании по ID
-	"http://127.0.0.1:8080/basket"  // Получить все элементы корзины
-	"http://127.0.0.1:8080/basket/" // Проверить есть ли товар в корзине
-	"http://127.0.0.1:8080/basket/add" // Добавить продукт в корзину или обновить количество
-	"http://127.0.0.1:8080/basket/increase" // Увеличить количество товара
-	"http://127.0.0.1:8080/basket/decrease" // Уменьшить количество товара
-	"http://127.0.0.1:8080/basket/remove" // Удалить товар из корзины

Однако при подключении с эмулятора необходимо заменить ```127.0.0.1``` на ```10.0.2.2```

2) Создала новый проект Flutter. Добавила библиотеку для работы с API (файл **pubspec.yaml**)

```
dependencies:
  dio: ^5.7.0
```
3) Для получения данных через API был создан файл **server-api.dart** в папке ***lib/server-api/***
4) Продолжила работу над проектом с 8-ой практической, получая данные через API

- Главный экран: каталог настольных игр в две колонки.

![Снимок экрана 2024-11-04 111150](https://github.com/user-attachments/assets/9459dd58-5961-41ec-81fc-bf5f463403f5)

![Снимок экрана 2024-11-04 111208](https://github.com/user-attachments/assets/c95df5ee-ec27-4df0-bce2-2f1c550c6a50)

- Экран избранное: товары фильтруются из основного списка по сотоянию ***isFavorite***. Состояние обновляется как с карточки товара, так и с экрана информации о товаре через метод PUT (Обновление статуса избранного по ID)
  
<img src='https://github.com/user-attachments/assets/d671d051-7d4c-4824-bbc4-544c3ba392aa' width=300 />

<img src='https://github.com/user-attachments/assets/900abf80-a609-4139-882c-295db3c45d50' width=300 />

<img src='https://github.com/user-attachments/assets/f605348b-d8b8-482e-991d-78d9020730b8' width=300 />

- Экран корзины: добавление товаров через карточку товара или через экран информации о товаре. Если товара еще нет в корзине, то он добавляется, в противном случае счетчик товара увеличивается на 1. Удаление товара из корзины происходит свайпом вправо или влево (обработка запроса DELETE по ID товара)

![Снимок экрана 2024-11-04 111601](https://github.com/user-attachments/assets/c868327c-fce1-4029-aed2-550aa0af1d43)

![Снимок экрана 2024-11-04 111623](https://github.com/user-attachments/assets/af7f3f11-3322-4106-bf7b-b0698cde9386)


<img src='https://github.com/user-attachments/assets/f8577345-dc28-45ab-b247-a97f5cea1f4a' width=300 />

![Снимок экрана 2024-11-04 111733](https://github.com/user-attachments/assets/cf86e029-fadf-4647-9aed-bef66bf55b71)


- Экран профиля остался без изменений. Данные остаются только на клиентской части с возможностью их редактирования
  
<img src='https://github.com/user-attachments/assets/0b10f130-7821-44ab-b14e-ac56aca6ce09' width=300 />

<img src='https://github.com/user-attachments/assets/2893cd58-f311-46ca-a0c5-602c9e65a58c' width=300 />

- Экран добавления товара отправляет запрос POST на сервер (Создать продукт)
  
<img src='https://github.com/user-attachments/assets/f15ce095-b9b9-46b2-ba2d-a6b29b7e624e' width=300 />

<img src='https://github.com/user-attachments/assets/875598a5-bb89-44e6-8301-b9d1f53bc232' width=300 />

- Через экран информации о товаре можно удалить товар из каталога (обработка запроса DELETE по ID товара)

<img src='https://github.com/user-attachments/assets/a68419d6-fef0-4e80-872f-94e92137ee45' width=300 />

<img src='https://github.com/user-attachments/assets/bddccce9-4035-40f3-b829-b4ff1ee18241' width=300 />

- Экран редактирования информации о товаре. Добавлена функция редактирования товара через обработку запроса PUT (Обновление информации о товаре при редактировании по ID).
  
<img src='https://github.com/user-attachments/assets/63aa6cda-6847-41a1-85f5-d413d24f300b' width=300 />

<img src='https://github.com/user-attachments/assets/92b70d9e-69e1-43e4-809c-bb4954972245' width=300 />

<img src='https://github.com/user-attachments/assets/b7f3cfc3-b299-4669-b2b6-67582b7afac4' width=300 />

<img src='https://github.com/user-attachments/assets/f2197ac3-84ca-4ae3-a13d-e5647939fc92' width=300 />

<img src='https://github.com/user-attachments/assets/ed6fa081-2300-4d7b-af31-158f5ad0de28' width=300 />

<img src='https://github.com/user-attachments/assets/9d73aee2-4002-4bf6-aa49-2fb7be3e4485' width=300 />

<img src='https://github.com/user-attachments/assets/46ffe478-9936-441a-9211-35df1d9455aa' width=300 />

