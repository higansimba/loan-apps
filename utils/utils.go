package utils

import (
	"fmt"
	"math"
	"math/rand"
)

// domains represents common email domains
var domains = []string{
	"gmail.com",
	"yahoo.com",
	"hotmail.com",
	"outlook.com",
	"example.com",
}

// generateRandomString creates a random string of specified length
func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyz0123456789"
	result := make([]byte, length)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result)
}

// GenerateEmail creates a random email address
func GenerateEmail() string {
	// r := time.Now().UnixNano()

	// Generate random username (5-15 characters)
	// usernameLength := r + 5
	username := GenerateRandomString(8)

	// Select random domain
	domain := domains[rand.Intn(len(domains))]

	// Combine username and domain
	return fmt.Sprintf("%s@%s", username, domain)
}

// GenerateUniqueEmails generates a specified number of unique emails
func GenerateUniqueEmails(count int) []string {
	emails := make(map[string]bool)
	result := make([]string, 0, count)

	for len(result) < count {
		email := GenerateEmail()
		if !emails[email] {
			emails[email] = true
			result = append(result, email)
		}
	}

	return result
}

func GenerateRandomNumber(min, max int) int {
	return rand.Intn(max-min+1) + min
}
func GenerateInstallmentRate(amount float64, tenor float64) (weeklyInstallment float64, totalInterest float64) {
	// Bunga flat 10% per tahun dari pokok pinjaman
	annualInterestRate := 0.10
	totalInterest = amount * annualInterestRate // Misal: 10% dari 5jt = 500rb

	// Total yang harus dibayar: pokok + bunga
	totalPayment := amount + totalInterest

	// Cicilan per minggu
	weeklyInstallment = totalPayment / tenor

	// Bulatkan ke rupiah terdekat
	weeklyInstallment = math.Round(weeklyInstallment)

	return weeklyInstallment, totalInterest
}
