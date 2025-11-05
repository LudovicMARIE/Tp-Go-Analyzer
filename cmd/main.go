package main

import (
	"fmt"

	"github.com/ludovicMARIE/Tp-Go-Analyzer/internal/checker"
)

func main() {
	targets := []string{"https://www.google.com",
		"https://www.notarealwebsite.abc",
		"https://github.com",
		"https://www.movie.database/film/details",
		"https://www.gaming.news/release/new-game",
		"https://www.health.clinic/appointment/online",
		"https://www.car.manufacturer/model/electric",
		"https://www.home.decor/ideas/living-room",
		"https://www.environmental.org/project/clean-water",
		"https://www.space.agency/mission/mars",
		"https://www.fashion.magazine/trend/summer",
		"https://www.tech.conference/schedule/day1",
		"https://www.food.blog/recipe/dessert",
		"https://www.online.course/programming/python",
		"https://www.travel.guide/city/paris",
		"https://www.music.label/artist/new-album",
		"https://www.sports.club/events/match",
		"https://www.photography.tips/technique/lighting",
		"https://www.diy.tools/review/drill",
		"https://www.pet.vet/service/vaccination",
		"https://www.gardening.store/seeds/flower",
		"https://www.finance.advice/retirement/planning",
		"https://www.history.podcast/episode/ww2",
		"https://www.language.exchange/partner/find",
		"https://www.book.review/author/classic",
		"https://www.movie.review/genre/comedy",
		"https://www.gaming.forum/topic/strategy",
	}

	results := make(chan checker.CheckResult)

	for _, url := range targets {
		go checker.CheckUrl(url, results)
	}

	for range targets {
		result := <-results
		if result.Err != nil {
			fmt.Printf("KO %s : erreur - %v\n", result.Target, result.Err)
		} else {
			fmt.Printf("OK %s - %s \n", result.Target, result.Status)
		}
	}

}
