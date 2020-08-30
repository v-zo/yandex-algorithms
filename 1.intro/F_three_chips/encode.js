fun = (data, N, indC) => {
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

    data.forEach((el, i, arr) => {
        if (indC === i) {
            return;
        }

        if (
            Math.abs(el + arr[ind2] + arr[indC] - N)
            < Math.abs(arr[ind1] + arr[ind2] + arr[indC] - N)
        ) {
            if (ind2 === i) {
                return;
            }
            ind1 = i;
            return;
        }

        if (
            Math.abs(el + arr[ind1] + arr[indC] - N)
            < Math.abs(arr[ind1] + arr[ind2] + arr[indC] - N)
        ) {
            if (ind1 === i) {
                return;
            }
            ind2 = i;
            return;
        }
    });

    return [data[ind1], data[ind2], data[indC]];
}

rec = (data, N) => {
    res = fun(data, N, 0);
    clos = Math.abs((res[0] + res[1] + res[2]) - N);

    for (i = 1; i < data.length; i++) {
        cc = fun(data, N, i);
        sum = Math.abs((cc[0] + cc[1] + cc[2]) - N);

        if (sum < clos) {
            clos = sum;
            res = cc;
        }
    }

    return (res[0] + res[1] + res[2]).toString();
}

const encode = (lines) => {
    const N = +lines[0];
    const data = lines[1].split(' ').map(n => +n)

    return rec(data, N);
}

module.exports.encode = encode;
module.exports.rec = rec;