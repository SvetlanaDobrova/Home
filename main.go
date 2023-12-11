package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5"
	"os"
)

func main() {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	var jsonStr string
	// = "{\"address\":\"Московская область, г. Дедовск\",\"area\":42.5,\"rooms\":[{\"area\":3,\"furniture\":[{\"height\":2,\"length\":0.4,\"name\":\"Трюмо\",\"width\":0.8},{\"height\":2,\"length\":0.1,\"name\":\"Вешалка\",\"width\":1},{\"height\":2.1,\"length\":0.4,\"name\":\"шкаф\",\"width\":1}],\"height\":2.4,\"length\":2.5,\"name\":\"Прихожая\",\"width\":1.2},{\"area\":6,\"furniture\":[{\"height\":1,\"length\":0.8,\"name\":\"Плита\",\"width\":1},{\"height\":2,\"length\":0.5,\"name\":\"Холодильник\",\"width\":0.5},{\"height\":0.4,\"length\":0.45,\"name\":\"Духовка\",\"width\":0.5},{\"height\":1,\"length\":0.7,\"name\":\"Посудомойка\",\"width\":0.8},{\"height\":1,\"length\":1.5,\"name\":\"Раковина\",\"width\":0.7},{\"height\":0.3,\"length\":0.35,\"name\":\"Микроволновка\",\"width\":0.6}],\"height\":2.4,\"length\":3,\"name\":\"Кухня\",\"width\":2},{\"area\":18,\"furniture\":[{\"height\":0.5,\"length\":2.5,\"name\":\"Кровать\",\"width\":3},{\"height\":2,\"length\":1,\"name\":\"Стол\",\"width\":2},{\"height\":2.3,\"length\":0.75,\"name\":\"Шкаф\",\"width\":2}],\"height\":2.4,\"length\":6,\"name\":\"Спальня\",\"width\":3},{\"area\":12.65,\"furniture\":[{\"height\":0.5,\"length\":1,\"name\":\"Диван\",\"width\":3},{\"height\":1.2,\"length\":2,\"name\":\"Стол\",\"width\":1},{\"height\":2.1,\"length\":0.5,\"name\":\"шкаф\",\"width\":3}],\"height\":2.4,\"length\":3,\"name\":\"Гостинная\",\"width\":5.5}],\"tenants\":[{\"age\":20,\"name\":\"Доброва Светлана Борисовна\"},{\"age\":62,\"name\":\"Маслова Любовь Сергеевна\"}]}"

	// Подключимся к БД postgres. Строка берется из переменной окружения DATABASE_URL
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	// Извлечем json с нашим проектом из таблицы home в переменную jsonStr
	err = conn.QueryRow(context.Background(), "select project from home ").Scan(&jsonStr)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	var jsDeserialize home

	// де-сериализуем json в struct home
	_ = json.Unmarshal([]byte(jsonStr), &jsDeserialize)

	/* var myFlat home = home{
		Area:    42.5,
		Address: "Московская область, г. Дедовск",
		Rooms: []room{
			{
				Name:   "Прихожая",
				Height: 2.4,
				Width:  1.2,
				Length: 2.5,
				Area:   1.2 * 2.5,
				Furniture: []furniture{
					{
						Height: 2,
						Width:  0.8,
						Length: 0.4,
						Name:   "Трюмо",
					},
					{
						Height: 2,
						Width:  1,
						Length: 0.1,
						Name:   "Вешалка",
					},
					{
						Height: 2.1,
						Width:  1,
						Length: 0.4,
						Name:   "шкаф",
					},
				},
			},
			{
				Name:   "Кухня",
				Height: 2.4,
				Width:  2,
				Length: 3,
				Area:   2 * 3,
				Furniture: []furniture{
					{
						Height: 1,
						Width:  1,
						Length: 0.8,
						Name:   "Плита",
					},
					{
						Height: 2,
						Width:  0.5,
						Length: 0.5,
						Name:   "Холодильник",
					},
					{
						Height: 0.4,
						Width:  0.5,
						Length: 0.45,
						Name:   "Духовка",
					},
					{
						Height: 1,
						Width:  0.8,
						Length: 0.7,
						Name:   "Посудомойка",
					},
					{
						Height: 1,
						Width:  0.7,
						Length: 1.5,
						Name:   "Раковина",
					},
					{
						Height: 0.3,
						Width:  0.6,
						Length: 0.35,
						Name:   "Микроволновка",
					},
				},
			},
			{
				Name:   "Спальня",
				Height: 2.4,
				Width:  3,
				Length: 6,
				Area:   6 * 3,
				Furniture: []furniture{
					{
						Height: 0.5,
						Width:  3,
						Length: 2.5,
						Name:   "Кровать",
					},
					{
						Height: 2,
						Width:  2,
						Length: 1,
						Name:   "Стол",
					},
					{
						Height: 2.3,
						Width:  2,
						Length: 0.75,
						Name:   "Шкаф",
					},
				},
			},
			{
				Name:   "Гостинная",
				Height: 2.4,
				Width:  5.5,
				Length: 3,
				Area:   5.5 * 2.3,
				Furniture: []furniture{
					{
						Height: 0.5,
						Width:  3,
						Length: 1,
						Name:   "Диван",
					},
					{
						Height: 1.2,
						Width:  1,
						Length: 2,
						Name:   "Стол",
					},
					{
						Height: 2.1,
						Width:  3,
						Length: 0.5,
						Name:   "шкаф",
					},
				},
			},
		},
		Tenants: []tenant{
			tenant{
				Name: "Доброва Светлана Борисовна",
				Age:  20,
			},
			{
				Name: "Маслова Любовь Сергеевна",
				Age:  62,
			},
		},
	} */

	// вывод структуры проекта
	showHome(jsDeserialize)

	//var str []byte
	//str, _ = json.Marshal(myFlat)
	//fmt.Printf("{[%v]}", string(str))
}
