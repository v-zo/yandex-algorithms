const encode = (lines) => {
    const N = lines[0];
    let emptys;
    (emptys = []).length = N; emptys.fill(0);
    const ids = lines[1].split(' ');

    ids.forEach(id => {
        emptys[id - 1] = id;
    })

    const f1 = 1 + emptys.findIndex((_, i, arr) => arr[i] === 0);
    const sliced = emptys.slice(f1,);
    const f2 = 1 + f1 + sliced.findIndex((_, i, arr) => arr[i] === 0)

    return `${f1} ${f2}`;
}

module.exports.encode = encode;