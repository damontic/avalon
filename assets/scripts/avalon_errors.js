class UnsupportedNumberOfPlayers extends Error {
    constructor(...params) {
        super(...params);
        this.name = "UnsupportedNumberOfPlayers";
    }
}
