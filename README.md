# makeitso
_"It does it for you"_

**THIS IS A JOKE. DO NOT USE THIS IN PRODUCTION. EVEN RUNNING LOCALLY, CHECK THE GENERATED CODE BEFORE RUNNING**

## Introduction

Welcome to makeitso, the (mostly) language-agnostic AI-powered preprocessor designed to make your life easier by doing the heavy lifting for you. Please note that this tool is intended for humorous purposes only and should not be used in any production environment. Even when using it locally, it is essential to check the generated code before running it.

## How does it work?

You don't want to know. It's a bunch of wires, paperclips, more paperclips, duct tape and ummm... ChatGPT. Feel free to check the source, but don't judge me. This just needed to be made. That's why it's not super clean. It just is.


## What?

The main idea behind makeitso is to help you with complicated or tedious tasks in your programming job. Instead of writing complex code and getting it wrong, just put a description of what the code needs to do and make AI get it wrong. This way, the blame does not fall on you, it falls on the AI.

## That sounds crazy. How would one use this?

Taking inspiration from the infamous jQuery, the syntax of makeitso is designed to make it familiar and easy to use.
To leverage the power of makeitso, include a simple prompt in your code. The prompt follows the pattern:

```javascript
$("your command here")
```

For example:

```javascript
$("run fizzbuzz from 1 to 30")
```

The above example is just a demonstration and can be replaced with any task you need assistance with. The $(prompt) can be used as a statement or an expression, depending on the surrounding code and the prompt itself.

### Preprocessor

Files written in this manner can be run though makeitso, which will follow all the instruction in those $(prompt)s, replacing them with actual code.

## Usage Examples

### Simple stuff
```javascript
const a = "foo";
const b = "bar";
const result = $("append the two variables a and b");
```

### Hard stuff
```javascript
const n = $("give me the seventeenth prime number bro");
console.log(n);
```

### Math (we know it's hard)
```javascript
const result = $('23 + 42');
```

### Querying the DOM (JS specific)

You can use makeitso to interact with the Document Object Model (DOM) in your web applications. To retrieve elements from the DOM, try the following syntax:

```javascript
const elems = $("div.item");
```

This will select all `div` elements with the class "item" in the document. Probably...

### Querying the DOM in English

makeitso supports querying the DOM using plain language for more natural interactions. For example:

```javascript
const elems = $('give me all buttons');
```

This will select all the button elements on the page. If you're lucky.

### Querying a Database

You can also use makeitso to query a database and retrieve data.

```javascript
const users = $("return all users from the local sqlite3 database users.db");
```

This command will fetch all the users' records from the database. Under certain circumstances.

### Querying Your Cat

```javascript
const mood = $("How is my cat doing?");
```

### Fixing (syntax) errors:
```javascript
function fooBar(a {
  consoleelog(a)
}

dooBar("hello world');

$('fix all syntax errors');
```

### Fixing mistakes
```javascript
console.log("2 + 2 = 5")
console.log("The D minor scale is C D E F G A B C")
console.log("The first president of the United States was Chuck Norris.")

$('fix all mistakes');
```

## Running

### Building
You need a working Go environment. Then simply:

```
make
make install (optionally)
make clean (optionally)
```

### Executing

*An OpenAI API key is required. Expose the key in the environment variable `OPENAI_KEY`*

To process a file, you simply run makeitso with the file as an argument. The output will be returned on stdout. That means if you want the output in a file, simply redirect it:

```
makeitso my_file.js > my_file.out.js
```

This will generate the file `my_file.out.js` with (hopefully) all `$(prompt)` statements replaced by delicious AI-generated code.

#### Flags

##### Language

By default, makeitso tries to guess the language of the source file. To specify a language, use the `-l` (for language) flag:

```
makeitso -l typescript my_file.ts
```

##### Flavor

To generate output in different styles, you can use the `-f` (flavor) flag. Available flavors are `junior`, `senior`, and `bigbrain`:

```
makeitso -f junior my_file.js
```

Feel free to experiment with makeitso, but always remember to double-check the generated code, as it may not always align with your expectations. 


#### Include in your build pipeline

I'm sure you'll find out a way how to do this, but here's a simple evil idea for your package.json:

```json
{
  "name": "lazycode",
  "version": "1.0.0",
  "main": "index.js",
  "scripts": {
    "prod": "makeitso -l javascript index.js > index.out.js && node index.out.js"
  }
}
```


## Final thoughts

You can see how powerful this can be.

And hopefully, you can see that it's not a good idea. It's a bad idea. A very bad one.

I'm sorry!

Please don't use this.
I had to make this.
Roko's Basilik made me do it.
