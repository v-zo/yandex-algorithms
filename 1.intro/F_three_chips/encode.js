const getCache = (data, ind1, ind2, arcmN) => Math.abs(data[ind1] + data[ind2] + arcmN);

const fun = (data, N, indC) => {
    let ind1, ind2;

    if (indC === 0) {
        ind1 = data[1] > data[2] ? 1 : 2;
        ind2 = data[1] > data[2] ? 2 : 1;
    } else if (indC === 1) {
        ind1 = data[0] > data[2] ? 0 : 2;
        ind2 = data[0] > data[2] ? 2 : 0;
    } else {
        ind1 = data[0] > data[1] ? 0 : 1;
        ind2 = data[0] > data[1] ? 1 : 0;
    }

    const arcmN = data[indC] - N;
    let cache = getCache(data, ind1, ind2, arcmN);

    data.forEach((el, i, arr) => {
        if (indC === i) {
            return;
        }

        const elSum = el + arcmN;

        if (ind2 !== i &&
            (Math.abs(elSum + arr[ind2]) < cache)
        ) {
            ind1 = i;
            cache = getCache(arr, ind1, ind2, arcmN)
            return;
        }

        if (ind1 !== i &&
            (Math.abs(elSum + arr[ind1]) < cache)
        ) {
            ind2 = i;
            cache = getCache(arr, ind1, ind2, arcmN)
            return;
        }
    });

    return data[ind1] + data[ind2] + data[indC];
}

const rec = (data, N) => {
    let res = fun(data, N, 0);
    let clos = Math.abs(res - N);

    for (i = 1; i < data.length; i++) {
        const curr = fun(data, N, i);
        const diff = Math.abs(curr - N);

        if (curr === N) {
            return curr.toString();
        }

        if (diff < clos) {
            clos = diff;
            res = curr;
        }
    }

    return res.toString();
}

const encode = (lines) => {
    const N = +lines[0];
    const data = lines[1].split(' ').map(n => +n)

    return rec(data, N);
}

module.exports.encode = encode;