const fs = require('fs');
const { encode } = require('./encode');

const input = fs
    .readFileSync('./input.txt', 'utf8');

const result = encode(input);

fs.writeFileSync('./output.txt', result);
