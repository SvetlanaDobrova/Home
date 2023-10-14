package main

import "fmt"

type home struct {
	area    float64
	address string
	rooms   []room
	tenants []tenant
}

func showHome(h home) {
	fmt.Println(h.address)
	fmt.Println(" Проживают:")
	showTenants(h)
	fmt.Println(" Комнаты:")
	for _, r := range h.rooms {
		showRoom(r)
	}
}

func main() {

	var myFlat home = home{
		area:    42.5,
		address: "Московская область, г. Дедовск",
		rooms: []room{
			{
				name:   "Прихожая",
				height: 2.4,
				width:  1.2,
				length: 2.5,
				area:   1.2 * 2.5,
				furniture: []furniture{
					{
						height: 2,
						width:  0.8,
						length: 0.4,
						name:   "Трюмо",
					},
					{
						height: 2,
						width:  1,
						length: 0.1,
						name:   "Вешалка",
					},
					{
						height: 2.1,
						width:  1,
						length: 0.4,
						name:   "шкаф",
					},
				},
			},
			{
				name:   "Кухня",
				height: 2.4,
				width:  2,
				length: 3,
				area:   2 * 3,
				furniture: []furniture{
					{
						height: 1,
						width:  1,
						length: 0.8,
						name:   "Плита",
					},
					{
						height: 1,
						width:  0.7,
						length: 1.5,
						name:   "Раковина",
					},
					{
						height: 0.3,
						width:  0.6,
						length: 0.35,
						name:   "Микроволновка",
					},
				},
			},
			{
				name:   "Спальня",
				height: 2.4,
				width:  3,
				length: 6,
				area:   6 * 3,
				furniture: []furniture{
					{
						height: 0.5,
						width:  3,
						length: 2.5,
						name:   "Кровать",
					},
					{
						height: 2,
						width:  2,
						length: 1,
						name:   "Стол",
					},
					{
						height: 2.3,
						width:  2,
						length: 0.75,
						name:   "Шкаф",
					},
				},
			},
		},
		tenants: []tenant{
			tenant{
				name: "Доброва Светлана Борисовна",
				age:  20,
			},
			{
				name: "Маслова Любовь Сергеевна",
				age:  62,
			},
		},
	}
	showHome(myFlat)
}