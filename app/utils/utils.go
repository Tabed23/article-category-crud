package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Tabed23/article-category-crud/app/types"
)

func WriteOnFile(c *types.Category) {
	file, err := os.OpenFile("./data/categories.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Printf("Error opening category.txt: %v", err)
	}
	_, _ = file.WriteString(fmt.Sprintf("%s,%s \n", c.CategoryID, c.Name))
}

func FindByIdFile(id string, c *types.Category) (*types.Category, error) {
	file, err := os.Open("./data/categories.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if strings.Contains(scanner.Text(), id) {
			found := scanner.Text()
			ok := strings.Split(found, ",")
			c.CategoryID = ok[0]
			c.Name = ok[1]
			return c, nil
		}
	}

	err = scanner.Err()
	return nil, err
}

func DeleteByIdFromFile(id string) (bool, error) {

	file, err := os.Open("./data/categories.txt")
	if err != nil {
		return false, err
	}
	defer file.Close()

	var founds []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		data := strings.Split(line, ",")
		if len(data) < 2 {
			continue
		}
		if id != data[0] {
			founds = append(founds, line)
		}
	}

	if err := scanner.Err(); err != nil {
		return false, err
	}

	file.Close()
	file, err = os.Create("./data/categories.txt")
	if err != nil {
		return false, err
	}
	defer file.Close()

	for _, line := range founds {
		_, err := file.WriteString(line + "\n")
		if err != nil {
			return false, err
		}
	}
	file.Sync()
	return true, nil
}
