const encode = (lines) => {
    const N = lines[0][0];
    const result = [];
    const c = N / 2 | 0;

    const rows = lines.slice(1,);

    for (let rad = 0; rad <= c; rad++) {
        for (let m = - rad + 1; m < rad; m++) {
            result.push(rows[c - rad][c + m]);
        }

        for (let k = c - rad; k < c + rad; k++) {
            result.push(rows[k][c + rad]);
        }

        for (let m = c + rad; m > c - rad; m--) {
            result.push(rows[c + rad][m]);
        }

        for (let k = c + rad; k >= c - rad; k--) {
            result.push(rows[k][c - rad]);
        }
    }

    return result.join('\n');
}

module.exports.encode = encode;