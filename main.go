package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

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

var basket = []Item{}

var products = []Product{
	{
		ID:          1,
		Title:       "Бункер",
		ImageURL:    "https://rolethedice.ru/wp-content/uploads/2021/02/%D0%A8%D0%B0%D0%B1%D0%BB%D0%BE%D0%BD2-6-1024x705.jpg",
		Description: "«Бункер» – спасение за пределами возможностей!",
		Rules:       "В начале партии случайным образом выбираются карта катастрофы и карты бункера. Затем игроки получают одинаковый набор карт характеристик: профессия, биология, здоровье, хобби, багаж и факт. Они характеризуют вашего персонажа. Далее игра длится несколько раундов, в ходе которых игроки раскрывают по одной своей карте другим игрокам и стараются убедить их, что данная характеристика может быть полезна для жизни в бункере. В конце каждого раунда игроки голосуют против того, кто по их мнению является наиболее бесполезным членом группы – этот игрок покидает игру.",
		Age:         18,
		Gamers:      "4-16 игроков",
		Time:        "30-60 минут",
		Price:       2990,
		Indicator:   1,
		IsFavorite:  false,
	},
	{
		ID:          2,
		Title:       "Ужас Аркхэма",
		ImageURL:    "https://22games.net/wp-content/uploads/2024/07/The-Dream-Eaters-Investigator-Expansion-1.jpg",
		Description: "Два мира, две древние силы – и только вы стоите между ними.",
		Rules:       "Игроки управляют действиями персонажей посредством своей колоды сыщика. Сценарии ведут кампании, предлагая интереснейшее приключение.\n\nКаждый раунд начинается с фазы Мифа. В эту фазу на карту агенды кладётся жетон ужаса, и каждый игрок тянет по карте из колоды контактов. Если на игровой карте проклятие, то проходится проверка указанного на проклятии скилла, и, если вы терпите неудачу, разыгрывается какой-то негативный эффект. Если на карте монстр, то он закрепляется за сыщиком и ему придётся либо сразиться с ним, либо попробовать проскользнуть.\n\nЗатем наступает фаза Расследования, где игроки в любом удобном для себя порядке выполняют по 3 действия из доступных:\n• Добор карты\n• Получение ресурсов\n• Активация действия с ранее выложенной карты\n• Призыв монстра к себе\n• Попытка увернуться от монстра\n• Атака монстра\n• Попытка сбора улик\n• Переход в другую локацию\n• Розыгрыш карты с руки за ресурсы\nПосле хода игроков наступает фаза Монстров. Различных противников очень много, некоторые враги двигаются по направлению к героям, а те, что уже подобрались, атакуют, автоматом навешивая ужас и раны на сыщиков.\n\nВ завершении раунда происходит фаза Передышки, в которой все игроки добирают по карте и получают по одному ресурсу.\n",
		Age:         13,
		Gamers:      "1-4 игрока",
		Time:        "1-2 часа",
		Price:       4490,
		Indicator:   2,
		IsFavorite:  false,
	},
	{
		ID:          3,
		Title:       "Descent",
		ImageURL:    "https://22games.net/wp-content/uploads/2023/10/Descent-Legends-of-the-Dark-The-Betrayers-War-11.jpg",
		Description: "Драконы, мятеж и древние враги – будущее Терринота под угрозой.",
		Rules:       "Игроки выбирают героев, получают стартовое снаряжение и отправляются в бой. Герои перемещаются по полю, сражаясь с монстрами, взаимодействуя с трёхмерным окружением и друг другом и добывая новое оружие, доспехи и артефакты. Интересная особенность: карта атаки составляется из двух двухсторонних карт оружия, так что вы сами готовите удар героя (впоследствии карта атаки переворачивается и активируются оставшиеся на другой стороне умения).\n\nПартия в Descent состоит из серии раундов, а каждый раунд – из фазы героев, где они ходят, атакуют противников и взаимодействуют с окружающим миром, и фазы тьмы, в которой активируются полученные дебафы, ходят враги, случаются сценарные события и замышляются коварные планы.\n\nЭто полностью кооперативная игра, то есть побеждают или проигрывают все игроки вместе. Победить можно только выполнив все задачи сценария, соответственно, если выполнить какую-то задачу становится невозможно – герои проигрывают.\n\nТакже герои проиграют, если один из них получает ранение, при этом уже страдая тяжёлым недугом. Итоги одного сценария могут повлиять на прохождение следующих.\n",
		Age:         13,
		Gamers:      "1-4 игрока",
		Time:        "3-4 часа",
		Price:       17990,
		Indicator:   3,
		IsFavorite:  true,
	},
	{
		ID:          4,
		Title:       "Манчкин",
		ImageURL:    "https://belosnejka52.ru/image/cache/data/products/razvivayushchieigrushkiknigi/nastolniedidakticheskieigri/9/FaXwcdvsErc-600x600-600x600.jpg",
		Description: "В мире «Манчкин» дружба и предательство идут рука об руку!",
		Rules:       "В начале партии игроки начинают как обычные люди своего пола (да-да, в игре есть даже различие полов!) первого уровня. В ваш ход вы выбиваете дверь, то есть переворачиваете верхнюю карту колоды дверей, и выполняете условия, написанные на карте. Если же это карта монстра, то вам не повезло и придётся вступить с ним в бой. Когда сила игрока больше или равна силе монстра, тогда игрок побеждает и может взять столько карт сокровищ, сколько указанно на карте с монстром. Во время боя другие игроки могут помогать товарищу и вступить с ним в бой, или же мешать и всячески усиливать вылезшего врага. Если игрок побеждает в бою, то он получает один уровень. В течение партии игроки будут применять на себя разные классы и профессии, менять свой пол, напяливать различные шмотки, пить подозрительные зелья, а также сражать самых разнообразных монстров!",
		Age:         12,
		Gamers:      "3-6 игроков",
		Time:        "30+ минут",
		Price:       1290,
		Indicator:   1,
		IsFavorite:  false,
	},
	{
		ID:          5,
		Title:       "Взрывные котята",
		ImageURL:    "https://22games.net/wp-content/uploads/2023/06/Exploding-Kittens-2-Player-Version-6.jpg",
		Description: "«Взрывные котята» – милота, которая может взорвать ваш вечер!",
		Rules:       "В центре стола находится колода карт, в которой, помимо прочих, есть взрывные котята. В свой ход игрок может разыграть карту с руки, а затем должен взять верхнюю карту из колоды. Если попался взрывной котёнок, то вы тут же выбываете из игры. Вы ведь только взорвались, между прочим!\n\nКаждый игрок получает на руки по восемь случайных карт, среди которых одна специальная карта «Обезвредь». С помощью этой карты вы сможете выжить, если вытащите взрывного котейку. Это самые ценные карты и их надо набирать побольше!\n\nИгроки ходят по очереди, тянут карты и очень стараются не взорваться. В этом им помогут специальные карты, которые можно разыграть в начале своего хода:\n• Карта «Нападай» – заставит следующего за вами игрока сделать подряд два хода, то есть, взять две карты, соответственно, шанс на подрыв тоже увеличится в два раза!\n• Карта «Неть» – отменяет действие предыдущей карты\n• Карта «Слиняй» – позволяет вам не брать карту в свой ход\n• Карта «Подлижись» – заставляет игрока отдать вам одну карту\n• Карта «Затасуй» – позволяет перемешать колоду. Полезно, если вы знаете, что верхняя карта подорвёт вас к чертям!\n• Карта «Подсмуртри грядущее» – позволяет подглядеть три верхние карты колоды!\n• Кошкокарты! Сами по себе бесполезные, если бы не ещё одно правило комбинаций. Например, разыграв две любые одинаковые карты, вы можете забрать карту у противника. За три одинаковые карты вы сможете выбрать определённую карту у врага. А за пять разных – можно покопаться в стопке сброса и забрать себе одну любую карту\n\nПобедит игрок, умудрившийся за всю партию не подорваться на котейке, то есть, последний оставшийся в живых игрок! Поздравляем, вы прирождённый котоман и сапёр по совместительству!",
		Age:         10,
		Gamers:      "2-5 игроков",
		Time:        "15+ минут",
		Price:       990,
		Indicator:   2,
		IsFavorite:  true,
	},
	{
		ID:          6,
		Title:       "Каркассон",
		ImageURL:    "https://static.insales-cdn.com/files/1/1792/2000640/original/carcassone-inside-box.jpg",
		Description: "Стройте своё королевство с умом и стратегией!",
		Rules:       "В «Каркассоне» действия игроков предельно просты. В свой ход игрок должен взять картонный квадрат с участком средневекового ландшафта и выложить его на стол к уже выложенным квадратам по принципу домино. Новый квадрат должен продолжать уже существующую картину мира – дорога смыкается с дорогой, пастбище с пастбищем, городская стена – с другим её участком. Когда благодаря размещению квадрата на столе возникает новый объект (например, дорога отходит в сторону от перекрёстка), игрок, выложивший этот квадрат, может застолбить этот объект за собой. Для этого он должен взять одну деревянную фигурку из своего резерва и поставить её на квадрат.\n\nТеперь в его интересах развить и завершить этот объект (в нашем примере – замкнуть дорогу следующим перекрёстком). Развивать объекты следует, соблюдая известную осторожность: если игрок не завершит этот объект до конца игры, он не получит за него ничего. По завершении, игрок снимает своего человечка с объекта и получает определённое число очков. \n",
		Age:         7,
		Gamers:      "2-5 игроков",
		Time:        "35+ минут",
		Price:       1990,
		Indicator:   3,
		IsFavorite:  false,
	},
	{
		ID:          7,
		Title:       "Городской убийца",
		ImageURL:    "https://www.cardplace.ru/uploads/cardplace/ff2d5fc3/dc651adf84755264681490352063dcef.jpg",
		Description: "«Городской убийца» – погрузитесь в мир паранойи и хитроумных маневров!",
		Rules:       "Игроки делятся на 2 команды: одна берёт на себя роль серийного убийцы, а другая играет за детектива. Также можно играть дуэльно или одному против всех. Пожалуй, оптимальный состав игроков: один убийца, два детектива.\n\nДействие игры происходит в середине прошлого века в одном из американских городов. Неназванный, но достаточно крупный город населяют множество ярких персонажей, которые представлены в игре 54 уникальными картами жителей. У каждого жителя есть характеристики: пол, профессия, социальная группа, возраст, телосложение и рост.\n\nПеред началом игры случайно определяются 20 жителей города, которые будут участвовать в партии. Игрок за убийцу перемешивает эти карты и смотрит одну из них в тайне от остальных – это и будет личность убийцы. Игрок записывает его личность и характеристики в свой листок. После этого колода с картами жителей перемешивается, и игрок-детектив раскладывает карты жителей на игровом поле, изображающем город – по одной в каждый квартал и по две в угловые кварталы города.\n\nТаким же образом из 24 карт мотивов на игру отбираются 6: игрок-убийца перемешивает их, берёт случайную карту в тайне от остальных игроков и записывает её рядом с характеристиками своей личности – это мотив убийцы, который игрок-детектив тоже должен вычислить. Все убийства в игре убийца будет совершать строго по правилам своего мотива. Затем убийца замешивает карту мотива с 5 другими картами и выкладывает перед собой все 6 карт на столе.\n\nЗапутать следы убийце поможет социальная группа его союзников. Игрок-убийца перемешивает колоду жетонов социальных групп и получает 3 жетона, среди которых он выбирает себе один и тоже записывает его на своем листке. Жители выбранной социальной группы будут мешать вести расследование детективу.",
		Age:         16,
		Gamers:      "2-4 игрока",
		Time:        "60 минут",
		Price:       2490,
		Indicator:   1,
		IsFavorite:  true,
	},
	{
		ID:          8,
		Title:       "Неудержимые единорожки",
		ImageURL:    "https://vsedrugoeshop.ru/upload/ammina.optimizer/jpg-webp/q80/upload/iblock/c28/7gm2ejvfpfllj91at010agphibre6x6r/6350854204.webp",
		Description: "Встречайте весёлую карточную игру для любой вечеринки.",
		Rules:       "В начале игры нужно отделить карты малышей-единорогов и памятки от основной колоды. После этого нужно выбрать одного из малышей и поместить себе в стойло, оставшиеся отправляются в отдельную колоду – ясли. Малыши не должны попадать в основную колоду или сброс: они могут находиться только в стойле или в яслях. После этого игроки ходят. Ход состоит из нескольких фаз: в первую фазу применяются карты с мгновенными эффектами, во вторую берётся карта из колоды, в третью совершается дополнительное действие, а в четвёртую сбрасываются лишние карты (максимум на руке – 7 карт, если лимит не увеличен каким-либо эффектом). Игра продолжается, пока не закончатся карты из основной колоды или один из игроков не наберёт необходимое число единорогов.\n\nПобеждает игрок, собравший в своём стойле 7 единорожков, если в игре участвуют 2-5 игроков, и 6 – если игроков от 6 до 8. Если игра закончилась из-за того, что кончились карты в основной колоде, то побеждает игрок, набравший наибольшее число единорожков.\n\nЕсли на победу претендуют несколько игроков, им придётся подсчитать количество букв в названиях своих карт единорожков, побеждает тот, у кого больше. Если же и по этому параметру на победу претендуют несколько игроков, то никто не выигрывает.",
		Age:         12,
		Gamers:      "2-8 игроков",
		Time:        "30-60 минут",
		Price:       1290,
		Indicator:   2,
		IsFavorite:  false,
	},
	{
		ID:          9,
		Title:       "Игра UNO",
		ImageURL:    "https://um-detki.ru/wp-content/uploads/2019/03/uno3.jpg",
		Description: "В этой маленькой коробочке притаилось большое веселье!",
		Rules:       "Ваша задача – первым набрать 500 очков. А чтобы это сделать, в каждом раунде вы должны быстрее остальных избавляться от карт на руках. Вначале партии открывается одна карта из колоды, а дальше по часовой стрелке игроки выкладывают одну карту с руки. И продолжается это до тех пор, пока кто-то не сбросит все свои карты. Запомните главное правило: если у вас осталась одна карта, обязательно крикните «Уно!», ведь если этого не сделать, а другие игроки заметят это, придётся добирать карты из колоды.\n\nПодробнее о картах\nВ состав игры входят 108 карт четырёх цветов, которые делятся на обычные, карты «Действия» и «Дикие» карты.\n\nВсе обычные карты имеют значение от 0 до 9 и разделены поровну на 4 цвета: зелёный, жёлтый, красный и голубой. Их в игре больше остальных. Чтобы выложить с руки обычную карту, она должна совпадать с верхней открытой картой на столе либо по цвету, либо по значению. То есть, если на столе лежит красная пятёрка, то вы можете выложить на неё или пятёрку любого цвета, или красную карту с любым значением. В конце раунда они приносят победные очки по своему номиналу.\n\nИграть с одними обычными картами было бы скучно, поэтому в UNO и существуют карты действий! Всего в игре их 24, они делятся на 3 типа. «Смена направления» меняет порядок хода игроков. Выкладывали карты по часовой стрелке? Теперь наоборот! «Пропусти ход» говорит само за себя: следующий игрок пропускает свой ход. А «Вытяни две» обязывает следующего игрока не и пропустить свой ход, и добрать 2 карты. Кстати, «Действия» разыгрываются, как и обычные карты: либо на тот же цвет, либо на то же «Действие». В конце раунда каждая карта «Действий» приносит 20 очков.\n\nДикие карты - это самые настоящие «джокеры»! Всего существует 2 вида «Диких» карт: «Выбери цвет» и «Выбери цвет, Вытяни четыре». Первую можно играть в любой свой ход на любую карту. Она позволяет вам назначит цвет, которым продолжится игра. А карта «Выбери цвет, Вытяни четыре» намного коварнее: её можно сыграть на любую карту, она так же даёт вам право назначить цвет следующей карты, а ещё заставляет игрока после вас снять 4 карты и пропустить свой ход! Но есть одно «Но» – её нельзя сыграть, пока у вас есть любая другая карта, которой можно продолжить игру. В конце раунда каждая «Дикая» карта приносит победителю 50 очков.",
		Age:         7,
		Gamers:      "2-10 игроков",
		Time:        "15-45 минут",
		Price:       790,
		Indicator:   3,
		IsFavorite:  false,
	},
	{
		ID:          10,
		Title:       "Мафия: Вся семья в сборе",
		ImageURL:    "https://avatars.mds.yandex.net/get-mpic/5231998/2a0000018e2424cb5a1f0eb06e85c1e4bc2d/orig",
		Description: "«Мафия» – классика психологических игр!",
		Rules:       "Игра делится на несколько фаз: сначала игроки выбирают ведущего, взакрытую случайным образом делятся на команды мирных жителей и мафии (включая специальные роли) с помощью карт, а затем наступает «ночь». В фазу ночи все игроки закрывают глаза, то есть «засыпают», но мафия не дремлет и выходит на охоту! В первую ночь мафия молча знакомится друг с другом, в последующие могут убить одного из мирных; помимо этого ночью мирный комиссар может узнать роль любого игрока. Наступает следующая фаза – «день», в которой вскрывает личность убитого ночью, затем все игроки совещаются, обмениваются мнениями о том, кто, скорее всего, является членом мафии и общим голосованием решает, кого казнить.",
		Age:         14,
		Gamers:      "7-17 игроков",
		Time:        "20+ минут",
		Price:       2490,
		Indicator:   1,
		IsFavorite:  false,
	},
}

// обработчик для GET-запроса, возвращает список продуктов
func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(products)
}

// Функция вычисления нового ID продукта
func getNextID(products []Product) int {
	maxID := 0
	for _, product := range products {
		if product.ID > maxID {
			maxID = product.ID
		}
	}
	return maxID + 1
}

// Обработчик для POST-запроса, добавляет продукт
func createProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var newProduct Product
	err := json.NewDecoder(r.Body).Decode(&newProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Received new Product: %+v\n", newProduct)

	newProduct.ID = getNextID(products)
	products = append(products, newProduct)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newProduct)
}

// Обработчик для обновления информации о продукте
func updateProduct(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/products/update/info/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}

	var updateProduct Product
	err = json.NewDecoder(r.Body).Decode(&updateProduct)
	if err != nil {
		fmt.Println("Error decoding request body:", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i].Age = updateProduct.Age
			products[i].Description = updateProduct.Description
			products[i].Gamers = updateProduct.Gamers
			products[i].ImageURL = updateProduct.ImageURL
			products[i].Indicator = updateProduct.Indicator
			products[i].IsFavorite = updateProduct.IsFavorite
			products[i].Price = updateProduct.Price
			products[i].Rules = updateProduct.Rules
			products[i].Time = updateProduct.Time
			products[i].Title = updateProduct.Title

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Добавление маршрута для получения одного продукта

func getProductByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем продукт с данным ID
	for _, Product := range products {
		if Product.ID == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(Product)
			return
		}
	}

	// Если продукт не найден
	http.Error(w, "Product not found", http.StatusNotFound)
}

// Удаление продукта по id
func deleteProductHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Получаем ID из URL
	idStr := r.URL.Path[len("/Products/delete/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid Product ID", http.StatusBadRequest)
		return
	}

	// Ищем и удаляем продукт с данным ID
	for i, Product := range products {
		if Product.ID == id {
			// Удаляем продукт из среза
			products = append(products[:i], products[i+1:]...)
			w.WriteHeader(http.StatusNoContent) // Успешное удаление, нет содержимого
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

// Функция для поиска продукта по ID
func findProductByID(id int) (*Product, bool) {
	for _, product := range products {
		if product.ID == id {
			return &product, true
		}
	}
	return nil, false
}

// обработчик для GET-запроса, возвращает список элементов в корзине
func getBasketHandler(w http.ResponseWriter, r *http.Request) {
	// Устанавливаем заголовки для правильного формата JSON
	w.Header().Set("Content-Type", "application/json")
	// Преобразуем список заметок в JSON
	json.NewEncoder(w).Encode(basket)
}

// Проверка наличия товара в корзине
func checkBasketItemHandler(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/Basket/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid product ID", http.StatusBadRequest)
		return
	}

	for _, item := range basket {
		if item.ID == id {
			response := map[string]interface{}{
				"isInBasket": true,
				"itemCount":  item.Counter,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	}

	response := map[string]interface{}{
		"isInBasket": false,
		"itemCount":  0,
	}
	json.NewEncoder(w).Encode(response)
}

// Обработчик для POST-запроса, добавляет продукт в корзину или обновляет количество
func addToBasketHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GameID int `json:"gameId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, item := range basket {
		if item.ID == req.GameID {
			basket[i].Counter++
			json.NewEncoder(w).Encode(basket[i])
			return
		}
	}

	newItem := Item{ID: req.GameID, Counter: 1}
	basket = append(basket, newItem)
	json.NewEncoder(w).Encode(newItem)
}

// Увеличение количества товара
func increaseBasketItemHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GameID int `json:"gameId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, item := range basket {
		if item.ID == req.GameID {
			basket[i].Counter++
			json.NewEncoder(w).Encode(basket[i])
			return
		}
	}
	http.Error(w, "Item not found in basket", http.StatusNotFound)
}

// Уменьшение количества товара в корзине
func decreaseBasketItemHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GameID int `json:"gameId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, item := range basket {
		if item.ID == req.GameID {
			if basket[i].Counter > 1 {
				basket[i].Counter--
				json.NewEncoder(w).Encode(basket[i])
			}
			return
		}
	}
	http.Error(w, "Item not found in basket", http.StatusNotFound)
}

// Удаление товара из корзины
func removeFromBasketHandler(w http.ResponseWriter, r *http.Request) {
	var req struct {
		GameID int `json:"gameId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, item := range basket {
		if item.ID == req.GameID {
			basket = append(basket[:i], basket[i+1:]...)
			json.NewEncoder(w).Encode(map[string]string{"message": "Item removed from basket"})
			return
		}
	}
	http.Error(w, "Item not found in basket", http.StatusNotFound)
}

// Обработчик для обновления статуса избранного
func updateFavoriteStatus(w http.ResponseWriter, r *http.Request) {

	idStr := r.URL.Path[len("/products/update/status/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid item ID", http.StatusBadRequest)
		return
	}
	var req struct {
		IsFavorite bool `json:"isFavorite"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	for i, product := range products {
		if product.ID == id {
			products[i].IsFavorite = req.IsFavorite

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(products[i])
			return
		}
	}

	http.Error(w, "Product not found", http.StatusNotFound)
}

func main() {
	http.HandleFunc("/products", getProductsHandler)                  // Получить все продукты
	http.HandleFunc("/products/create", createProductHandler)         // Создать продукт
	http.HandleFunc("/products/", getProductByIDHandler)              // Получить продукт по ID
	http.HandleFunc("/products/delete/", deleteProductHandler)        // Удалить продукт
	http.HandleFunc("/products/update/status/", updateFavoriteStatus) // Обновление статуса избранного
	http.HandleFunc("/products/update/info/", updateProduct)          // Обновить информацию о продукте

	http.HandleFunc("/basket", getBasketHandler)                   // Получить все элементы корзины
	http.HandleFunc("/basket/", checkBasketItemHandler)            // Проверить есть ли товар в корзине
	http.HandleFunc("/basket/add", addToBasketHandler)             // Добавить продукт в корзину или обновить количество
	http.HandleFunc("/basket/increase", increaseBasketItemHandler) // Увеличить количество товара
	http.HandleFunc("/basket/decrease", decreaseBasketItemHandler) // Уменьшить количество товара
	http.HandleFunc("/basket/remove", removeFromBasketHandler)     // Удалить товар из корзины

	fmt.Println("Server is running on port 8080!")
	http.ListenAndServe(":8080", nil)
}
