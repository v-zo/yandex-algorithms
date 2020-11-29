const encode = (line) => {
    let currentNum = +line;
    let reversed = 0;
    while (currentNum) {
        const digit = currentNum % 10;
        reversed = (reversed * 10) + digit;
        currentNum = parseInt(currentNum / 10);
    }

    return reversed.toString();
}

module.exports.encode = encode;