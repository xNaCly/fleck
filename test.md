# Heading 1

<h1 class="heading" id="this-is-the-first-heading">
  This is the first heading
</h1>

This paragraph contains _italic_ and **bold** elements.

This is a paragraph containing an `inline code *_[]#!()` element (_with special characters_) and an inline hash: # as well as inline dash: -.

```js
console.log("test");

let a = [3, 5, 1, 2, 6];
a.sort((x, y) => x - y));

// this is sorted asc 😳
console.log(a);

```

This code block is split ^o^

```
// bit wise in forEach?
a.forEach((x)=>{
    console.log(x & 1);
});

// new rentner 🧓
let rentner = {
    name: "Gerald",
    age: 28,
    car: "911"
}
```

Now a better programming language:

```go
package main

import "log"

func main(){
    log.Println("Hello world!")
}
```

And now an even better programming language:

```c
#include <stdio.h>
#include <stdlib.h>

int main(void){
    for (int i = 0; i < 10; i++){
        if (i % 3 == 0) {
            printf("divisible by 3\n");
        } else if (i % 5 == 0){
            printf("divisible by 5\n");
        } else if (i % 15 == 0){
            printf("divisible by 5 & 3\n");
        } else {
            printf("%d", i);
        }
    }
    printf("You just go fizzbuzzed!\n");
    return EXIT_SUCCESS;
}
```

Now not even a programming language:

```html
<h1 class="heading" id="this-is-the-first-heading">
  This is the first heading
</h1>
```

> This is a blockquote

---

> This is a
> tripple line
> blockquote

[xnacly's homepage](https://xnacly.me)

[this is not a link

this is also not a link]()

[test](this can never be a link

[](a link with an empty title?)

![xnacly's profile picture](https://avatars.githubusercontent.com/u/47723417?v=4)

![](https://avatars.githubusercontent.com/u/47723417?v=4)

![image without a source, this will be an alt text]()

## Heading 2

### Heading 3

#### Heading 4

##### Heading 5

###### Heading 6

1. list

- list
- [x] checked list
- [ ] unchecked list

@include{README.md}
Today's date: @today{2006-01-02}
author: @shell{whoami}
