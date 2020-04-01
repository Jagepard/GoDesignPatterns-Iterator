/**
 * @author  : Jagepard <jagepard@yandex.ru>
 * @license https://mit-license.org/ MIT
 */

package main

// Client is ...
type Client struct {
	bucket []Item
}

var i int = 0

func (client Client) addItemToTheBucket(item Item) {
	client.bucket[i] = item
	i++
}
