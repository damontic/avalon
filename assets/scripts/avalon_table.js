class AvalonTable {
    constructor(avalon, x, y, w, h, data) {
        this.avalon = avalon
        this.x = x
        this.y = y
        this.width = w
        this.height = h
        this.data = data
    }
    draw() {
        push()
        fill(0, 255, 0)
        rect(
            this.x,
            this.y,
            this.width,
            this.height
        )
        pop()
    }
}
