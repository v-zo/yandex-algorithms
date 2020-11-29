const { encode } = require('./encode')

it('find numbers test 1', () => {
    expect(
        encode(
            ['7', '6 4 1 2 3']
        )
    ).toBe('5 7')
})

// it('find numbers test 2', () => {
//     expect(
//         encode(
//             ['17', '17 15 13 11 3 8 6 10 12 1 4 14 5 2 16']
//         )
//     ).toBe('17 9')
// })