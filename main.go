package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	"strings"
)

// Тащим вакансии
func main() {
	// Создание коллектора
	c := colly.NewCollector()

	// Обработка каждого элемента, соответствующего указанному селектору
	c.OnHTML("div#a11y-main-content", func(e *colly.HTMLElement) {
		// Используем метод Each для итерации по всем найденным элементам
		e.DOM.Find("a.bloko-link").Each(func(_ int, s *goquery.Selection) {
			// Извлечение текста каждого элемента и удаление лишних пробелов
			vacancy := strings.TrimSpace(s.Text())
			// Проверка, не пустой ли текст вакансии
			if vacancy != "" {
				// Печать текста каждой вакансии
				fmt.Printf("Vacancy found: %s\n", vacancy)
			}
		})
	})

	// Вывод "Visiting ..." перед отправкой запроса
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Начало сканирования указанной страницы
	c.Visit("https://spb.hh.ru/search/vacancy?text=Golang&from=suggest_post&area=2&hhtmFrom=main&hhtmFromLabel=vacancy_search_line")
}
