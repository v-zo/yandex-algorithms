const createPlaylist = (lines) => {
    const n = +lines[0];
    const rusIds = lines[1].split(' ');
    const engIds = lines[2].split(' ');

    const result = [];

    for (let i = 0; i < n; i++) {
        result.push(rusIds[i]);
        result.push(engIds[i]);
    }

    return result.join(' ');
}

module.exports.createPlaylist = createPlaylist;