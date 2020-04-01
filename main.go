/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

func main() {
	client := Client{bucket: make([]Item, 3)}
	client.addItemToTheBucket(Item{name: "Колготки", price: 150, description: "штопаные"})
	client.addItemToTheBucket(Item{name: "Мясо", price: 250, description: "тухлое"})
	client.addItemToTheBucket(Item{name: "Батон", price: 40, description: ""})

	employee := Iterator{bucket: client.bucket}
	employee.iterateItems()
}
