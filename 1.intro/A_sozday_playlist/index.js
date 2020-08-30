const fs = require('fs');
const { createPlaylist } = require('./createPlaylist');

const input = fs
    .readFileSync('./input.txt', 'utf8')
    .split('\n');

const result = createPlaylist(input);

fs.writeFileSync('./output.txt', result);
