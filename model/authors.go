package model

import "fmt"

//AuthorDto - Author structure
type Author struct {
	UUID      string `json:"uuid"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Birthday  string `json:"birthday"`
	Death     string `json:"death"`
}

func (a Author) String() string {
	return fmt.Sprintf("AuthorDto{ UUID = %s, FirstName = %s, LastName = %s, Birthday = %s, Death = %s}",
		a.UUID, a.FirstName, a.LastName, a.Birthday, a.Death)
}

func (a *AuthorList) GetAuthorByUuid(uuid string) (Author, error) {
	err := fmt.Errorf("Could not find any author with that UUID %s", uuid)

	for _, author := range *a {
		if author.UUID == uuid {
			return author, nil
		}
	}

	return Author{}, err
}

func (a *AuthorList) DeleteAuthorWithUuid(uuid string) (error) {
	var err = fmt.Errorf("Could not find any author with that uuid %s", uuid)
	var updatedAuthors AuthorList
	for _, author := range *a {
		if author.UUID == uuid {
			err = nil
		} else {
			updatedAuthors = append(updatedAuthors, author)
		}
	}
	if err == nil {
		*a = updatedAuthors
	}
	return err
}

func (a *AuthorList) UpdateAuthorWithUuid(updatedAuthor Author) (Author, error) {
	var err = fmt.Errorf("could not find author by uuid %s", updatedAuthor.UUID)
	var newAuthors AuthorList
	for _, author := range *a {
		if author.UUID == updatedAuthor.UUID {
			newAuthors = append(newAuthors, updatedAuthor)
			err = nil
		} else {
			newAuthors = append(newAuthors, author)
		}
	}
	if err == nil {
		*a = newAuthors
	}
	return updatedAuthor, err
}

func (a *AuthorList) AddAuthor(author Author) {
	*a = append(*a, author)
}

type AuthorList []Author

var Authors AuthorList