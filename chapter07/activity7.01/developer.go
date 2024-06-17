package main

import (
	"errors"
	"fmt"
	"strings"
)

type Developer struct {
	Individual        Employee
	HourlyRate        int
	HoursWorkedInYear int
	Review            map[string]any
}

func (d Developer) Pay() (string, float64) {
	fullName := d.Individual.FullName()
	yearPay := float64(d.HourlyRate * d.HoursWorkedInYear)

	return fullName, yearPay
}

func (d Developer) ReviewRating() error {
	total := 0

	for _, v := range d.Review {
		rating, err := overallReview(v)
		if err != nil {
			return err
		}

		total += rating
	}

	averageRating := float64(total) / float64(len(d.Review))
	fmt.Printf("%s got a review rating of %.2f\n", d.Individual.FullName(), averageRating)
	return nil
}

func overallReview(i any) (int, error) {
	switch v := i.(type) {
	case int:
		return v, nil
	case string:
		rating, err := convertReviewToInt(v)
		if err != nil {
			return 0, err
		}
		return rating, nil
	default:
		return 0, errors.New("unknown type")
	}
}

func convertReviewToInt(str string) (int, error) {
	switch strings.ToLower(str) {
	case "excellent":
		return 5, nil
	case "good":
		return 4, nil
	case "fair":
		return 3, nil
	case "poor":
		return 2, nil
	case "unsatisfactory":
		return 1, nil
	default:
		return 0, errors.New("invalid rating: " + str)
	}
}
