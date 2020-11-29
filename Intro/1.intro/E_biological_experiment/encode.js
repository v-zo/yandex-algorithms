const encode = (lines) => {
    const N = +lines[0];
    const data = lines.slice(1,).map(item => item.split(' '));

    const entities = data.reduce((acc, item) => {
        acc[item[0]] = {
            time: +item[1],
            childs: [item[2], item[3]].filter(x => x > 0)
        }

        return acc;
    }, {});

    const ids = data.map(item => item[0]);

    const gens = [{
        count: entities[ids[0]].time,
        length: 1
    }];

    reducer = (a, b) => entities[a].time + entities[b].time;

    processChilds = (childs, gen) => {
        gen++;

        if (!gens[gen]) {
            gens[gen] = {
                count: 0,
                length: 0
            }
        }


        childs.forEach(childId => {
            const child = entities[childId];

            gens[gen].count += child.time;
            gens[gen].length += 1;

            if (!child.childs.length) {
                return;
            }

            processChilds(child.childs, gen);
        })
    }

    processChilds(entities[ids[0]].childs, 0);

    const result = gens
        .map(({ count, length }) => count / length)
        .filter(x => x === 0 ? true : x);

    return result.map(x => Number
        .parseFloat(x)
        .toFixed(2)
    ).join(' ');
}

module.exports.encode = encode;