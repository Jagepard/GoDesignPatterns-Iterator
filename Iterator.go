/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

import "fmt"

// Iterator is ...
type Iterator struct {
	bucket []Item
}

func (iterator Iterator) iterateItems() {
	for _, element := range iterator.bucket {
		fmt.Printf("%s %d %s \n", element.name, element.price, element.description)
	}
}
