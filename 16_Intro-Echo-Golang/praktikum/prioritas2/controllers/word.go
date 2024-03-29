package controllers

import (
	"net/http"
	"prioritas2_echo_golang/models"
	"slices"
	"strings"

	"github.com/labstack/echo/v4"
)

var Words = []models.Word{}

func GetAllWords(c echo.Context) error {

	resOk := models.ResponseSuccess{
		Data:    Words,
		Message: "all words",
	}

	return c.JSON(http.StatusOK, resOk)
}

func AddWord(c echo.Context) error {
	word := models.Word{}

	resErr := models.ResponseError{
		Message: "failed to add word data",
	}

	if err := c.Bind(&word); err != nil {
		return c.JSON(http.StatusBadRequest, resErr)
	}

	if word.Word == "" {
		return c.JSON(http.StatusBadRequest, resErr)
	}

	word.Length = len(word.Word)
	word.NumOfVocals = numOfVocals(word.Word)
	word.Palindrome = isPalindrome(word.Word)

	resOk := models.ResponseSuccess{
		Message: "word added",
		Data:    word,
	}

	Words = append(Words, word)

	return c.JSON(http.StatusCreated, resOk)
}

func numOfVocals(word string) int {
	var count int

	for _, c := range strings.ToLower(word) {
		if strings.ContainsAny(string(c), "aiueo") {
			count++
		}
	}

	return count
}

func isPalindrome(word string) bool {
	word = strings.ToLower(word)

	revSliceWord := strings.Split(word, "")
	slices.Reverse(revSliceWord)

	return slices.Equal(strings.Split(word, ""), revSliceWord)
}
