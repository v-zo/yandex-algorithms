const { encode } = require('./encode')

it('case 1', () => {
    expect(
        encode(
            ['6', '-1 -1 -9 -7 3 -6']
        )
    ).toBe('1')
})

it('case 2', () => {
    expect(
        encode(
            ['5', '7 -8 2 -8 -3']
        )
    ).toBe('6')
})

it('case 3', () => {
    expect(
        encode(
            ['8', '6 2 8 -3 1 1 6 10']
        )
    ).toBe('8')
})

it('case 4', () => {
    expect(
        encode(
            ['6', '-1 -1 -9 -7 3 -6']
        )
    ).toBe('1')
})

it('case 5', () => {
    expect(
        encode(
            ['5', '7 -8 2 -8 -3']
        )
    ).toBe('6')
})

it('case 6', () => {
    expect(
        encode(
            ['8', '6 2 8 -3 1 1 6 10']
        )
    ).toBe('8')
})

it('case 7', () => {
    expect(
        encode(
            ['4', '-10 -9 8 -8']
        )
    ).toBe('-9')
})