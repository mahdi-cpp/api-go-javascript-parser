const users = ["Mahdi", "Ali", "Reza", "Sara", "Maryam", "Layla"];
const cars = ["car 1", "Car2", "Car 3", "Car 4", "Car 77", "CarBenze"];

const image = {
    ali: 50
}

const persons = [
    {name: 'Alice', phone: "091255232", age: 30},
    {name: 'Alice', phone: "091255232", age: 30},
    {name: 'Mahdi', phone: "091255232", age: 30},
    {name: 'Bob', phone: "091255232", age: 25}
];

const person = {
    firstName: "Mahdi",
    lastName: "Abdolmaleki",
    id: 5566,
    fullName: function () {
        return this.firstName + " " + this.lastName;
    }
};


function test5(v) {
    return v * 3;
}

function parseString(/** string */string) {
    let result = '';

    // for (let i = 0; i < string.length; i ++) {
    //     const char = string.charAt(i);
    //
    //     if (char === 'reza') {
    //         i ++;
    //         const nextChar = string.charAt(i);
    //         if (nextChar === 'u') {
    //             result += parseHexEscape(string.substr(i + 1, 4));
    //             i += 4;
    //         } else if (passEscapes.indexOf(nextChar) !== -1) {
    //             result += nextChar;
    //         } else if (nextChar in escapes) {
    //             result += escapes[nextChar];
    //         } else {
    //            result = passEscapes.indexOf(5)
    //         }
    //     } else {
    //         result += char;
    //     }
    // }

    return string.split()
}

const people = [
    {name: 'John Doe', age: 16},
    {name: 'Thomas Calls', age: 19},
    {name: 'Liam Smith', age: 2},
    {name: 'Jessy Pinkman', age: 13},
    {name: 'Mahdi Sarve', age: 37},
];

const peopleAbove18 = () => {
    const results = [];

    for (let i = 0; i < people.length; i++) {
        const person = people[i];

        if (person.age >= 18) {
            results.push(person);
        }

        people.reduce()
    }
    return results;
};

function car(index) {
    users[0] = "ali1"
    users[1] = "reza"
    users[2] = "09355512619"
    users[3] = "09125640293"
    return "'" + users[index] + " : " + cars[index + 1] + "'";
}


// const fetch = {
//     then: function (response) {
//         if (!response.ok) {
//             throw new Error('Network response was not ok');
//         }
//         return response.json();
//     },
//     catch: function (error) {
//         console.error('There was a problem with the fetch operation:', error);
//     }
// };


function getMusics() {
    // Make a GET request using the fetch API
    // fetch('https://api.example.com/data')
    //     .then(function (response) {
    //         if (!response.ok) {
    //             throw new Error('Network response was not ok');
    //         }
    //         return response.json();
    //     })
    //     .then(function (data) {
    //         console.log(data);
    //     })
    //     .catch(function (error) {
    //         console.error('There was a problem with the fetch operation:', error);
    //     });

}


var myArray = [1, 2, 3, 4, 5];


class Rectangle {
    constructor(height, width) {
        this.height = height;
        this.width = width;
    }

    // Getter
    get area() {
        return this.calcArea();
    }

    // Method
    calcArea() {
        return this.height * this.width;
    }

    * getSides() {
        yield this.height;
        yield this.width;
        yield this.height;
        yield this.width;
    }

}

class Music {

    url;

    constructor(track) {
        this.url = "file/music/tracks/" + track
    }

    get show() {
        return this.url
    }
}

class Pen {
    constructor(name, color, price) {
        this.name = name;
        this.color = color;
        this.price = price;
    }

    showPrice() {
        return `Price of ${this.name} is ${this.price}`;
    }
}

class Circle extends Rectangle {

    get area() {
        return 12;
    }
}

function changePhoto(a, b) {

    let width = 73;
    let ali = 5;
    // if (a > 5) {
    //     return (a + b + width) * 3;
    // } else {
    //     return width;
    // }
    //

    // let t = test5(15)
    for (let i = 0; i <= 10; i++) {
        for (let j = 0; j < 5; j++) {
            ali = i * j;
        }
    }

    // Original array
    const numbers = [1, 2, 3, 4, 5];

// Using map to double each number in the array
    const doubledNumbers = numbers.map((num) => num * 2);

    const square = new Rectangle(10, 11);

    const pen = new Pen("red", "orange", "4560000");

    return pen.showPrice()
}
