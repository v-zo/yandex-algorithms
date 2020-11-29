const fs = require('fs');
const { encode } = require('./encode');

const input = fs
    .readFileSync('./input.txt', 'utf8')
    .split('\n')
    .map(st => st.split(' '));

const result = encode(input);

fs.writeFileSync('./output.txt', result);
