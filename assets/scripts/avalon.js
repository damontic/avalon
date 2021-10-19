class Avalon {
    constructor(width, height){
        this.width = width
        this.height = height
        // backgrounds from farthest to closer
        this.background_zs = [-500]
        this.merlin_card = new AvalonCard(IMAGE_MERLIN, IMAGE_CHARACTER_CARD_BACK, 100, 100)
        this.morgana_card = new AvalonCard(IMAGE_MORGANA, IMAGE_CHARACTER_CARD_BACK, 250, 100)
    }
    draw() {
        background(BLACK)
        push()
        // Sets the origin to the upper left of the canvas
        // Add some margin space
        translate(
            - this.width / 2,
            - this.height / 2
        )

        if (rooms != undefined) {
            this.merlin_card.draw()
            this.morgana_card.draw()
        }
        pop()
    }
}
