# Heading 1

$p,q = - \frac{p}{2} \pm \sqrt{\left(\frac{p}{2}\right)-q}$

$$
\begin{align}
\lim_{x \to \infty} (a_n+b_n) = \lim_{x \to \infty} a_n + \lim_{x \to \infty} b_n \\
\lim_{x \to \infty} (C \cdot a_n) = C \cdot \lim_{x \to \infty} a_n \\
\lim_{x \to \infty} (a_n \cdot b_n) = \lim_{x \to \infty} a_n \cdot \lim_{x \to \infty} b_n \\
\lim_{x \to \infty} |a_n| = \left| \lim_{x \to \infty} a_n \right|
\end{align}
$$

Source:

```latex
$$
\begin{align}
\lim_{x \to \infty} (a_n+b_n) = \lim_{x \to \infty} a_n + \lim_{x \to \infty} b_n \\
\lim_{x \to \infty} (C \cdot a_n) = C \cdot \lim_{x \to \infty} a_n \\
\lim_{x \to \infty} (a_n \cdot b_n) = \lim_{x \to \infty} a_n \cdot \lim_{x \to \infty} b_n \\
\lim_{x \to \infty} |a_n| = \left| \lim_{x \to \infty} a_n \right|
\end{align}
$$
```

> **Info**
>
> ### Heading
>
> fleck does not support inline html, but it does inline math $a^{32} / \frac{19}{1209}$ and even block math:
>
> $$
> \begin{align}
> p,q = - \frac{p}{2} \pm \sqrt{\left(\frac{p}{2}\right)-q} \\
> p,q = - \frac{p}{2} \pm \sqrt{\left(\frac{p}{2}\right)-q}
> \end{align}
> $$

> **Warning**
>
> fleck does not support inline html.

> **Danger**
>
> fleck does not support inline html.

> **Note**
>
> fleck does not support inline html.

```js
// `npm run this-shit`
console.log("🤬");

console.log(`${name} is dumb`);
```

<h1 class="heading" id="this-is-the-first-heading">
  This is the first heading
</h1>

This paragraph contains _italic_, **bold**, ==highlighted== and ~~striketrough~~ elements .

This is a paragraph containing an `inline code *_[]#!()` element (_with special characters_) and an inline hash: # as well as other special characters -=$`.

```js
console.log("test");

let a = [3, 5, 1, 2, 6];
a.sort((x, y) => x - y));

// this is sorted asc 😳
console.log(a);
```

This code block is split ^o^

```js
// bit wise in forEach?
a.forEach((x) => {
  console.log(x & 1);
});

// new rentner 🧓
let rentner = {
  name: "Gerald",
  age: 28,
  car: "911",
};
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

> This is a _simple_ blockquote

---

> This is a multi line blockquote
>
> ```js
> console.log("test");
> // even with comments in the codeblock
> console.log("test");
> ```
>
> ![xnacly's profile picture](https://avatars.githubusercontent.com/u/47723417?v=4)
>
> [this is not a link

this is also not a link]()

[test](this can never be a link

[](a link with an empty title?)

[xnacly's homepage](https://xnacly.me)

![](https://avatars.githubusercontent.com/u/47723417?v=4)

![image without a source, this will be an alt text]()

## Heading 2

### Heading 3

#### Heading 4

##### Heading 5

###### Heading 6

- list this is a very long line, possibly even too long, SAY WHATTTT
  test test test loong shit [link](google.com) `inline code`, **bold**, _italic_ and even
  arbitrary
  line
  breaks
  wtf
- list 2, this li even includes an image ![this includes an image](https://avatars.githubusercontent.com/u/47723417?v=4)

- [x] checked list
- [ ] unchecked list

## Included file:

---

@include{README.md}

---

Today's date: @today{2006-01-02}
author: @shell{whoami}
