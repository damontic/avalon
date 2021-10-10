class AvalonMenu {
    constructor(avalon, x, y, w, h) {
        this.avalon = avalon
        this.x = x
        this.y = y
        this.width = w
        this.height = h

        let roomsData = {}
        this.roomsTable = new AvalonTable(
            avalon,
            this.x,
            this.y,
            this.width,
            this.height,
            roomsData
        )

        this.createRoomButton = new AvalonButton(
            avalon,
            'Create Game!',
            this.avalon.left_margin,
            this.avalon.upper_margin + this.avalon.title.height,
            FONT_INCONSOLATA
        )
    }
    draw() {
        push()
        translate(0,0,this.avalon.background_zs[0])
        this.roomsTable.draw()
        this.createRoomButton.draw()
        pop()
    }
}
