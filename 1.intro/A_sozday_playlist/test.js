const { createPlaylist } = require('./createPlaylist')

it('createPlaylist test 1', () => {
    expect(
        createPlaylist(
            ['3', '1 2 3', '4 5 6']
        )
    ).toBe('1 4 2 5 3 6')
})

it('createPlaylist test 2', () => {
    expect(
        createPlaylist(
            ['1', '1', '2']
        )
    ).toBe('1 2')
})